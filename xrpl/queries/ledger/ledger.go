package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	ledgertypes "github.com/Peersyst/xrpl-go/xrpl/queries/ledger/types"
)

// ############################################################################
// Request
// ############################################################################

type Request struct {
	LedgerHash   common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex  common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Full         bool                   `json:"full,omitempty"`
	Accounts     bool                   `json:"accounts,omitempty"`
	Expand       bool                   `json:"expand,omitempty"`
	Transactions bool                   `json:"transactions,omitempty"`
	OwnerFunds   bool                   `json:"owner_funds,omitempty"`
	Binary       bool                   `json:"binary,omitempty"`
	Queue        bool                   `json:"queue,omitempty"`
	Type         ledger.EntryType       `json:"type,omitempty"`
}

func (*Request) Method() string {
	return "ledger"
}

// TODO: Implement V2
func (*Request) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type Response struct {
	Ledger      ledgertypes.BaseLedger  `json:"ledger"`
	LedgerHash  string                  `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex      `json:"ledger_index"`
	Validated   bool                    `json:"validated,omitempty"`
	QueueData   []ledgertypes.QueueData `json:"queue_data,omitempty"`
}
