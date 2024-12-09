package types

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

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
