package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
)

// ############################################################################
// Request
// ############################################################################

type EntryRequest struct {
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	TxHash      string                 `json:"tx_hash"`
}

func (*EntryRequest) Method() string {
	return "transaction_entry"
}

// ############################################################################
// Response
// ############################################################################

type EntryResponse struct {
	LedgerIndex common.LedgerIndex          `json:"ledger_index"`
	LedgerHash  common.LedgerHash           `json:"ledger_hash,omitempty"`
	Metadata    transaction.TxObjMeta       `json:"metadata"`
	Tx          transaction.FlatTransaction `json:"tx_json"`
}
