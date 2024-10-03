package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type TxResponse struct {
	Date        uint                         `json:"date"`
	Hash        types.Hash256                `json:"hash"`
	LedgerIndex common.LedgerIndex           `json:"ledger_index"`
	Meta        transactions.TxMeta          `json:"meta"`
	Validated   bool                         `json:"validated"`
	Tx          transactions.FlatTransaction `json:",omitempty"`
}
