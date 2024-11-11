package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

const (
	ErrAccountObjectUnmarshal string = "Unmarshal JSON AccountObjects"
)

type ObjectsResponse struct {
	Account            types.Address             `json:"account"`
	AccountObjects     []ledger.FlatLedgerObject `json:"account_objects"`
	LedgerHash         common.LedgerHash         `json:"ledger_hash,omitempty"`
	LedgerIndex        common.LedgerIndex        `json:"ledger_index,omitempty"`
	LedgerCurrentIndex common.LedgerIndex        `json:"ledger_current_index,omitempty"`
	Limit              int                       `json:"limit,omitempty"`
	Marker             any                       `json:"marker,omitempty"`
	Validated          bool                      `json:"validated,omitempty"`
}
