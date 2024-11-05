package ledger

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
)

type Request struct {
	LedgerHash   common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex  common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Full         bool                   `json:"full,omitempty"`
	Accounts     bool                   `json:"accounts,omitempty"`
	Transactions bool                   `json:"transactions,omitempty"`
	OwnerFunds   bool                   `json:"owner_funds,omitempty"`
	Binary       bool                   `json:"binary,omitempty"`
	Queue        bool                   `json:"queue,omitempty"`
	Type         ledger.EntryType       `json:"type,omitempty"`
}

func (*Request) Method() string {
	return "ledger"
}

// TODO: Implement
func (*Request) Validate() error {
	return nil
}

func (r *Request) UnmarshalJSON(data []byte) error {
	type lrHelper struct {
		LedgerHash   common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex  json.RawMessage   `json:"ledger_index,omitempty"`
		Full         bool              `json:"full,omitempty"`
		Accounts     bool              `json:"accounts,omitempty"`
		Transactions bool              `json:"transactions,omitempty"`
		OwnerFunds   bool              `json:"owner_funds,omitempty"`
		Binary       bool              `json:"binary,omitempty"`
		Queue        bool              `json:"queue,omitempty"`
		Type         ledger.EntryType  `json:"type,omitempty"`
	}
	var h lrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = Request{
		LedgerHash:   h.LedgerHash,
		Full:         h.Full,
		Accounts:     h.Accounts,
		Transactions: h.Transactions,
		OwnerFunds:   h.OwnerFunds,
		Binary:       h.Binary,
		Queue:        h.Queue,
		Type:         h.Type,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
