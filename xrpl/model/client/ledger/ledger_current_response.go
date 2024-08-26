package ledger

import "github.com/Peersyst/xrpl-go/xrpl/model/client/common"

type LedgerCurrentResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index"`
}
