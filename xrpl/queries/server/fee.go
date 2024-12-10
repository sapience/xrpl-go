package server

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	servertypes "github.com/Peersyst/xrpl-go/xrpl/queries/server/types"
)

// ############################################################################
// Request
// ############################################################################

type FeeRequest struct {
}

func (*FeeRequest) Method() string {
	return "fee"
}

// TODO: Implement V2
func (*FeeRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################
type FeeResponse struct {
	CurrentLedgerSize  string                `json:"current_ledger_size"`
	CurrentQueueSize   string                `json:"current_queue_size"`
	Drops              servertypes.FeeDrops  `json:"drops"`
	ExpectedLedgerSize string                `json:"expected_ledger_size"`
	LedgerCurrentIndex common.LedgerIndex    `json:"ledger_current_index"`
	Levels             servertypes.FeeLevels `json:"levels"`
	MaxQueueSize       string                `json:"max_queue_size"`
}
