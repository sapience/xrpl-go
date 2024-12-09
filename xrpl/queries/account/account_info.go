package account

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

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

func (r *InfoRequest) UnmarshalJSON(data []byte) error {
	type airHelper struct {
		Account     types.Address     `json:"account"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		Queue       bool              `json:"queue,omitempty"`
		SignerList  bool              `json:"signer_list,omitempty"`
		Strict      bool              `json:"strict,omitempty"`
	}
	var h airHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = InfoRequest{
		Account:    h.Account,
		LedgerHash: h.LedgerHash,
		Queue:      h.Queue,
		SignerList: h.SignerList,
		Strict:     h.Strict,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

type InfoResponse struct {
	AccountData        ledger.AccountRoot  `json:"account_data"`
	SignerLists        []ledger.SignerList `json:"signer_lists,omitempty"`
	LedgerCurrentIndex common.LedgerIndex  `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex  `json:"ledger_index,omitempty"`
	QueueData          QueueData           `json:"queue_data,omitempty"`
	Validated          bool                `json:"validated"`
}
