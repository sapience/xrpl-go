package transactions

import (
	"encoding/json"
	"errors"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/model/utils"
	"github.com/Peersyst/xrpl-go/xrpl/utils/typecheck"
)

const (
	// Do not use the default path; only use paths included in the Paths field.
	// This is intended to force the transaction to take arbitrage opportunities.
	// Most clients do not need this.
	tfRippleNotDirect uint = 65536
	// If the specified Amount cannot be sent without spending more than SendMax,
	// reduce the received amount instead of failing outright. See Partial
	// Payments for more details.
	tfPartialPayment uint = 131072
	// Only take paths where all the conversions have an input:output ratio that
	// is equal or better than the ratio of Amount:SendMax. See Limit Quality for
	// details.
	tfLimitQuality uint = 262144
)

// A Payment transaction represents a transfer of value from one account to another.
type Payment struct {
	BaseTx
	// API v1: Only available in API v1.
	// The maximum amount of currency to deliver.
	// For non-XRP amounts, the nested field names MUST be lower-case.
	// If the tfPartialPayment flag is set, deliver up to this amount instead.
	Amount types.CurrencyAmount

	// API v2: Only available in API v2.
	// The maximum amount of currency to deliver.
	// For non-XRP amounts, the nested field names MUST be lower-case.
	// If the tfPartialPayment flag is set, deliver up to this amount instead.
	DeliverMax types.CurrencyAmount `json:",omitempty"`

	// (Optional) Minimum amount of destination currency this transaction should deliver.
	// Only valid if this is a partial payment.
	// For non-XRP amounts, the nested field names are lower-case.
	DeliverMin types.CurrencyAmount `json:",omitempty"`

	// The unique address of the account receiving the payment.
	Destination types.Address

	// (Optional) Arbitrary tag that identifies the reason for the payment to the destination, or a hosted recipient to pay.
	DestinationTag uint32 `json:",omitempty"`

	// (Optional) Arbitrary 256-bit hash representing a specific reason or identifier for this payment
	InvoiceID types.Hash256 `json:",omitempty"`

	// (Optional, auto-fillable) Array of payment paths to be used for this transaction.
	// Must be omitted for XRP-to-XRP transactions.
	Paths [][]PathStep `json:",omitempty"`

	// (Optional) Highest amount of source currency this transaction is allowed to cost,
	// including transfer fees, exchange rates, and slippage.
	// Does not include the XRP destroyed as a cost for submitting the transaction.
	// For non-XRP amounts, the nested field names MUST be lower-case.
	// Must be supplied for cross-currency/cross-issue payments.
	// Must be omitted for XRP-to-XRP payments.
	SendMax types.CurrencyAmount `json:",omitempty"`
}

// TxType returns the type of the transaction (Payment).
func (Payment) TxType() TxType {
	return PaymentTx
}

// Flatten returns the flattened map of the Payment transaction.
func (p *Payment) Flatten() FlatTransaction {
	// Add BaseTx fields
	flattened := p.BaseTx.Flatten()

	// Add Payment-specific fields
	flattened["TransactionType"] = "Payment"

	if p.Amount != nil {
		flattened["Amount"] = p.Amount.Flatten()
	}

	if p.DeliverMax != nil {
		flattened["DeliverMax"] = p.DeliverMax.Flatten()
	}

	if p.DeliverMin != nil {
		flattened["DeliverMin"] = p.DeliverMin.Flatten()
	}

	if p.Destination != "" {
		flattened["Destination"] = p.Destination.String()
	}

	if p.DestinationTag != 0 {
		flattened["DestinationTag"] = p.DestinationTag
	}

	if p.InvoiceID != "" {
		flattened["InvoiceID"] = p.InvoiceID.String()
	}

	if len(p.Paths) > 0 {
		flattened["Paths"] = p.Paths
	}

	if p.SendMax != nil {
		flattened["SendMax"] = p.SendMax.Flatten()
	}

	return flattened
}

// SetRippleNotDirectFlag sets the RippleNotDirect flag.
//
// RippleNotDirect: Do not use the default path; only use paths included in the Paths field.
// This is intended to force the transaction to take arbitrage opportunities.
// Most clients do not need this.
func (p *Payment) SetRippleNotDirectFlag() {
	p.Flags |= tfRippleNotDirect
}

// SetPartialPaymentFlag sets the PartialPayment flag.
//
// PartialPayment: If the specified Amount cannot be sent without spending more than SendMax,
// reduce the received amount instead of failing outright. See Partial
// Payments for more details.
func (p *Payment) SetPartialPaymentFlag() {
	p.Flags |= tfPartialPayment
}

// SetLimitQualityFlag sets the LimitQuality flag.
//
// LimitQuality: Only take paths where all the conversions have an input:output ratio that
// is equal or better than the ratio of Amount:SendMax. See Limit Quality for
// details.
func (p *Payment) SetLimitQualityFlag() {
	p.Flags |= tfLimitQuality
}

// UnmarshalJSON unmarshals the Payment transaction from JSON.
func (p *Payment) UnmarshalJSON(data []byte) error {
	type pHelper struct {
		BaseTx
		Amount         json.RawMessage
		Destination    types.Address
		DestinationTag uint32          `json:",omitempty"`
		InvoiceID      types.Hash256   `json:",omitempty"`
		Paths          [][]PathStep    `json:",omitempty"`
		SendMax        json.RawMessage `json:",omitempty"`
		DeliverMin     json.RawMessage `json:",omitempty"`
	}
	var h pHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*p = Payment{
		BaseTx:         h.BaseTx,
		Destination:    h.Destination,
		DestinationTag: h.DestinationTag,
		InvoiceID:      h.InvoiceID,
		Paths:          h.Paths,
	}
	var amount, max, min types.CurrencyAmount
	var err error
	amount, err = types.UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return err
	}
	max, err = types.UnmarshalCurrencyAmount(h.SendMax)
	if err != nil {
		return err
	}
	min, err = types.UnmarshalCurrencyAmount(h.DeliverMin)
	if err != nil {
		return err
	}
	p.Amount = amount
	p.DeliverMin = min
	p.SendMax = max

	return nil
}

// ValidatePayment validates the Payment struct and make sure all the fields are correct.
func ValidatePayment(tx FlatTransaction) error {
	ValidateBaseTransaction(tx)

	// Check if the field Amount is set
	if _, ok := tx["Amount"]; !ok {
		return errors.New("missing field Amount")
	}

	// Check if the field Amount is valid
	if !IsAmount(tx["Amount"]) {
		return errors.New("invalid field Amount")
	}

	var err error

	// Check if the field Destination is set and valid
	err = ValidateRequiredField(tx, "Destination", typecheck.IsString)
	if err != nil {
		return err
	}

	// Check if the field DestinationTag is valid
	err = ValidateOptionalField(tx, "DestinationTag", typecheck.IsUint32)
	if err != nil {
		return err
	}

	// Check if the field InvoiceId is valid
	err = ValidateOptionalField(tx, "InvoiceId", typecheck.IsString)
	if err != nil {
		return err
	}

	// Check if the field Paths is valid
	if tx["Paths"] != nil {
		if !IsPaths(tx["Paths"].([][]map[string]interface{})) {
			return errors.New("invalid field Paths")
		}
	}

	// Check if the field SendMax is valid
	err = ValidateOptionalField(tx, "SendMax", IsAmount)
	if err != nil {
		return err
	}

	// Check if the field DeliverMax is valid
	err = ValidateOptionalField(tx, "DeliverMax", IsAmount)
	if err != nil {
		return err
	}

	// Check if the field DeliverMin is valid
	err = ValidateOptionalField(tx, "DeliverMin", IsAmount)
	if err != nil {
		return err
	}

	// Check partial payment fields
	err = checkPartialPayment(tx)
	if err != nil {
		return err
	}

	return nil
}

func checkPartialPayment(tx FlatTransaction) error {
	// Check if DeliverMin is set
	if _, ok := tx["DeliverMin"]; ok {
		// Check if Flags flag is set
		if _, ok := tx["Flags"]; !ok {
			return errors.New("Payment transaction: tfPartialPayment flag required with DeliverMin")
		}

		// Extract Flags
		flagsField, isUint := (tx["Flags"]).(uint)
		var isTfPartialPayment bool
		if isUint {
			isTfPartialPayment = utils.IsFlagEnabled(flagsField, uint(tfPartialPayment))
		}

		// TODO: check if tfPartialPayment is enabled if/when Flags is an object/map instead of a uint

		if !isTfPartialPayment {
			return errors.New("Payment transaction: tfPartialPayment flag required with DeliverMin")
		}

		err := ValidateOptionalField(tx, "DeliverMin", IsAmount)
		if err != nil {
			return err
		}
	}

	return nil
}
