package ledger

import "github.com/Peersyst/xrpl-go/xrpl/queries/common"

// ############################################################################
// Request
// ############################################################################

type ClosedRequest struct {
}

func (*ClosedRequest) Method() string {
	return "ledger_closed"
}

// TODO: Implement V2
func (*ClosedRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type ClosedResponse struct {
	LedgerHash  string             `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
}
