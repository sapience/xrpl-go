package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type DepositPreauth struct {
	BaseTx
	Authorize   types.Address `json:",omitempty"`
	Unauthorize types.Address `json:",omitempty"`
}

func (*DepositPreauth) TxType() TxType {
	return DepositPreauthTx
}
