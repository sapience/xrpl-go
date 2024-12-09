package path

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	pathtypes "github.com/Peersyst/xrpl-go/xrpl/queries/path/types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type BookOffersRequest struct {
	TakerGets   types.IssuedCurrencyAmount `json:"taker_gets"`
	TakerPays   types.IssuedCurrencyAmount `json:"taker_pays"`
	LedgerHash  common.LedgerHash          `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerIndex`json:"ledger_index,omitempty"`
	Limit       int                        `json:"limit,omitempty"`
	Taker       types.Address              `json:"taker,omitempty"`
}

func (*BookOffersRequest) Method() string {
	return "book_offers"
}

// TODO: Implement
func (*BookOffersRequest) Validate() error {
	return nil
}

type BookOffersResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash  `json:"ledger_hash,omitempty"`
	Offers             []pathtypes.BookOffer        `json:"offers"`
	Validated          bool               `json:"validated,omitempty"`
}
