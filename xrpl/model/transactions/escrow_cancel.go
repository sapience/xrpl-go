package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type EscrowCancel struct {
	BaseTx
	Owner         types.Address
	OfferSequence uint
}

func (*EscrowCancel) TxType() TxType {
	return EscrowCancelTx
}

// TODO: Implement flatten
func (s *EscrowCancel) Flatten() map[string]interface{} {
	return nil
}
