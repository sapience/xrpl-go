package ledger

import "github.com/Peersyst/xrpl-go/xrpl/queries/common"

// ############################################################################
// Request
// ############################################################################

type CurrentRequest struct {
}

func (*CurrentRequest) Method() string {
	return "ledger_current"
}

// ############################################################################
// Response
// ############################################################################

type CurrentResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index"`
}
