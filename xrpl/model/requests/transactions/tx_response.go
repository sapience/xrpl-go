package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type TxResponse struct {
	Date        uint                         `json:"date"`
	Hash        types.Hash256                `json:"hash"`
	LedgerIndex common.LedgerIndex           `json:"ledger_index"`
	Meta        transactions.TxMeta          `json:"meta"`
	Validated   bool                         `json:"validated"`
	Tx          transactions.FlatTransaction `json:",omitempty"`
}
