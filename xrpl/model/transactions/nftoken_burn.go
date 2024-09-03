package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type NFTokenBurn struct {
	BaseTx
	NFTokenID types.NFTokenID
	Owner     types.Address `json:",omitempty"`
}

func (*NFTokenBurn) TxType() TxType {
	return NFTokenBurnTx
}

// TODO: Implement flatten
func (s *NFTokenBurn) Flatten() FlatTransaction {
	return nil
}
