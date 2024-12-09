package server

import "github.com/Peersyst/xrpl-go/xrpl/queries/common"

type LedgerAcceptRequest struct {
}

func (*LedgerAcceptRequest) Method() string {
	return "leder_accept"
}

type LedgerAcceptResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index"`
}
