package account

import (
	"encoding/json"
	"errors"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type ChannelsRequest struct {
	Account            types.Address          `json:"account"`
	DestinationAccount types.Address          `json:"destination_account,omitempty"`
	LedgerIndex        common.LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash      `json:"ledger_hash,omitempty"`
	Limit              int                    `json:"limit,omitempty"`
	Marker             any                    `json:"marker,omitempty"`
}

func (*ChannelsRequest) Method() string {
	return "account_channels"
}

// Validate method to be added to each request struct
func (r *ChannelsRequest) Validate() error {
	if r.Account == "" {
		return errors.New("no account ID specified")
	}

	return nil
}

func (r *ChannelsRequest) UnmarshalJSON(data []byte) error {
	type acrHelper struct {
		Account            types.Address     `json:"account"`
		DestinationAccount types.Address     `json:"destination_account"`
		LedgerIndex        json.RawMessage   `json:"ledger_index,omitempty"`
		LedgerHash         common.LedgerHash `json:"ledger_hash,omitempty"`
		Limit              int               `json:"limit,omitempty"`
		Marker             any               `json:"marker,omitempty"`
	}
	var h acrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = ChannelsRequest{
		Account:            h.Account,
		DestinationAccount: h.DestinationAccount,
		LedgerHash:         h.LedgerHash,
		Limit:              h.Limit,
		Marker:             h.Marker,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
