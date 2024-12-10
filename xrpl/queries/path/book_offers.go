package path

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	pathtypes "github.com/Peersyst/xrpl-go/xrpl/queries/path/types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

type BookOffersRequest struct {
	TakerGets   pathtypes.BookOfferCurrency `json:"taker_gets"`
	TakerPays   pathtypes.BookOfferCurrency `json:"taker_pays"`
	Taker       types.Address               `json:"taker,omitempty"`
	LedgerHash  common.LedgerHash           `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerIndex          `json:"ledger_index,omitempty"`
	Limit       int                         `json:"limit,omitempty"`
}

func (*BookOffersRequest) Method() string {
	return "book_offers"
}

// TODO: Implement
func (*BookOffersRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type BookOffersResponse struct {
	LedgerCurrentIndex common.LedgerIndex    `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex    `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash     `json:"ledger_hash,omitempty"`
	Offers             []pathtypes.BookOffer `json:"offers"`
	Validated          bool                  `json:"validated,omitempty"`
}
