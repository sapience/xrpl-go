package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type CheckCancel struct {
	BaseTx
	CheckID types.Hash256
}

func (*CheckCancel) TxType() TxType {
	return CheckCancelTx
}
