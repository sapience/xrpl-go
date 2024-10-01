package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
)

type TransactionEntryResponse struct {
	LedgerIndex common.LedgerIndex           `json:"ledger_index"`
	LedgerHash  common.LedgerHash            `json:"ledger_hash,omitempty"`
	Metadata    transactions.TxObjMeta       `json:"metadata"`
	Tx          transactions.FlatTransaction `json:"tx_json"`
}
