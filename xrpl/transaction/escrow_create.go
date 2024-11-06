package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type EscrowCreate struct {
	BaseTx
	Amount         types.XRPCurrencyAmount
	Destination    types.Address
	CancelAfter    uint32 `json:",omitempty"`
	FinishAfter    uint32 `json:",omitempty"`
	Condition      string `json:",omitempty"`
	DestinationTag uint32 `json:",omitempty"`
}

func (*EscrowCreate) TxType() TxType {
	return EscrowCreateTx
}

// TODO: Implement flatten
func (s *EscrowCreate) Flatten() FlatTransaction {
	return nil
}
