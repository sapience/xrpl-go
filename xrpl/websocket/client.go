package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync/atomic"
	"time"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/hash"
	transaction "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/mitchellh/mapstructure"

	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/server"
	streamtypes "github.com/Peersyst/xrpl-go/xrpl/queries/subscription/types"
	requests "github.com/Peersyst/xrpl-go/xrpl/queries/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/xrpl/websocket/interfaces"
	wstypes "github.com/Peersyst/xrpl-go/xrpl/websocket/types"
	ws "github.com/gorilla/websocket"

	commonconstants "github.com/Peersyst/xrpl-go/xrpl/common"
)

const (
	DefaultFeeCushion float32 = 1.2
	DefaultMaxFeeXRP  float32 = 2
)

var (
	ErrIncorrectID          = errors.New("incorrect id")
	ErrNotConnectedToServer = errors.New("not connected to server")
	ErrRequestTimedOut      = errors.New("request timed out")
)

type Client struct {
	cfg  ClientConfig
	conn *Connection

	// Channels
	errChan          chan error
	requestChan      chan *ClientResponse
	ledgerClosedChan chan *streamtypes.LedgerStream
	validationChan   chan *streamtypes.ValidationStream
	transactionChan  chan *streamtypes.TransactionStream
	peerStatusChan   chan *streamtypes.PeerStatusStream
	orderBookChan    chan *streamtypes.OrderBookStream
	bookChangesChan  chan *streamtypes.BookChangesStream
	consensusChan    chan *streamtypes.ConsensusStream

	idCounter atomic.Uint32
	NetworkID uint32
}

// Creates a new websocket client with cfg.
// This client will open and close a websocket connection for each request.
func NewClient(cfg ClientConfig) *Client {
	return &Client{
		cfg:         cfg,
		requestChan: make(chan *ClientResponse),
		errChan:     make(chan error),
		conn:        NewConnection(cfg.host),
	}
}

// Connect opens a websocket connection to the server. It starts reading messages in a goroutine.
func (c *Client) Connect() error {
	err := c.conn.Connect()
	if err != nil {
		return err
	}
	go c.readMessages()
	return nil
}

// Disconnect closes the websocket connection.
func (c *Client) Disconnect() error {
	return c.conn.Disconnect()
}

// IsConnected returns true if the client is connected to the server.
func (c *Client) IsConnected() bool {
	return c.conn.IsConnected()
}

func (c *Client) FaucetProvider() commonconstants.FaucetProvider {
	return c.cfg.faucetProvider
}

// Autofill fills in the missing fields in a transaction.
func (c *Client) Autofill(tx *transaction.FlatTransaction) error {
	if err := c.setValidTransactionAddresses(tx); err != nil {
		return err
	}

	err := c.setTransactionFlags(tx)
	if err != nil {
		return err
	}

	if _, ok := (*tx)["NetworkID"]; !ok {
		if c.NetworkID != 0 {
			(*tx)["NetworkID"] = c.NetworkID
		}
	}
	if _, ok := (*tx)["Sequence"]; !ok {
		err := c.setTransactionNextValidSequenceNumber(tx)
		if err != nil {
			return err
		}
	}
	if _, ok := (*tx)["Fee"]; !ok {
		err := c.calculateFeePerTransactionType(tx, 0)
		if err != nil {
			return err
		}
	}
	if _, ok := (*tx)["LastLedgerSequence"]; !ok {
		err := c.setLastLedgerSequence(tx)
		if err != nil {
			return err
		}
	}
	if txType, ok := (*tx)["TransactionType"].(transaction.TxType); ok {
		if acc, ok := (*tx)["Account"].(types.Address); txType == transaction.AccountDeleteTx && ok {
			err := c.checkAccountDeleteBlockers(acc)
			if err != nil {
				return err
			}
		}
		if txType == transaction.PaymentTx {
			err := c.checkPaymentAmounts(tx)
			if err != nil {
				return err
			}
		}
		if txType == transaction.BatchTx {
			// Convert FlatTransaction to Batch for autofilling
			if batchTx, ok := (*tx)["RawTransactions"]; ok {
				batch := &transaction.Batch{
					RawTransactions: batchTx.([]types.RawTransaction),
				}
				if account, ok := (*tx)["Account"].(types.Address); ok {
					batch.Account = account
				}
				err := c.AutofillBatch(batch)
				if err != nil {
					return err
				}
				// Update the original transaction with the autofilled RawTransactions
				(*tx)["RawTransactions"] = batch.RawTransactions
			}
		}
	}
	return nil
}

// AutofillMultisigned fills in the missing fields in a multisigned transaction.
// This function is used to fill in the missing fields in a multisigned transaction.
// It fills in the missing fields in the transaction and calculates the fee per number of signers.
func (c *Client) AutofillMultisigned(tx *transaction.FlatTransaction, nSigners uint64) error {
	err := c.Autofill(tx)
	if err != nil {
		return err
	}

	err = c.calculateFeePerTransactionType(tx, nSigners)
	if err != nil {
		return err
	}

	return nil
}

// FundWallet funds a wallet with XRP from the faucet.
// If the wallet does not have a classic address, it will return an error.
func (c *Client) FundWallet(wallet *wallet.Wallet) error {
	if wallet.ClassicAddress == "" {
		return errors.New("fund wallet: cannot fund a wallet without a classic address")
	}

	err := c.cfg.faucetProvider.FundWallet(wallet.ClassicAddress)
	if err != nil {
		return err
	}

	return nil
}

// Request sends a request to the server and returns the response.
// This function is used to send requests to the server.
// It returns the response from the server.
func (c *Client) Request(req interfaces.Request) (*ClientResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id := c.idCounter.Add(1)

	msg, err := c.formatRequest(req, int(id), nil)
	if err != nil {
		return nil, err
	}

	if !c.IsConnected() {
		return nil, ErrNotConnectedToServer
	}

	err = c.conn.WriteMessage(msg)
	if err != nil {
		return nil, err
	}

	res, err := c.awaitResponse(int(id))
	if err != nil {
		return nil, err
	}

	if res.ID != int(id) {
		return nil, ErrIncorrectID
	}
	if err := res.CheckError(); err != nil {
		return nil, err
	}

	return res, nil
}

// SubmitTxBlob sends a pre-signed transaction blob to the server.
// It decodes the blob to confirm that it contains either a signature
// or a signing public key, and then submits it using a submission request.
// The failHard flag determines how strictly errors are handled.
func (c *Client) SubmitTxBlob(txBlob string, failHard bool) (*requests.SubmitResponse, error) {
	tx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return nil, err
	}

	_, okTxSig := tx["TxSignature"].(string)
	_, okPubKey := tx["SigningPubKey"].(string)

	if !okTxSig && !okPubKey {
		return nil, ErrMissingTxSignatureOrSigningPubKey
	}

	return c.submitRequest(&requests.SubmitRequest{
		TxBlob:   txBlob,
		FailHard: failHard,
	})
}

// SubmitTx signs the transaction (if necessary) and submits it to the server
// via a submission request. It applies the provided submit options to decide whether
// to autofill missing fields and enforce failHard mode during submission.
func (c *Client) SubmitTx(tx transaction.FlatTransaction, opts *wstypes.SubmitOptions) (*requests.SubmitResponse, error) {
	txBlob, err := c.getSignedTx(tx, opts.Autofill, opts.Wallet)
	if err != nil {
		return nil, err
	}

	return c.submitRequest(&requests.SubmitRequest{
		TxBlob:   txBlob,
		FailHard: opts.FailHard,
	})
}

// SubmitMultisigned sends a multisigned transaction to the server and returns the response.
// This function is used to send multisigned transactions to the server.
// It returns the response from the server.
func (c *Client) SubmitMultisigned(txBlob string, failHard bool) (*requests.SubmitMultisignedResponse, error) {
	tx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return nil, err
	}
	signers, okSigners := tx["Signers"].([]interface{})

	if okSigners && len(signers) > 0 {
		for _, sig := range signers {
			signer := sig.(map[string]any)
			signerData := signer["Signer"].(map[string]any)
			if signerData["SigningPubKey"] == "" && signerData["TxnSignature"] == "" {
				return nil, errors.New("signer data is empty")
			}
		}
	}

	return c.submitMultisignedRequest(&requests.SubmitMultisignedRequest{
		Tx:       tx,
		FailHard: failHard,
	})
}

// SubmitTxBlobAndWait sends a pre-signed transaction blob to the server,
// decodes it to retrieve the required LastLedgerSequence, submits the blob,
// and then waits until the transaction is confirmed in a ledger. It returns
// the transaction response if the submission is successful.
func (c *Client) SubmitTxBlobAndWait(txBlob string, failHard bool) (*requests.TxResponse, error) {
	tx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return nil, err
	}

	lastLedgerSequence, ok := tx["LastLedgerSequence"].(uint32)
	if !ok {

		return nil, ErrMissingLastLedgerSequenceInTransaction

	}
	txResponse, err := c.SubmitTxBlob(txBlob, failHard)
	if err != nil {
		return nil, err
	}

	if txResponse.EngineResult != "tesSUCCESS" {
		return nil, &ClientError{ErrorString: "transaction failed to submit with engine result: " + txResponse.EngineResult}
	}

	txHash, err := hash.SignTxBlob(txBlob)
	if err != nil {
		return nil, err
	}

	return c.waitForTransaction(txHash, lastLedgerSequence)
}

// SubmitTxAndWait prepares a transaction by ensuring it is fully signed,
// submits it to the server, and waits for ledger confirmation.
// It validates that the transaction's EngineResult is successful before returning
// the transaction response.
func (c *Client) SubmitTxAndWait(tx transaction.FlatTransaction, opts *wstypes.SubmitOptions) (*requests.TxResponse, error) {
	// Get the signed transaction blob.
	txBlob, err := c.getSignedTx(tx, opts.Autofill, opts.Wallet)
	if err != nil {
		return nil, err
	}

	// Delegate to SubmitTxBlobAndWait to handle submission, engine result check,
	// ledger sequence validation, and waiting for confirmation.
	return c.SubmitTxBlobAndWait(txBlob, opts.FailHard)
}

func (c *Client) waitForTransaction(txHash string, lastLedgerSequence uint32) (*requests.TxResponse, error) {
	var txResponse *requests.TxResponse
	i := 0

	for i < c.cfg.maxRetries {
		// Get the current ledger index
		currentLedger, err := c.GetLedgerIndex()
		if err != nil {
			return nil, err
		}

		// Check if the transaction has been included in the current ledger
		if currentLedger.Int() >= int(lastLedgerSequence) {
			break
		}

		// Request the transaction from the server
		res, err := c.Request(&requests.TxRequest{
			Transaction: txHash,
		})
		if err != nil {
			return nil, err
		}

		err = res.GetResult(&txResponse)
		if err != nil {
			return nil, err
		}

		// Check if the transaction has been included in the current ledger
		if txResponse.LedgerIndex.Int() >= int(lastLedgerSequence) {
			break
		}

		// Wait for the retry delay before retrying
		time.Sleep(c.cfg.retryDelay)
		i++
	}

	if txResponse == nil {
		return nil, errors.New("transaction not found")
	}

	return txResponse, nil
}

func (c *Client) submitMultisignedRequest(req *requests.SubmitMultisignedRequest) (*requests.SubmitMultisignedResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var subRes requests.SubmitMultisignedResponse
	err = res.GetResult(&subRes)
	if err != nil {
		return nil, err
	}
	return &subRes, nil
}

func (c *Client) submitRequest(req *requests.SubmitRequest) (*requests.SubmitResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var subRes requests.SubmitResponse
	err = res.GetResult(&subRes)
	if err != nil {
		return nil, err
	}
	return &subRes, nil
}

func (c *Client) formatRequest(req interfaces.Request, id int, marker any) ([]byte, error) {
	m := make(map[string]any)
	m["id"] = id
	m["command"] = req.Method()
	m["api_version"] = req.APIVersion()
	if marker != nil {
		m["marker"] = marker
	}
	dec, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &m})
	err := dec.Decode(req)
	if err != nil {
		return nil, err
	}

	return json.Marshal(m)
}

// TODO: Implement this when IsValidXAddress is implemented
func (c *Client) getClassicAccountAndTag(address string) (string, uint32) {
	return address, 0
}

func (c *Client) convertTransactionAddressToClassicAddress(tx *transaction.FlatTransaction, fieldName string) {
	if address, ok := (*tx)[fieldName].(string); ok {
		classicAddress, _ := c.getClassicAccountAndTag(address)
		(*tx)[fieldName] = classicAddress
	}
}

func (c *Client) validateTransactionAddress(tx *transaction.FlatTransaction, addressField, tagField string) error {
	classicAddress, tag := c.getClassicAccountAndTag((*tx)[addressField].(string))
	(*tx)[addressField] = classicAddress

	if tag != uint32(0) {
		if txTag, ok := (*tx)[tagField].(uint32); ok && txTag != tag {
			return fmt.Errorf("the %s, if present, must be equal to the tag of the %s", addressField, tagField)
		}
		(*tx)[tagField] = tag
	}

	return nil
}

// Sets valid addresses for the transaction.
func (c *Client) setValidTransactionAddresses(tx *transaction.FlatTransaction) error {
	// Validate if "Account" address is an xAddress
	if err := c.validateTransactionAddress(tx, "Account", "SourceTag"); err != nil {
		return err
	}

	if _, ok := (*tx)["Destination"]; ok {
		if err := c.validateTransactionAddress(tx, "Destination", "DestinationTag"); err != nil {
			return err
		}
	}

	// DepositPreuaht
	c.convertTransactionAddressToClassicAddress(tx, "Authorize")
	c.convertTransactionAddressToClassicAddress(tx, "Unauthorize")
	// EscrowCancel, EscrowFinish
	c.convertTransactionAddressToClassicAddress(tx, "Owner")
	// SetRegularKey
	c.convertTransactionAddressToClassicAddress(tx, "RegularKey")

	return nil
}

// Sets the next valid sequence number for a given transaction.
func (c *Client) setTransactionNextValidSequenceNumber(tx *transaction.FlatTransaction) error {
	if _, ok := (*tx)["Account"].(string); !ok {
		return errors.New("missing Account in transaction")
	}
	res, err := c.GetAccountInfo(&account.InfoRequest{
		Account:     types.Address((*tx)["Account"].(string)),
		LedgerIndex: common.LedgerTitle("current"),
	})

	if err != nil {
		return err
	}

	(*tx)["Sequence"] = uint32(res.AccountData.Sequence)
	return nil
}

// Calculates the current transaction fee for the ledger.
// Note: This is a public API that can be called directly.
func (c *Client) getFeeXrp(cushion float32) (string, error) {
	res, err := c.GetServerInfo(&server.InfoRequest{})
	if err != nil {
		return "", err
	}

	if res.Info.ValidatedLedger.BaseFeeXRP == 0 {
		return "", errors.New("getFeeXrp: could not get BaseFeeXrp from ServerInfo")
	}

	loadFactor := res.Info.LoadFactor
	if res.Info.LoadFactor == 0 {
		loadFactor = 1
	}

	fee := res.Info.ValidatedLedger.BaseFeeXRP * float32(loadFactor) * cushion

	if fee > c.cfg.maxFeeXRP {
		fee = c.cfg.maxFeeXRP
	}

	// Round fee to NUM_DECIMAL_PLACES
	roundedFee := float32(math.Round(float64(fee)*math.Pow10(int(currency.MaxFractionLength)))) / float32(math.Pow10(int(currency.MaxFractionLength)))

	// Convert the rounded fee back to a string with NUM_DECIMAL_PLACES
	return fmt.Sprintf("%.*f", currency.MaxFractionLength, roundedFee), nil
}

// Calculates the fee per transaction type.
//
// Enhanced implementation that replicates xrpl.js calculateFeePerTransactionType logic,
// including special cases for EscrowFinish, AccountDelete, AMMCreate, Batch, and multi-signing.
func (c *Client) calculateFeePerTransactionType(tx *transaction.FlatTransaction, nSigners uint64) error {
	// Get base network fee
	netFeeXRP, err := c.getFeeXrp(c.cfg.feeCushion)
	if err != nil {
		return err
	}

	netFeeDrops, err := currency.XrpToDrops(netFeeXRP)
	if err != nil {
		return err
	}

	// Convert to uint64 for calculations
	baseFeeUint, err := strconv.ParseUint(netFeeDrops, 10, 64)
	if err != nil {
		return err
	}

	baseFee := baseFeeUint

	// Get transaction type
	transactionType := ""
	if txType, ok := (*tx)["TransactionType"]; ok {
		if str, ok := txType.(string); ok {
			transactionType = str
		}
	}

	// Check if this is a special transaction cost type
	isSpecialTxCost := transactionType == "AccountDelete" || transactionType == "AMMCreate"

	// EscrowFinish Transaction with Fulfillment
	if transactionType == "EscrowFinish" {
		if fulfillment, ok := (*tx)["Fulfillment"]; ok && fulfillment != nil {
			if fulfillmentStr, ok := fulfillment.(string); ok && fulfillmentStr != "" {
				fulfillmentBytesSize := (len(fulfillmentStr) + 1) / 2 // Math.ceil(length / 2)
				// BaseFee × (33 + (Fulfillment size in bytes / 16))
				baseFee = baseFeeUint * (33 + uint64(fulfillmentBytesSize)/16)
			}
		}
	} else if isSpecialTxCost {
		// For AccountDelete and AMMCreate, use owner reserve fee
		reserveFee, err := c.fetchOwnerReserveFee()
		if err != nil {
			return err
		}
		baseFee = reserveFee
	} else if transactionType == "Batch" {
		// For Batch transactions, calculate fee for all inner transactions
		rawTxFees, err := c.calculateBatchFees(tx)
		if err != nil {

			return err
		}

		// baseFee = BigNumber.sum(baseFee.times(2), rawTxFees)
		baseFee = baseFeeUint*2 + rawTxFees
	}

	// Multi-signed Transaction: BaseFee × (1 + Number of Signatures Provided)
	if nSigners > 0 {
		signersFee := baseFeeUint * nSigners
		baseFee = baseFee + signersFee
	}

	// Apply max fee limit (but not for special transaction cost types)
	var totalFee uint64
	if isSpecialTxCost {
		totalFee = baseFee
	} else {
		maxFeeDrops, err := currency.XrpToDrops(fmt.Sprintf("%.6f", c.cfg.maxFeeXRP))
		if err != nil {
			return err
		}
		maxFeeUint, err := strconv.ParseUint(maxFeeDrops, 10, 64)
		if err != nil {
			return err
		}
		if baseFee < maxFeeUint {
			totalFee = baseFee
		} else {
			totalFee = maxFeeUint
		}
	}

	(*tx)["Fee"] = strconv.FormatUint(totalFee, 10)
	return nil
}

// Sets the latest validated ledger sequence for the transaction.
// Modifies the `LastLedgerSequence` field in the tx.
func (c *Client) setLastLedgerSequence(tx *transaction.FlatTransaction) error {
	index, err := c.GetLedgerIndex()
	if err != nil {
		return err
	}

	(*tx)["LastLedgerSequence"] = index.Uint32() + commonconstants.LedgerOffset
	return err
}

// Checks for any blockers that prevent the deletion of an account.
// Returns nil if there are no blockers, otherwise returns an error.
func (c *Client) checkAccountDeleteBlockers(address types.Address) error {
	accObjects, err := c.GetAccountObjects(&account.ObjectsRequest{
		Account:              address,
		LedgerIndex:          common.LedgerTitle("validated"),
		DeletionBlockersOnly: true,
	})
	if err != nil {
		return err
	}

	if len(accObjects.AccountObjects) > 0 {
		return errors.New("account %s cannot be deleted; there are Escrows, PayChannels, RippleStates, or Checks associated with the account")
	}
	return nil
}

func (c *Client) checkPaymentAmounts(tx *transaction.FlatTransaction) error {
	if _, ok := (*tx)["DeliverMax"]; ok {
		if _, ok := (*tx)["Amount"]; !ok {
			(*tx)["Amount"] = (*tx)["DeliverMax"]
		} else if (*tx)["Amount"] != (*tx)["DeliverMax"] {
			return errors.New("payment transaction: Amount and DeliverMax fields must be identical when both are provided")
		}
	}
	return nil
}

// Sets a transaction's flags to its numeric representation.
// TODO: Add flag support for AMMDeposit, AMMWithdraw,
// NFTTOkenCreateOffer, NFTokenMint, OfferCreate, XChainModifyBridge (not supported).
func (c *Client) setTransactionFlags(tx *transaction.FlatTransaction) error {
	flags, ok := (*tx)["Flags"].(uint32)
	if !ok && flags > 0 {
		(*tx)["Flags"] = int(0)
		return nil
	}

	_, ok = (*tx)["TransactionType"].(string)
	if !ok {
		return errors.New("transaction type is missing in transaction")
	}

	return nil
}

func (c *Client) awaitResponse(id int) (*ClientResponse, error) {
	for {
		select {
		case res := <-c.requestChan:
			if res.ID == id {
				return res, nil
			}
		case <-time.After(c.cfg.timeout):
			return nil, ErrRequestTimedOut
		}
	}
}

func (c *Client) handleMessage(message []byte) {
	var stream wstypes.Message
	c.unmarshalMessage(message, &stream)
	if stream.IsRequest() {
		c.handleRequest(message)
	} else if stream.IsStream() {
		c.handleStream(stream.Type, message)
	}
}

func (c *Client) handleRequest(message []byte) {
	var res ClientResponse
	c.unmarshalMessage(message, &res)
	c.requestChan <- &res
}

func (c *Client) unmarshalMessage(message []byte, v any) {
	if err := json.Unmarshal(message, v); err != nil {
		if c.errChan == nil {
			c.errChan = make(chan error)
		}
		c.errChan <- err
	}
}

func (c *Client) handleStream(t streamtypes.Type, message []byte) {
	switch t {
	case streamtypes.LedgerStreamType:
		var ledger streamtypes.LedgerStream
		c.unmarshalMessage(message, &ledger)

		if c.ledgerClosedChan != nil {
			c.ledgerClosedChan <- &ledger
		}
	case streamtypes.TransactionStreamType:
		var transaction streamtypes.TransactionStream
		c.unmarshalMessage(message, &transaction)
		if c.transactionChan != nil {
			c.transactionChan <- &transaction
		}
	case streamtypes.ValidationStreamType:
		var validation streamtypes.ValidationStream
		c.unmarshalMessage(message, &validation)
		if c.validationChan != nil {
			c.validationChan <- &validation
		}
	case streamtypes.PeerStatusStreamType:
		var peerStatus streamtypes.PeerStatusStream
		c.unmarshalMessage(message, &peerStatus)
		if c.peerStatusChan != nil {
			c.peerStatusChan <- &peerStatus
		}
	case streamtypes.ConsensusStreamType:
		var consensus streamtypes.ConsensusStream
		c.unmarshalMessage(message, &consensus)
		if c.consensusChan != nil {
			c.consensusChan <- &consensus
		}
	default:
		if c.errChan == nil {
			c.errChan = make(chan error)
		}
		c.errChan <- fmt.Errorf("unknown stream type: %v", t)
	}
}

func (c *Client) readMessages() {
	retryCount := 0
	maxRetries := c.cfg.maxReconnects

	for {
		if c.conn == nil {
			return
		}
		message, err := c.conn.ReadMessage()
		switch {
		case ws.IsCloseError(err) || ws.IsUnexpectedCloseError(err):
			if retryCount >= maxRetries {
				if c.errChan == nil {
					c.errChan = make(chan error)
				}
				c.errChan <- fmt.Errorf("max reconnection attempts (%d) reached", maxRetries)
				return
			}
			retryCount++
			connErr := c.conn.Connect()
			if connErr != nil {
				if c.errChan == nil {
					c.errChan = make(chan error)
				}
				c.errChan <- connErr
				return
			}
		case err != nil:
			c.errChan <- err
			return
		default:
			// Send the message to the channel
			c.handleMessage(message)
			// Reset retry count on successful message
			retryCount = 0
		}
	}
}

// getSignedTx ensures the transaction is fully signed and returns the transaction blob.
// If the transaction is already signed, it encodes and returns it. Otherwise, it autofills (if enabled)
// and signs the transaction using the provided wallet.
func (c *Client) getSignedTx(tx transaction.FlatTransaction, autofill bool, wallet *wallet.Wallet) (string, error) {
	// Check if the transaction is already signed: both fields must be non-empty.
	sig, sigOk := tx["TxSignature"].(string)
	pubKey, pubKeyOk := tx["SigningPubKey"].(string)
	if sigOk && sig != "" && pubKeyOk && pubKey != "" {
		blob, err := binarycodec.Encode(tx)
		if err != nil {
			return "", err
		}
		return blob, nil
	}

	// If not signed, ensure a wallet is provided.
	if wallet == nil {
		return "", ErrMissingWallet
	}

	// Optionally autofill the transaction.
	if autofill {
		if err := c.Autofill(&tx); err != nil {
			return "", err
		}
	}

	// Sign the transaction.
	txBlob, _, err := wallet.Sign(tx)
	if err != nil {
		return "", err
	}
	return txBlob, nil
}

// fetchOwnerReserveFee fetches the owner reserve fee from the server state.
// Replicates the JavaScript fetchOwnerReserveFee function.
func (c *Client) fetchOwnerReserveFee() (uint64, error) {
	response, err := c.GetServerState(&server.StateRequest{})
	if err != nil {
		return 0, err
	}

	reserveInc := response.State.ValidatedLedger.ReserveInc
	if reserveInc == 0 {
		return 0, errors.New("could not fetch Owner Reserve")
	}

	return uint64(reserveInc), nil
}

// calculateBatchFees calculates the total fees for all inner transactions in a Batch.
// Replicates the JavaScript logic for Batch transaction fee calculation.
func (c *Client) calculateBatchFees(tx *transaction.FlatTransaction) (uint64, error) {
	var totalFees uint64 = 0

	// Get RawTransactions from the batch transaction
	rawTransactions, ok := (*tx)["RawTransactions"]
	if !ok {
		return 0, errors.New("RawTransactions field missing from Batch transaction")
	}

	// Convert to array of interfaces
	rawTxArray, ok := rawTransactions.([]interface{})
	if !ok {
		return 0, errors.New("RawTransactions field is not an array")
	}

	// Iterate through each raw transaction
	for _, rawTxItem := range rawTxArray {

		// Each item should be a map containing a "RawTransaction" field
		rawTxWrapper, ok := rawTxItem.(map[string]interface{})
		if !ok {
			return 0, errors.New("RawTransaction item is not an object")
		}

		// Extract the actual transaction from the wrapper
		innerTx, ok := rawTxWrapper["RawTransaction"]
		if !ok {
			return 0, errors.New("RawTransaction field missing from wrapper")
		}

		innerTxMap, ok := innerTx.(map[string]interface{})
		if !ok {
			return 0, errors.New("RawTransaction field is not an object")
		}

		// Convert to FlatTransaction for fee calculation
		flatInnerTx := transaction.FlatTransaction(innerTxMap)

		// Calculate fee for this inner transaction (no multi-signing for inner transactions)
		err := c.calculateFeePerTransactionType(&flatInnerTx, 0)
		if err != nil {

			return 0, err
		}

		// Extract the calculated fee
		feeValue, ok := flatInnerTx["Fee"]
		if !ok {
			return 0, errors.New("fee field missing after calculation")
		}

		feeStr, ok := feeValue.(string)
		if !ok {
			return 0, errors.New("fee field is not a string")
		}

		// Convert fee string to uint64 and add to total
		feeUint, err := strconv.ParseUint(feeStr, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse fee '%s': %w", feeStr, err)
		}

		totalFees += feeUint
	}

	return totalFees, nil
}

// AutofillBatch fills in the missing fields in a batch transaction.
func (c *Client) AutofillBatch(tx *transaction.Batch) error {
	accountSequences := make(map[string]uint32)

	for _, t := range tx.RawTransactions {
		if t.RawTransaction["Sequence"] == nil && t.RawTransaction["TicketSequence"] == nil {
			accountAddr := t.RawTransaction["Account"].(string)

			if _, exists := accountSequences[accountAddr]; exists {
				t.RawTransaction["Sequence"] = accountSequences[accountAddr]
				accountSequences[accountAddr]++

			} else {
				flatTx := transaction.FlatTransaction(t.RawTransaction)
				nextSequence, err := c.getTransactionNextValidSequenceNumber(&flatTx)
				if err != nil {
					return err
				}
				var sequence uint32
				if accountAddr == string(tx.Account) {
					sequence = nextSequence + 1
				} else {
					sequence = nextSequence
				}
				accountSequences[accountAddr] = sequence + 1
				t.RawTransaction["Sequence"] = sequence
			}
		}

		if t.RawTransaction["Fee"] == nil {
			t.RawTransaction["Fee"] = "0"
		} else if t.RawTransaction["Fee"] != "0" {
			return types.ErrBatchInnerTransactionInvalid
		}

		if t.RawTransaction["SigningPubKey"] == nil {
			t.RawTransaction["SigningPubKey"] = ""
		} else if t.RawTransaction["SigningPubKey"] != "" {
			return types.ErrBatchInnerTransactionInvalid
		}

		if t.RawTransaction["TxnSignature"] != nil {
			return types.ErrBatchInnerTransactionInvalid
		}

		if t.RawTransaction["Signers"] != nil {
			return types.ErrBatchNestedTransaction
		}

		if t.RawTransaction["NetworkID"] == nil {
			needsNetworkID := c.NetworkID != 0
			if needsNetworkID {
				t.RawTransaction["NetworkID"] = c.NetworkID
			}
		}

	}
	return nil
}

// getTransactionNextValidSequenceNumber gets the next valid sequence number for a transaction.
func (c *Client) getTransactionNextValidSequenceNumber(tx *transaction.FlatTransaction) (uint32, error) {
	if _, ok := (*tx)["Account"].(string); !ok {
		return 0, errors.New("transaction is missing Account field")
	}

	res, err := c.GetAccountInfo(&account.InfoRequest{
		Account:     types.Address((*tx)["Account"].(string)),
		LedgerIndex: common.LedgerTitle("current"),
	})

	if err != nil {
		return 0, err
	}

	return uint32(res.AccountData.Sequence), nil
}
