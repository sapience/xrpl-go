package ledger

import "github.com/Peersyst/xrpl-go/xrpl/queries/common"

type ClosedRequest struct {
}

func (*ClosedRequest) Method() string {
	return "ledger_closed"
}

type ClosedResponse struct {
	LedgerHash  string             `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
}
