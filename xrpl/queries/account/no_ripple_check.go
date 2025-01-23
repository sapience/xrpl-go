package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type NoRippleCheckRequest struct {
	common.BaseRequest
	Account      types.Address          `json:"account"`
	Role         string                 `json:"role"`
	Transactions bool                   `json:"transactions,omitempty"`
	Limit        int                    `json:"limit,omitempty"`
	LedgerHash   common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex  common.LedgerSpecifier `json:"ledger_index,omitempty"`
}

func (*NoRippleCheckRequest) Method() string {
	return "noripple_check"
}

func (*NoRippleCheckRequest) APIVersion() int {
	return version.RippledAPIV2
}

func (*NoRippleCheckRequest) Validate() error {
	return nil
}

type NoRippleCheckResponse struct {
	LedgerCurrentIndex common.LedgerIndex            `json:"ledger_current_index"`
	Problems           []string                      `json:"problems"`
	Transactions       []transaction.FlatTransaction `json:"transactions"`
}
