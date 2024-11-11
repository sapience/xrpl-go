package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
)

type DataResponse struct {
	LedgerIndex string            `json:"ledger_index"`
	LedgerHash  common.LedgerHash `json:"ledger_hash"`
	State       []State           `json:"state"`
	Marker      any               `json:"marker"`
}

type State struct {
	Data            string                  `json:"data,omitempty"`
	LedgerEntryType ledger.EntryType        `json:",omitempty"`
	LedgerObject    ledger.FlatLedgerObject `json:"-"`
	Index           string                  `json:"index"`
}
