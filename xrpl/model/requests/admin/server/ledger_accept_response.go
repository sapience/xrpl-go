package server

import "github.com/Peersyst/xrpl-go/xrpl/model/requests/common"

type LedgerAcceptResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index"`
}
