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

	var dest, sendMax types.CurrencyAmount
	var err error
	dest, err = types.UnmarshalCurrencyAmount(h.DestinationAmount)
	if err != nil {
		return err
	}
	r.DestinationAmount = dest

	sendMax, err = types.UnmarshalCurrencyAmount(h.SendMax)
	if err != nil {
		return err
	}
	r.SendMax = sendMax

	return nil
}

type FindResponse struct {
	Alternatives       []Alternative        `json:"alternatives"`
	DestinationAccount types.Address        `json:"destination_account"`
	DestinationAmount  types.CurrencyAmount `json:"destination_amount"`
	SourceAccount      types.Address        `json:"source_account"`
	FullReply          bool                 `json:"full_reply"`
	Closed             bool                 `json:"closed,omitempty"`
	Status             bool                 `json:"status,omitempty"`
}

func (r *FindResponse) UnmarshalJSON(data []byte) error {
	type pfrHelper struct {
		Alternatives       []Alternative   `json:"alternatives"`
		DestinationAccount types.Address   `json:"destination_account"`
		DestinationAmount  json.RawMessage `json:"destination_amount"`
		SourceAccount      types.Address   `json:"source_account"`
		FullReply          bool            `json:"full_reply"`
		Closed             bool            `json:"closed,omitempty"`
		Status             bool            `json:"status,omitempty"`
	}
	var h pfrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = FindResponse{
		Alternatives:       h.Alternatives,
		DestinationAccount: h.DestinationAccount,
		SourceAccount:      h.SourceAccount,
		FullReply:          h.FullReply,
		Closed:             h.Closed,
		Status:             h.Status,
	}

	dst, err := types.UnmarshalCurrencyAmount(h.DestinationAmount)
	if err != nil {
		return err
	}
	r.DestinationAmount = dst

	return nil
}

type Alternative struct {
	PathsComputed     [][]transaction.PathStep `json:"paths_computed"`
	SourceAmount      types.CurrencyAmount     `json:"source_amount"`
	DestinationAmount types.CurrencyAmount     `json:"destination_amount,omitempty"`
}

func (p *Alternative) UnmarshalJSON(data []byte) error {
	type paHelper struct {
		PathsComputed     [][]transaction.PathStep `json:"paths_computed"`
		SourceAmount      json.RawMessage          `json:"source_amount"`
		DestinationAmount json.RawMessage          `json:"destination_amount,omitempty"`
	}
	var h paHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	p.PathsComputed = h.PathsComputed

	var src, dst types.CurrencyAmount
	var err error

	src, err = types.UnmarshalCurrencyAmount(h.SourceAmount)
	if err != nil {
		return err
	}
	p.SourceAmount = src

	dst, err = types.UnmarshalCurrencyAmount(h.DestinationAmount)
	if err != nil {
		return err
	}
	p.DestinationAmount = dst

	return nil
}
