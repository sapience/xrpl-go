package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/Peersyst/xrpl-go/xrpl/client"
	"github.com/Peersyst/xrpl-go/xrpl/model/client/account"
	"github.com/Peersyst/xrpl-go/xrpl/model/client/common"
	"github.com/Peersyst/xrpl-go/xrpl/model/client/server"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/utils"
	"github.com/mitchellh/mapstructure"
)

func (c *WebsocketClient) formatRequest(req client.XRPLRequest, id int, marker any) ([]byte, error) {
	m := make(map[string]any)
	m["id"] = id
	m["command"] = req.Method()
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
func (c *WebsocketClient) getClassicAccountAndTag(address string) (string, uint32) {
	return address, 0
}

func (c *WebsocketClient) convertTransactionAddressToClassicAddress(tx *map[string]interface{}, fieldName string) {
	if address, ok := (*tx)[fieldName].(string); ok {
		classicAddress, _ := c.getClassicAccountAndTag(address)
		(*tx)[fieldName] = classicAddress
	}
}

func (c *WebsocketClient) validateTransactionAddress(tx *map[string]interface{}, addressField, tagField string) error {
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
func (c *WebsocketClient) setValidTransactionAddresses(tx *map[string]interface{}) error {
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
func (c *WebsocketClient) setTransactionNextValidSequenceNumber(tx *map[string]interface{}) error {
	if _, ok := (*tx)["Account"].(string); !ok {
		return errors.New("missing Account in transaction")
	}
	res, _, err := c.GetAccountInfo(&account.AccountInfoRequest{
		Account:     types.Address((*tx)["Account"].(string)),
		LedgerIndex: common.LedgerTitle("current"),
	})

	if err != nil {
		return err
	}

	(*tx)["Sequence"] = res.AccountData.Sequence
	return nil
}

// Calculates the current transaction fee for the ledger.
// Note: This is a public API that can be called directly.
func (c *WebsocketClient) getFeeXrp(cushion float32) (string, error) {
	res, err := c.GetServerInfo(&server.ServerInfoRequest{})
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

	if fee > c.MaxFeeXRP {
		fee = c.MaxFeeXRP
	}

	// Round fee to NUM_DECIMAL_PLACES
	roundedFee := float32(math.Round(float64(fee)*math.Pow10(client.NUM_DECIMAL_PLACES))) / float32(math.Pow10(client.NUM_DECIMAL_PLACES))

	// Convert the rounded fee back to a string with NUM_DECIMAL_PLACES
	return fmt.Sprintf("%.*f", client.NUM_DECIMAL_PLACES, roundedFee), nil
}

// Calculates the fee per transaction type.
//
// TODO: Add fee support for `EscrowFinish` `AccountDelete`, `AMMCreate`, and multisigned transactions.
func (c *WebsocketClient) calculateFeePerTransactionType(tx *map[string]interface{}) error {
	fee, err := c.getFeeXrp(c.FeeCushion)
	if err != nil {
		return err
	}

	// TODO: Replace with XrpToDrops util when implemented
	feeFloat, err := strconv.ParseFloat(fee, 64)
	if err != nil {
		return err
	}
	feeDrops := feeFloat * utils.DROPS_PER_XRP

	roundedFee := math.Ceil(feeDrops)
	(*tx)["Fee"] = fmt.Sprintf("%.0f", roundedFee)

	return nil
}

// Sets the latest validated ledger sequence for the transaction.
// Modifies the `LastLedgerSequence` field in the tx.
func (c *WebsocketClient) setLastLedgerSequence(tx *map[string]interface{}) error {
	index, err := c.GetLedgerIndex()
	if err != nil {
		return err
	}

	(*tx)["LastLedgerSequence"] = index.Uint32() + LEDGER_OFFSET
	return err
}

// Checks for any blockers that prevent the deletion of an account.
// Returns nil if there are no blockers, otherwise returns an error.
func (c *WebsocketClient) checkAccountDeleteBlockers(address types.Address) error {
	accObjects, err := c.GetAccountObjects(&account.AccountObjectsRequest{
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

func (c *WebsocketClient) checkPaymentAmounts(tx *map[string]interface{}) error {
	if _, ok := (*tx)["DeliverMax"]; ok {
		if _, ok := (*tx)["Amount"]; !ok {
			(*tx)["Amount"] = (*tx)["DeliverMax"]
		} else {
			if (*tx)["Amount"] != (*tx)["DeliverMax"] {
				return errors.New("payment transaction: Amount and DeliverMax fields must be identical when both are provided")
			}
		}
	}
	return nil
}

// Sets a transaction's flags to its numeric representation.
// TODO: Add flag support for AccountSet, AMMDeposit, AMMWithdraw,
// NFTTOkenCreateOffer, NFTokenMint, OfferCreate, XChainModifyBridge (not supported).
func (c *WebsocketClient) setTransactionFlags(tx *map[string]interface{}) error {
	flags, ok := (*tx)["Flags"].(uint32)
	if !ok && flags > 0 {
		(*tx)["Flags"] = uint32(0)
		return nil
	}

	_, ok = (*tx)["TransactionType"].(string)
	if !ok {
		return errors.New("transaction type is missing in transaction")
	}

	return nil
}
