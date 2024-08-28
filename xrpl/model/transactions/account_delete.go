package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type AccountDelete struct {
	BaseTx
	Destination    types.Address
	DestinationTag uint `json:",omitempty"`
}

func (*AccountDelete) TxType() TxType {
	return AccountDeleteTx
}

// TODO: Implement flatten
func (s *AccountDelete) Flatten() map[string]interface{} {
	return nil
}
