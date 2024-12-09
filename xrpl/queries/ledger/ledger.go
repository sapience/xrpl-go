package ledger

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
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

type Response struct {
	Ledger      Header             `json:"ledger"`
	LedgerHash  string             `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
	Validated   bool               `json:"validated,omitempty"`
	QueueData   []QueueData        `json:"queue_data,omitempty"`
}

type Header struct {
	AccountHash         string                         `json:"account_hash"`
	AccountState        []ledger.FlatLedgerObject      `json:"accountState,omitempty"`
	CloseFlags          int                            `json:"close_flags"`
	CloseTime           int                            `json:"close_time"`
	CloseTimeHuman      string                         `json:"close_time_human"`
	CloseTimeResolution int                            `json:"close_time_resolution"`
	Closed              bool                           `json:"closed"`
	LedgerHash          string                         `json:"ledger_hash"`
	LedgerIndex         string                         `json:"ledger_index"`
	ParentCloseTime     int                            `json:"parent_close_time"`
	ParentHash          string                         `json:"parent_hash"`
	TotalCoins          types.XRPCurrencyAmount        `json:"total_coins"`
	TransactionHash     string                         `json:"transaction_hash"`
	Transactions        []transaction.FlatTransaction `json:"transactions,omitempty"`
}

type QueueData struct {
	Account          types.Address                `json:"account"`
	Tx               transaction.FlatTransaction `json:"tx"`
	RetriesRemaining int                          `json:"retries_remaining"`
	PreflightResult  string                       `json:"preflight_result"`
	LastResult       string                       `json:"last_result,omitempty"`
	AuthChange       bool                         `json:"auth_change,omitempty"`
	Fee              types.XRPCurrencyAmount      `json:"fee,omitempty"`
	FeeLevel         types.XRPCurrencyAmount      `json:"fee_level,omitempty"`
	MaxSpendDrops    types.XRPCurrencyAmount      `json:"max_spend_drops,omitempty"`
}
