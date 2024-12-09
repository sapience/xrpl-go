package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type TxRequest struct {
	Transaction string             `json:"transaction"`
	Binary      bool               `json:"binary,omitempty"`
	MinLedger   common.LedgerIndex `json:"min_ledger,omitempty"`
	MaxLedger   common.LedgerIndex `json:"max_ledger,omitempty"`
}

func (*TxRequest) Method() string {
	return "tx"
}

type TxResponse struct {
	Date        uint                         `json:"date"`
	Hash        types.Hash256                `json:"hash"`
	LedgerIndex common.LedgerIndex           `json:"ledger_index"`
	Meta        transaction.TxMeta          `json:"meta"`
	Validated   bool                         `json:"validated"`
	Tx          transaction.FlatTransaction `json:",omitempty"`
}
