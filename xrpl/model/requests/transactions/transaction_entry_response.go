package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
)

type TransactionEntryResponse struct {
	LedgerIndex common.LedgerIndex           `json:"ledger_index"`
	LedgerHash  common.LedgerHash            `json:"ledger_hash,omitempty"`
	Metadata    transactions.TxObjMeta       `json:"metadata"`
	Tx          transactions.FlatTransaction `json:"tx_json"`
}
