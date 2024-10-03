package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
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
func (s *EscrowCancel) Flatten() FlatTransaction {
	return nil
}
