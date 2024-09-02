package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
)

const (
	ErrAccountTxUnmarshal string = "Unmarshal JSON AccountTransaction"
)

type AccountTransaction struct {
	LedgerIndex uint64                       `json:"ledger_index"`
	Meta        transactions.TxMeta          `json:"meta"`
	Tx          transactions.FlatTransaction `json:"tx"`
	TxBlob      string                       `json:"tx_blob"`
	Validated   bool                         `json:"validated"`
}
