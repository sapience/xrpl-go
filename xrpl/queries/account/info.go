package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	accounttypes "github.com/Peersyst/xrpl-go/xrpl/queries/account/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

type InfoRequest struct {
	Account     types.Address          `json:"account"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	Queue       bool                   `json:"queue,omitempty"`
	SignerList  bool                   `json:"signer_list,omitempty"`
	Strict      bool                   `json:"strict,omitempty"`
}

func (*InfoRequest) Method() string {
	return "account_info"
}

func (*InfoRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type InfoResponse struct {
	AccountData        ledger.AccountRoot  `json:"account_data"`
	SignerLists        []ledger.SignerList `json:"signer_lists,omitempty"`
	LedgerCurrentIndex common.LedgerIndex  `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex  `json:"ledger_index,omitempty"`
	QueueData          accounttypes.QueueData           `json:"queue_data,omitempty"`
	Validated          bool                `json:"validated"`
}
