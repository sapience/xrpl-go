package ledger

import "github.com/Peersyst/xrpl-go/xrpl/model/client/common"

type LedgerClosedResponse struct {
	LedgerHash  string             `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
}
