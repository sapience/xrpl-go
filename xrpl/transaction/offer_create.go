package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// An OfferCreate transaction places an Offer in the decentralized exchange.
//
// Example:
//
// ```json
//
//	{
//	    "TransactionType": "OfferCreate",
//	    "Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
//	    "Fee": "12",
//	    "Flags": 0,
//	    "LastLedgerSequence": 7108682,
//	    "Sequence": 8,
//	    "TakerGets": "6000000",
//	    "TakerPays": {
//	      "currency": "GKO",
//	      "issuer": "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
//	      "value": "2"
//	    }
//	}
//
// ```
type OfferCreate struct {
	BaseTx
	// (Optional) Time after which the Offer is no longer active, in seconds since the Ripple Epoch.
	Expiration uint `json:",omitempty"`
	// (Optional) An Offer to delete first, specified in the same way as OfferCancel.
	OfferSequence uint `json:",omitempty"`
	// The amount and type of currency being sold.
	TakerGets types.CurrencyAmount
	// The amount and type of currency being bought.
	TakerPays types.CurrencyAmount
}

// TxType returns the type of the transaction (OfferCreate).
func (*OfferCreate) TxType() TxType {
	return OfferCreateTx
}

// Flatten returns a map of the OfferCreate transaction fields.
func (s *OfferCreate) Flatten() FlatTransaction {
	flattened := s.BaseTx.Flatten()

	if s.Expiration != 0 {
		flattened["Expiration"] = s.Expiration
	}
	if s.OfferSequence != 0 {
		flattened["OfferSequence"] = s.OfferSequence
	}
	flattened["TakerGets"] = s.TakerGets.Flatten()
	flattened["TakerPays"] = s.TakerPays.Flatten()

	return flattened
}

// Validates the OfferCreate transaction.
func (o *OfferCreate) Validate() (bool, error) {
	_, err := o.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if ok, err := IsAmount(o.TakerGets, "TakerGets", true); !ok {
		return false, err
	}

	if ok, err := IsAmount(o.TakerPays, "TakerPays", true); !ok {
		return false, err
	}

	return true, nil
}
