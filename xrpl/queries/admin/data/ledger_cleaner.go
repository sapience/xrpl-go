package data

import "github.com/Peersyst/xrpl-go/xrpl/queries/common"

type LedgerCleanerRequest struct {
	Ledger     common.LedgerIndex `json:"ledger,omitempty"`
	MaxLedger  common.LedgerIndex `json:"max_ledger,omitempty"`
	MinLedger  common.LedgerIndex `json:"min_ledger,omitempty"`
	Full       bool               `json:"full,omitempty"`
	FixTxns    bool               `json:"fix_txns,omitempty"`
	CheckNodes bool               `json:"check_nodes,omitempty"`
	Stop       bool               `json:"stop,omitempty"`
}

func (*LedgerCleanerRequest) Method() string {
	return "ledger_cleaner"
}

type LedgerCleanerResponse struct {
	Message string `json:"message"`
}
