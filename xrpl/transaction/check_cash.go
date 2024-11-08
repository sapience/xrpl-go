package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type CheckCash struct {
	BaseTx
	CheckID    types.Hash256
	Amount     types.CurrencyAmount `json:",omitempty"`
	DeliverMin types.CurrencyAmount `json:",omitempty"`
}

func (*CheckCash) TxType() TxType {
	return CheckCashTx
}

// TODO: Implement flatten
func (s *CheckCash) Flatten() FlatTransaction {
	return nil
}
