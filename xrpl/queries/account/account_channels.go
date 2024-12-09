package account

import (
	"encoding/json"
	"errors"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// The account_channels method returns information about an account's Payment
// Channels. This includes only channels where the specified account is the
// channel's source, not the destination. (A channel's "source" and "owner" are
// the same.) All information retrieved is relative to a particular version of
// the ledger. Returns an {@link AccountChannelsResponse}.
type ChannelsRequest struct {
	Account            types.Address          `json:"account"`
	DestinationAccount types.Address          `json:"destination_account,omitempty"`
	LedgerIndex        common.LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash      `json:"ledger_hash,omitempty"`
	Limit              int                    `json:"limit,omitempty"`
	Marker             any                    `json:"marker,omitempty"`
}

// Method returns the method name for the ChannelsRequest.
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

// The expected response from the account_channels method.
type ChannelsResponse struct {
	Account     types.Address      `json:"account"`
	Channels    []ChannelResult    `json:"channels"`
	LedgerIndex common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash  `json:"ledger_hash,omitempty"`
	Validated   bool               `json:"validated,omitempty"`
	Limit       int                `json:"limit,omitempty"`
	Marker      any                `json:"marker,omitempty"`
}
