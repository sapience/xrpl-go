package path

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type BookOffersRequest struct {
	TakerGets   types.IssuedCurrencyAmount `json:"taker_gets"`
	TakerPays   types.IssuedCurrencyAmount `json:"taker_pays"`
	LedgerHash  common.LedgerHash          `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier     `json:"ledger_index,omitempty"`
	Limit       int                        `json:"limit,omitempty"`
	Taker       types.Address              `json:"taker,omitempty"`
}

func (*BookOffersRequest) Method() string {
	return "book_offers"
}

func (r *BookOffersRequest) UnmarshalJSON(data []byte) error {
	type borHelper struct {
		TakerGets   types.IssuedCurrencyAmount `json:"taker_gets"`
		TakerPays   types.IssuedCurrencyAmount `json:"taker_pays"`
		LedgerHash  common.LedgerHash          `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage            `json:"ledger_index,omitempty"`
		Limit       int                        `json:"limit,omitempty"`
		Taker       types.Address              `json:"taker,omitempty"`
	}
	var h borHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*r = BookOffersRequest{
		TakerGets:  h.TakerGets,
		TakerPays:  h.TakerPays,
		LedgerHash: h.LedgerHash,
		Limit:      h.Limit,
		Taker:      h.Taker,
	}
	var i common.LedgerSpecifier
	i, err = common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

type BookOffersResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash  `json:"ledger_hash,omitempty"`
	Offers             []BookOffer        `json:"offers"`
}

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
