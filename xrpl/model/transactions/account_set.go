package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type AccountSet struct {
	BaseTx
	ClearFlag     uint          `json:",omitempty"`
	Domain        string        `json:",omitempty"`
	EmailHash     types.Hash128 `json:",omitempty"`
	MessageKey    string        `json:",omitempty"`
	NFTokenMinter string        `json:",omitempty"`
	SetFlag       uint          `json:",omitempty"`
	TransferRate  uint          `json:",omitempty"`
	TickSize      uint8         `json:",omitempty"`
	WalletLocator types.Hash256 `json:",omitempty"`
	WalletSize    uint          `json:",omitempty"`
}

func (*AccountSet) TxType() TxType {
	return AccountSetTx
}

// TODO: Implement flatten
func (s *AccountSet) Flatten() map[string]interface{} {
	return nil
}
