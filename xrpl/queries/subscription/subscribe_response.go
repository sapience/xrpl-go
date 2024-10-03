package subscribe

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
)

type SubscribeResponse struct {
	LoadBase         uint               `json:"load_base,omitempty"`
	LoadFactor       uint               `json:"load_factor,omitempty"`
	Random           string             `json:"random,omitempty"`
	ServerStatus     string             `json:"server_status,omitempty"`
	FeeBase          uint               `json:"fee_base,omitempty"`
	FeeRef           uint               `json:"fee_ref,omitempty"`
	LedgerHash       common.LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex      common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerTime       uint64             `json:"ledger_time,omitempty"`
	ReserveBase      uint               `json:"reserve_base,omitempty"`
	ReserveInc       uint               `json:"reserve_inc,omitempty"`
	ValidatedLedgers string             `json:"validated_ledgers,omitempty"`
	Offers           []ledger.Offer     `json:"offers,omitempty"`
}
