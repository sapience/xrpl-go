package data

import "github.com/Peersyst/xrpl-go/xrpl/model/requests/common"

type LedgerRequestRequest struct {
	LedgerIndex common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash  `json:"ledger_hash,omitempty"`
}

func (*LedgerRequest) Method() string {
	return "ledger_request"
}
