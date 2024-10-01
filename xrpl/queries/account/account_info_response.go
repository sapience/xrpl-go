package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
)

type AccountInfoResponse struct {
	AccountData        ledger.AccountRoot  `json:"account_data"`
	SignerLists        []ledger.SignerList `json:"signer_lists,omitempty"`
	LedgerCurrentIndex common.LedgerIndex  `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex  `json:"ledger_index,omitempty"`
	QueueData          QueueData           `json:"queue_data,omitempty"`
	Validated          bool                `json:"validated"`
}
