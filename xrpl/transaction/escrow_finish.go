package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// Deliver XRP from a held payment to the recipient.
//
// Example:
//
// ```json
//
//	{
//	    "Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
//	    "TransactionType": "EscrowFinish",
//	    "Owner": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
//	    "OfferSequence": 7,
//	    "Condition": "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
//	    "Fulfillment": "A0028000"
//	}
//
// ````
type EscrowFinish struct {
	BaseTx
	// Address of the source account that funded the held payment.
	Owner types.Address
	// Transaction sequence of EscrowCreate transaction that created the held payment to finish.
	OfferSequence uint
	// (Optional) Hex value matching the previously-supplied PREIMAGE-SHA-256 crypto-condition of the held payment.
	Condition string `json:",omitempty"`
	// Optional) Hex value of the PREIMAGE-SHA-256 crypto-condition fulfillment matching the held payment's Condition.
	Fulfillment string `json:",omitempty"`
}

// TxType returns the transaction type for this transaction (EscrowFinish).
func (*EscrowFinish) TxType() TxType {
	return EscrowFinishTx
}

// Flatten returns the flattened map of the EscrowFinish transaction.
func (e *EscrowFinish) Flatten() FlatTransaction {
	flattened := e.BaseTx.Flatten()

	flattened["TransactionType"] = "EscrowFinish"

	if e.Owner != "" {
		flattened["Owner"] = e.Owner
	}

	if e.OfferSequence != 0 {
		flattened["OfferSequence"] = e.OfferSequence
	}

	if e.Condition != "" {
		flattened["Condition"] = e.Condition
	}

	if e.Fulfillment != "" {
		flattened["Fulfillment"] = e.Fulfillment
	}

	return flattened
}

// Validate checks if the EscrowFinish struct is valid.
func (e *EscrowFinish) Validate() (bool, error) {
	ok, err := e.BaseTx.Validate()
	if err != nil || !ok {
		return false, err
	}

	if !addresscodec.IsValidClassicAddress(e.Owner.String()) {
		return false, errors.New("invalid xrpl address for the Owner field")
	}

	if e.OfferSequence == 0 {
		return false, errors.New("missing OfferSequence")
	}

	return true, nil
}
