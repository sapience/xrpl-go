package transactions

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type TrustSet struct {
	BaseTx
	LimitAmount types.CurrencyAmount
	QualityIn   uint `json:",omitempty"`
	QualityOut  uint `json:",omitempty"`
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
		QualityIn   uint `json:",omitempty"`
		QualityOut  uint `json:",omitempty"`
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
