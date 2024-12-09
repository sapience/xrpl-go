package types

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type BookOffer struct {
	ledger.Offer
	OwnerFunds      string               `json:"owner_funds,omitempty"`
	TakerGetsFunded types.CurrencyAmount `json:"taker_gets_funded,omitempty"`
	TakerPaysFunded types.CurrencyAmount `json:"taker_pays_funded,omitempty"`
	Quality         string               `json:"quality,omitempty"`
}

func (o *BookOffer) UnmarshalJSON(data []byte) error {
	type boHelper struct {
		OwnerFunds      string          `json:"offer_funds,omitempty"`
		TakerGetsFunded json.RawMessage `json:"taker_gets_funded,omitempty"`
		TakerPaysFunded json.RawMessage `json:"taker_pays_funded,omitempty"`
		Quality         string          `json:"quality,omitempty"`
	}
	var h boHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	var offer ledger.Offer
	err = json.Unmarshal(data, &offer)
	if err != nil {
		return err
	}
	*o = BookOffer{
		Offer:      offer,
		OwnerFunds: h.OwnerFunds,
		Quality:    h.Quality,
	}
	var g, p types.CurrencyAmount
	g, err = types.UnmarshalCurrencyAmount(h.TakerGetsFunded)
	if err != nil {
		return err
	}
	o.TakerGetsFunded = g
	p, err = types.UnmarshalCurrencyAmount(h.TakerPaysFunded)
	if err != nil {
		return err
	}
	o.TakerPaysFunded = p

	return nil
}
