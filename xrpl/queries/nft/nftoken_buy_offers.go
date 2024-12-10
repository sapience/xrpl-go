package path

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"

	nfttypes "github.com/Peersyst/xrpl-go/xrpl/queries/nft/types"
)

// ############################################################################
// Request
// ############################################################################

type NFTokenBuyOffersRequest struct {
	NFTokenID   types.NFTokenID        `json:"nft_id"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
}

func (*NFTokenBuyOffersRequest) Method() string {
	return "nft_buy_offers"
}

// ############################################################################
// Response
// ############################################################################

type NFTokenBuyOffersResponse struct {
	NFTokenID types.NFTokenID         `json:"nft_id"`
	Offers    []nfttypes.NFTokenOffer `json:"offers"`
}
