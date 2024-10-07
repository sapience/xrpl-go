package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type CheckCancel struct {
	BaseTx
	CheckID types.Hash256
}

// TODO: Implement flatten
func (*CheckCancel) TxType() TxType {
	return CheckCancelTx
}

func (s *CheckCancel) Flatten() FlatTransaction {
	return nil
}
