package path

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	nfttypes "github.com/Peersyst/xrpl-go/xrpl/queries/nft/types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

type NFTokenSellOffersRequest struct {
	NFTokenID   types.NFTokenID        `json:"nft_id"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
}

func (*NFTokenSellOffersRequest) Method() string {
	return "nft_sell_offers"
}

// ############################################################################
// Response
// ############################################################################

type NFTokenSellOffersResponse struct {
	NFTokenID types.NFTokenID         `json:"nft_id"`
	Offers    []nfttypes.NFTokenOffer `json:"offers"`
}
