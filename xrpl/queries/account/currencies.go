package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

type CurrenciesRequest struct {
	Account     types.Address          `json:"account"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Strict      bool                   `json:"strict,omitempty"`
}

func (*CurrenciesRequest) Method() string {
	return "account_currencies"
}

// ############################################################################
// Response
// ############################################################################

type CurrenciesResponse struct {
	LedgerHash        common.LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex       common.LedgerIndex `json:"ledger_index"`
	ReceiveCurrencies []string           `json:"receive_currencies"`
	SendCurrencies    []string           `json:"send_currencies"`
	Validated         bool               `json:"validated"`
}
