package transactions

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/transactions/types"
)

type Payment struct {
	BaseTx
	/*
		API v1: Only available in API v1.
		The maximum amount of currency to deliver.
		For non-XRP amounts, the nested field names MUST be lower-case.
		If the tfPartialPayment flag is set, deliver up to this amount instead.
	*/
	Amount types.CurrencyAmount
	/*
		API v2: Only available in API v2.
		The maximum amount of currency to deliver.
		For non-XRP amounts, the nested field names MUST be lower-case.
		If the tfPartialPayment flag is set, deliver up to this amount instead.
	*/
	DeliverMax types.CurrencyAmount `json:",omitempty"`
	/*
		(Optional) Minimum amount of destination currency this transaction should deliver.
		Only valid if this is a partial payment.
		For non-XRP amounts, the nested field names are lower-case.
	*/
	DeliverMin types.CurrencyAmount `json:",omitempty"`
	/*
		The unique address of the account receiving the payment.
	*/
	Destination types.Address
	/*
		(Optional) Arbitrary tag that identifies the reason for the payment to the destination, or a hosted recipient to pay.
	*/
	DestinationTag uint `json:",omitempty"`
	/*
		(Optional) Arbitrary 256-bit hash representing a specific reason or identifier for this payment
	*/
	InvoiceID uint `json:",omitempty"`
	/*
		(Optional, auto-fillable) Array of payment paths to be used for this transaction.
		Must be omitted for XRP-to-XRP transactions.
	*/
	Paths [][]PathStep `json:",omitempty"`
	/*
		(Optional) Highest amount of source currency this transaction is allowed to cost,
		including transfer fees, exchange rates, and slippage.
		Does not include the XRP destroyed as a cost for submitting the transaction.
		For non-XRP amounts, the nested field names MUST be lower-case.
		Must be supplied for cross-currency/cross-issue payments.
		Must be omitted for XRP-to-XRP payments.
	*/
	SendMax types.CurrencyAmount `json:",omitempty"`
}

// "New" creates a new Payment object based on the provided Payment struct.
// It validates the different fields of the payment.
// If no validation error is found, it initializes a new Payment object with the provided fields and returns it.
func New(payment Payment) Payment {
	// Validate the BaseTx field of the payment
	ValidateBaseTx(&payment.BaseTx)

	// TODO: validate other required and optional fields

	return Payment{
		BaseTx:         payment.BaseTx,
		Amount:         payment.Amount,
		Destination:    payment.Destination,
		DestinationTag: payment.DestinationTag,
		InvoiceID:      payment.InvoiceID,
		Paths:          payment.Paths,
		SendMax:        payment.SendMax,
		DeliverMin:     payment.DeliverMin,
	}
}

func (*Payment) TxType() TxType {
	return PaymentTx
}

func (p *Payment) UnmarshalJSON(data []byte) error {
	type pHelper struct {
		BaseTx
		Amount         json.RawMessage
		Destination    types.Address
		DestinationTag uint            `json:",omitempty"`
		InvoiceID      uint            `json:",omitempty"`
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
