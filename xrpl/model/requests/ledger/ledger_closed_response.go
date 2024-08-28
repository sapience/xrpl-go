package ledger

import "github.com/Peersyst/xrpl-go/xrpl/model/requests/common"

type LedgerClosedResponse struct {
	LedgerHash  string             `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
}
