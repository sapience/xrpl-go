package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
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
