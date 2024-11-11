package clio

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type LedgerResponse struct {
	Ledger      Ledger             `json:"ledger"`
	LedgerHash  common.LedgerHash  `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
	Validated   bool               `json:"validated"`
}

type Ledger struct {
	AccountHash         string                         `json:"account_hash"`
	AccountState        []ledger.FlatLedgerObject      `json:"accountState,omitempty"`
	CloseFlags          int                            `json:"close_flags"`
	CloseTime           uint                           `json:"close_time"`
	CloseTimeHuman      string                         `json:"close_time_human"`
	CloseTimeResolution int                            `json:"close_time_resolution"`
	Closed              bool                           `json:"closed"`
	LedgerHash          common.LedgerHash              `json:"ledger_hash"`
	LedgerIndex         string                         `json:"ledger_index"`
	ParentCloseTime     uint                           `json:"parent_close_time"`
	ParentHash          string                         `json:"parent_hash"`
	TotalCoins          types.XRPCurrencyAmount        `json:"total_coins"`
	TransactionHash     string                         `json:"transaction_hash"`
	Transactions        []transactions.FlatTransaction `json:"transactions,omitempty"`
}
