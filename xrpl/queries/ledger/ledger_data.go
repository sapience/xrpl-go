package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	ledgertypes "github.com/Peersyst/xrpl-go/xrpl/queries/ledger/types"
)

// ############################################################################
// Request
// ############################################################################

type DataRequest struct {
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Binary      bool                   `json:"binary,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
	Type        ledger.EntryType       `json:"type,omitempty"`
}

func (*DataRequest) Method() string {
	return "ledger_data"
}

// TODO: Implement V2
func (*DataRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type DataResponse struct {
	LedgerIndex string              `json:"ledger_index"`
	LedgerHash  common.LedgerHash   `json:"ledger_hash"`
	State       []ledgertypes.State `json:"state"`
	Marker      any                 `json:"marker"`
}
