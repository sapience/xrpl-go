package path

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type SubCommand string

const (
	CREATE SubCommand = "create"
	CLOSE  SubCommand = "close"
	STATUS SubCommand = "status"
)

type FindRequest struct {
	Subcommand         SubCommand             `json:"subcommand"`
	SourceAccount      types.Address          `json:"source_account,omitempty"`
	DestinationAccount types.Address          `json:"destination_account,omitempty"`
	DestinationAmount  types.CurrencyAmount   `json:"destination_amount,omitempty"`
	SendMax            types.CurrencyAmount   `json:"send_max,omitempty"`
	Paths              []transaction.PathStep `json:"paths,omitempty"`
}

func (*FindRequest) Method() string {
	return "path_find"
}

func (r *FindRequest) UnmarshalJSON(data []byte) error {
	type pfrHelper struct {
		Subcommand         SubCommand             `json:"subcommand"`
		SourceAccount      types.Address          `json:"source_account,omitempty"`
		DestinationAccount types.Address          `json:"destination_account,omitempty"`
		DestinationAmount  json.RawMessage        `json:"destination_amount,omitempty"`
		SendMax            json.RawMessage        `json:"send_max,omitempty"`
		Paths              []transaction.PathStep `json:"paths,omitempty"`
	}
	var h pfrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = FindRequest{
		Subcommand:         h.Subcommand,
		SourceAccount:      h.SourceAccount,
		DestinationAccount: h.DestinationAccount,
		Paths:              h.Paths,
	}

	var dest, max types.CurrencyAmount
	var err error
	dest, err = types.UnmarshalCurrencyAmount(h.DestinationAmount)
	if err != nil {
		return err
	}
	r.DestinationAmount = dest

	max, err = types.UnmarshalCurrencyAmount(h.SendMax)
	if err != nil {
		return err
	}
	r.SendMax = max

	return nil
}
