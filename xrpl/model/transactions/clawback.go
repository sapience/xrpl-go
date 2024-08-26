package transactions

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type Clawback struct {
	BaseTx

	// The amount of currency to deliver, and it must be non-XRP. The nested field
	// names MUST be lower-case. The `issuer` field MUST be the holder's address,
	// whom to be clawed back.
	Amount types.CurrencyAmount
}

func (*Clawback) TxType() TxType {
	return ClawbackTx
}

// UnmarshalJSON unmarshals the JSON data into a Clawback struct.
func (c *Clawback) UnmarshalJSON(data []byte) error {
	// Define a helper struct to hold the unmarshaled data
	type cHelper struct {
		BaseTx
		Amount json.RawMessage
	}

	var h cHelper

	// Unmarshal the JSON data into the helper struct
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}

	// Assign the values from the helper struct to the Clawback struct
	*c = Clawback{
		BaseTx: h.BaseTx,
	}

	// Unmarshal the Amount field into a CurrencyAmount struct
	var amount types.CurrencyAmount
	amount, err = types.UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return err
	}

	// Assign the unmarshaled CurrencyAmount to the Clawback struct
	c.Amount = amount

	return nil
}
