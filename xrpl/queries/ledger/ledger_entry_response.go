package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
)

type LedgerEntryResponse struct {
	Index       string                  `json:"index"`
	LedgerIndex common.LedgerIndex      `json:"ledger_index"`
	Node        ledger.FlatLedgerObject `json:"node,omitempty"`
	NodeBinary  string                  `json:"node_binary,omitempty"`
}
