package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type AccountDelete struct {
	BaseTx
	Destination    types.Address
	DestinationTag uint32 `json:",omitempty"`
}

func (*AccountDelete) TxType() TxType {
	return AccountDeleteTx
}

// TODO: Implement flatten
func (s *AccountDelete) Flatten() FlatTransaction {
	return nil
}
