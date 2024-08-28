package transactions

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

// Create or modify a trust line linking two accounts.
type TrustSet struct {
	// Base transaction fields
	BaseTx
	//Object defining the trust line to create or modify, in the format of a Currency Amount.
	LimitAmount types.CurrencyAmount
	// (Optional) Value incoming balances on this trust line at the ratio of this number per 1,000,000,000 units.
	// A value of 0 is shorthand for treating balances at face value. For example, if you set the value to 10,000,000, 1% of incoming funds remain with the sender.
	// If an account sends 100 currency, the sender retains 1 currency unit and the destination receives 99 units. This option is included for parity: in practice, you are much more likely to set a QualityOut value.
	// Note that this fee is separate and independent from token transfer fees.
	QualityIn uint32 `json:",omitempty"`
	// (Optional) Value outgoing balances on this trust line at the ratio of this number per 1,000,000,000 units.
	// A value of 0 is shorthand for treating balances at face value. For example, if you set the value to 10,000,000, 1% of outgoing funds would remain with the issuer.
	// If the sender sends 100 currency units, the issuer retains 1 currency unit and the destination receives 99 units. Note that this fee is separate and independent from token transfer fees.
	QualityOut uint32 `json:",omitempty"`
}

func (*TrustSet) TxType() TxType {
	return TrustSetTx
}

func (t *TrustSet) Flatten() map[string]interface{} {
	flattened := t.BaseTx.Flatten()

	flattened["TransactionType"] = "TrustSet"

	if t.LimitAmount != nil {
		flattened["LimitAmount"] = t.LimitAmount.Flatten()
	}
	if t.QualityIn != 0 {
		flattened["QualityIn"] = t.QualityIn
	}
	if t.QualityOut != 0 {
		flattened["QualityOut"] = t.QualityOut
	}

	return flattened
}

func (t *TrustSet) UnmarshalJSON(data []byte) error {
	type tsHelper struct {
		BaseTx
		LimitAmount json.RawMessage
		QualityIn   uint32 `json:",omitempty"`
		QualityOut  uint32 `json:",omitempty"`
	}
	var h tsHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*t = TrustSet{
		BaseTx:     h.BaseTx,
		QualityIn:  h.QualityIn,
		QualityOut: h.QualityOut,
	}
	limit, err := types.UnmarshalCurrencyAmount(h.LimitAmount)
	if err != nil {
		return err
	}
	t.LimitAmount = limit

	return nil
}
