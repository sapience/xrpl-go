package transactions

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type CheckCreate struct {
	BaseTx
	Destination    types.Address
	SendMax        types.CurrencyAmount
	DestinationTag uint          `json:",omitempty"`
	Expiration     uint          `json:",omitempty"`
	InvoiceID      types.Hash256 `json:",omitempty"`
}

func (*CheckCreate) TxType() TxType {
	return CheckCreateTx
}

// TODO: Implement flatten
func (s *CheckCreate) Flatten() FlatTransaction {
	return nil
}

func (c *CheckCreate) UnmarshalJSON(data []byte) error {
	type ccHelper struct {
		BaseTx
		Destination    types.Address
		SendMax        json.RawMessage
		DestinationTag uint          `json:",omitempty"`
		Expiration     uint          `json:",omitempty"`
		InvoiceID      types.Hash256 `json:",omitempty"`
	}
	var h ccHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*c = CheckCreate{
		BaseTx:         h.BaseTx,
		Destination:    h.Destination,
		DestinationTag: h.DestinationTag,
		Expiration:     h.Expiration,
		InvoiceID:      h.InvoiceID,
	}

	max, err := types.UnmarshalCurrencyAmount(h.SendMax)
	if err != nil {
		return err
	}
	c.SendMax = max

	return nil
}
