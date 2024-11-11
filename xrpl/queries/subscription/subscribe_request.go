package subscribe

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type Request struct {
	Streams          []string        `json:"streams,omitempty"`
	Accounts         []types.Address `json:"accounts,omitempty"`
	AccountsProposed []types.Address `json:"accounts_proposed,omitempty"`
	Books            []OrderBook     `json:"books,omitempty"`
	URL              string          `json:"url,omitempty"`
	URLUsername      string          `json:"url_username,omitempty"`
	URLPassword      string          `json:"url_password,omitempty"`
}

func (*Request) Method() string {
	return "subscribe"
}

type OrderBook struct {
	TakerGets types.IssuedCurrencyAmount `json:"taker_gets"`
	TakerPays types.IssuedCurrencyAmount `json:"taker_pays"`
	Taker     types.Address              `json:"taker"`
	Snapshot  bool                       `json:"snapshot,omitempty"`
	Both      bool                       `json:"both,omitempty"`
}
