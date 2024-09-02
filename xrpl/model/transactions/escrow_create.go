package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type EscrowCreate struct {
	BaseTx
	Amount         types.XRPCurrencyAmount
	Destination    types.Address
	CancelAfter    uint   `json:",omitempty"`
	FinishAfter    uint   `json:",omitempty"`
	Condition      string `json:",omitempty"`
	DestinationTag uint   `json:",omitempty"`
}

func (*EscrowCreate) TxType() TxType {
	return EscrowCreateTx
}

// TODO: Implement flatten
func (s *EscrowCreate) Flatten() FlatTransaction {
	return nil
}
