package transactions

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/validations"
	"github.com/Peersyst/xrpl-go/xrpl/model/utils"
	typeoffns "github.com/Peersyst/xrpl-go/xrpl/utils/typeof-fns"
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

func (p *Payment) Flatten() map[string]interface{} {
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
func ValidatePayment(tx map[string]interface{}) {
	validations.ValidateBaseTransaction(tx)

	if _, ok := tx["Amount"]; !ok {
		panic("Missing field Amount")
	}

	// IsAmount
	if !utils.IsAmount(tx["Amount"]) {
		panic("Invalid field Amount")
	}

	validations.ValidateRequiredField(tx, "Destination", typeoffns.IsString)
	validations.ValidateOptionalField(tx, "DestinationTag", typeoffns.IsUint32)
	validations.ValidateOptionalField(tx, "InvoiceId", typeoffns.IsString)

	// Check if the field Paths is valid
	if tx["Paths"] != nil {
		if !utils.IsPaths(tx["Paths"].([][]map[string]interface{})) {
			panic("Invalid field Paths")
		}
	}

	validations.ValidateOptionalField(tx, "SendMax", utils.IsAmount)
	validations.ValidateOptionalField(tx, "DeliverMin", utils.IsAmount)
	validations.ValidateOptionalField(tx, "DeliverMax", utils.IsAmount)

	// Check partial payment fields
}
