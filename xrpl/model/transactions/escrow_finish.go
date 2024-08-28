package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type EscrowFinish struct {
	BaseTx
	Owner         types.Address
	OfferSequence uint
	Condition     string `json:",omitempty"`
	Fulfillment   string `json:",omitempty"`
}

func (*EscrowFinish) TxType() TxType {
	return EscrowFinishTx
}

// TODO: Implement flatten
func (s *EscrowFinish) Flatten() map[string]interface{} {
	return nil
}
