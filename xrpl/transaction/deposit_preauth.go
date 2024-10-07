package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type DepositPreauth struct {
	BaseTx
	Authorize   types.Address `json:",omitempty"`
	Unauthorize types.Address `json:",omitempty"`
}

func (*DepositPreauth) TxType() TxType {
	return DepositPreauthTx
}

// TODO: Implement flatten
func (s *DepositPreauth) Flatten() FlatTransaction {
	return nil
}
