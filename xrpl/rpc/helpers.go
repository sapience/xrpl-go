package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/xrpl/currency"
	account "github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	server "github.com/Peersyst/xrpl-go/xrpl/queries/server"
	requests "github.com/Peersyst/xrpl-go/xrpl/queries/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"

	jsoniter "github.com/json-iterator/go"

	commonconstants "github.com/Peersyst/xrpl-go/xrpl/common"
)

const (
	// Sidechains are expected to have network IDs above this.
	// Networks with ID above this restricted number are expected specify an accurate NetworkID field
	// in every transaction to that chain to prevent replay attacks.
	// Mainnet and testnet are exceptions. More context: https://github.com/XRPLF/rippled/pull/4370
	RestrictedNetworks       = 1024
	RequiredNetworkIDVersion = "1.11.0"
)

// isNotLaterRippledVersion determines whether the source rippled version is not later than the target rippled version.
// Example usage: isNotLaterRippledVersion("1.10.0", "1.11.0") returns true.
//
//	isNotLaterRippledVersion("1.10.0", "1.10.0-b1") returns false.
func isNotLaterRippledVersion(source, target string) bool {
	if source == target {
		return true
	}

	sourceDecomp := strings.Split(source, ".")
	targetDecomp := strings.Split(target, ".")

	if len(sourceDecomp) < 3 || len(targetDecomp) < 3 {
		return false
	}

	sourceMajor, err := strconv.Atoi(sourceDecomp[0])
	if err != nil {
		return false
	}
	sourceMinor, err := strconv.Atoi(sourceDecomp[1])
	if err != nil {
		return false
	}
	targetMajor, err := strconv.Atoi(targetDecomp[0])
	if err != nil {
		return false
	}
	targetMinor, err := strconv.Atoi(targetDecomp[1])
	if err != nil {
		return false
	}

	// Compare major version
	if sourceMajor != targetMajor {
		return sourceMajor < targetMajor
	}

	// Compare minor version
	if sourceMinor != targetMinor {
		return sourceMinor < targetMinor
	}

	sourcePatch := strings.Split(sourceDecomp[2], "-")
	targetPatch := strings.Split(targetDecomp[2], "-")

	sourcePatchVersion, err := strconv.Atoi(sourcePatch[0])
	if err != nil {
		return false
	}
	targetPatchVersion, err := strconv.Atoi(targetPatch[0])
	if err != nil {
		return false
	}

	// Compare patch version
	if sourcePatchVersion != targetPatchVersion {
		return sourcePatchVersion < targetPatchVersion
	}

	// Compare release version
	if len(sourcePatch) != len(targetPatch) {
		return len(sourcePatch) > len(targetPatch)
	}

	if len(sourcePatch) == 2 {
		// Compare different release types
		if !strings.HasPrefix(sourcePatch[1], string(targetPatch[1][0])) {
			return sourcePatch[1] < targetPatch[1]
		}

		// Compare beta version
		if strings.HasPrefix(sourcePatch[1], "b") {
			sourceBeta, err := strconv.Atoi(sourcePatch[1][1:])
			if err != nil {
				return false
			}
			targetBeta, err := strconv.Atoi(targetPatch[1][1:])
			if err != nil {
				return false
			}
			return sourceBeta < targetBeta
		}

		// Compare rc version
		if strings.HasPrefix(sourcePatch[1], "rc") {
			sourceRC, err := strconv.Atoi(sourcePatch[1][2:])
			if err != nil {
				return false
			}
			targetRC, err := strconv.Atoi(targetPatch[1][2:])
			if err != nil {
				return false
			}
			return sourceRC < targetRC
		}
	}

	return false
}

// txNeedsNetworkID determines if the transaction required a networkID to be valid.
// Transaction needs networkID if later than restricted ID and build version is >= 1.11.0
func (c *Client) txNeedsNetworkID() (bool, error) {
	if c.NetworkID != 0 && c.NetworkID > RestrictedNetworks {
		res, err := c.GetServerInfo(&server.InfoRequest{})
		if err != nil {
			return false, err
		}

		if res.Info.BuildVersion != "" {
			return isNotLaterRippledVersion(RequiredNetworkIDVersion, res.Info.BuildVersion), nil
		}
	}
	return false, nil
}

// CreateRequest formats the parameters and method name ready for sending request
// Params will have been serialised if required and added to request struct before being passed to this method
func createRequest(reqParams XRPLRequest) ([]byte, error) {
	var body Request

	reqParams.SetAPIVersion(
		reqParams.APIVersion(),
	)

	body = Request{
		Method: reqParams.Method(),
		// each param object will have a struct with json serialising tags
		Params: [1]interface{}{reqParams},
	}

	// Omit the Params field if method doesn't require any
	paramBytes, err := jsoniter.Marshal(body.Params)
	if err != nil {
		return nil, err
	}
	paramString := string(paramBytes)
	if strings.Compare(paramString, "[{}]") == 0 {
		// need to remove params field from the body if it is empty
		body = Request{
			Method: reqParams.Method(),
		}

		jsonBytes, err := jsoniter.Marshal(body)
		if err != nil {
			return nil, err
		}

		return jsonBytes, nil
	}

	jsonBytes, err := jsoniter.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON-RPC request for method %s with parameters %+v: %w", reqParams.Method(), reqParams, err)
	}

	return jsonBytes, nil
}

// checkForError reads the http response and formats the error if it exists
func checkForError(res *http.Response) (Response, error) {

	var jr Response

	b, err := io.ReadAll(res.Body)
	if err != nil || b == nil {
		return jr, err
	}

	// In case a different error code is returned
	if res.StatusCode != 200 {
		return jr, &ClientError{ErrorString: string(b)}
	}

	jDec := json.NewDecoder(bytes.NewReader(b))
	jDec.UseNumber()
	err = jDec.Decode(&jr)
	if err != nil {
		return jr, err
	}

	// result will have 'error' if error response
	if _, ok := jr.Result["error"]; ok {
		return jr, &ClientError{ErrorString: jr.Result["error"].(string)}
	}

	return jr, nil
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

// Sets the next valid sequence number for a given transaction.
func (c *Client) setTransactionNextValidSequenceNumber(tx *transaction.FlatTransaction) error {
	seqNumber, err := c.getTransactionNextValidSequenceNumber(tx)
	if err != nil {
		return err
	}
	(*tx)["Sequence"] = seqNumber
	return nil
}

func (c *Client) getTransactionNextValidSequenceNumber(tx *transaction.FlatTransaction) (uint32, error) {
	if _, ok := (*tx)["Account"].(string); !ok {
		return 0, errors.New("missing Account in transaction")
	}

	res, err := c.GetAccountInfo(&account.InfoRequest{
		Account:     types.Address((*tx)["Account"].(string)),
		LedgerIndex: common.LedgerTitle("current"),
	})

	if err != nil {
		return 0, err
	}

	return uint32(res.AccountData.TicketCount), nil
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
// TODO: Add fee support for `EscrowFinish` `AccountDelete`, `AMMCreate`, and multisigned transactions.
func (c *Client) calculateFeePerTransactionType(tx *transaction.FlatTransaction, nSigners uint64) error {
	fee, err := c.getFeeXrp(c.cfg.feeCushion)
	if err != nil {
		return err
	}

	feeDrops, err := currency.XrpToDrops(fee)
	if err != nil {
		return err
	}

	if nSigners > 0 {
		// Convert feeDrops to uint64 for safe arithmetic
		baseFee, err := strconv.ParseUint(feeDrops, 10, 64)
		if err != nil {
			return err
		}

		// Calculate total signers fee: fee * nSigners
		signersFee := baseFee * nSigners

		// Add base fee and signers fee
		totalFee := baseFee + signersFee

		// Convert back to string
		feeDrops = strconv.FormatUint(totalFee, 10)
	}

	(*tx)["Fee"] = feeDrops

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

// getSignedTx ensures the transaction is fully signed and returns the transaction blob.
// If the transaction is already signed, it encodes and returns it. Otherwise, it autofills (if enabled)
// and signs the transaction using the provided wallet.
func (c *Client) getSignedTx(tx transaction.FlatTransaction, autofill bool, wallet *wallet.Wallet) (string, error) {
	// Check if the transaction is already signed: both fields must be non-empty.
	sig, sigOk := tx["TxnSignature"].(string)
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
