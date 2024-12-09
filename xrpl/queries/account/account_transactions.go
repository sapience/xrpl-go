package account

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type TransactionsRequest struct {
	Account        types.Address          `json:"account"`
	LedgerIndexMin int                    `json:"ledger_index_min,omitempty"`
	LedgerIndexMax int                    `json:"ledger_index_max,omitempty"`
	LedgerHash     common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex    common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Binary         bool                   `json:"binary,omitempty"`
	Forward        bool                   `json:"forward,omitempty"`
	Limit          int                    `json:"limit,omitempty"`
	Marker         any                    `json:"marker,omitempty"`
}

func (*TransactionsRequest) Method() string {
	return "account_tx"
}

func (r *TransactionsRequest) UnmarshalJSON(data []byte) error {
	type atrHelper struct {
		Account        types.Address     `json:"account"`
		LedgerIndexMin int               `json:"ledger_index_min,omitempty"`
		LedgerIndexMax int               `json:"ledger_index_max,omitempty"`
		LedgerHash     common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex    json.RawMessage   `json:"ledger_index,omitempty"`
		Binary         bool              `json:"binary,omitempty"`
		Forward        bool              `json:"forward,omitempty"`
		Limit          int               `json:"limit,omitempty"`
		Marker         any               `json:"marker,omitempty"`
	}
	var h atrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = TransactionsRequest{
		Account:        h.Account,
		LedgerIndexMin: h.LedgerIndexMin,
		LedgerIndexMax: h.LedgerIndexMax,
		LedgerHash:     h.LedgerHash,
		Binary:         h.Binary,
		Forward:        h.Forward,
		Limit:          h.Limit,
		Marker:         h.Marker,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

type TransactionsResponse struct {
	Account        types.Address      `json:"account"`
	LedgerIndexMin common.LedgerIndex `json:"ledger_index_min"`
	LedgerIndexMax common.LedgerIndex `json:"ledger_index_max"`
	Limit          int                `json:"limit"`
	Marker         any                `json:"marker,omitempty"`
	Transactions   []Transaction      `json:"transactions"`
	Validated      bool               `json:"validated"`
}
