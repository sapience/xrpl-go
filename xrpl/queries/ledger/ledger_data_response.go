package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
)

type LedgerDataResponse struct {
	LedgerIndex string            `json:"ledger_index"`
	LedgerHash  common.LedgerHash `json:"ledger_hash"`
	State       []LedgerState     `json:"state"`
	Marker      any               `json:"marker"`
}

type LedgerState struct {
	Data            string                  `json:"data,omitempty"`
	LedgerEntryType ledger.LedgerEntryType  `json:",omitempty"`
	LedgerObject    ledger.FlatLedgerObject `json:"-"`
	Index           string                  `json:"index"`
}
