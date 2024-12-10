package clio

import (
	cliotypes "github.com/Peersyst/xrpl-go/xrpl/queries/clio/types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

type NFTsByIssuerRequest struct {
	Issuer   types.Address `json:"issuer"`
	Marker   any           `json:"marker,omitempty"`
	Limit    int           `json:"limit,omitempty"`
	NftTaxon uint32        `json:"nft_taxon,omitempty"`
}

func (*NFTsByIssuerRequest) Method() string {
	return "nfts_by_issuer"
}

// TODO: Implement V2
func (*NFTsByIssuerRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type NFTsByIssuerResponse struct {
	Issuer       types.Address       `json:"issuer"`
	NFTs         []cliotypes.NFToken `json:"nfts"`
	Marker       any                 `json:"marker,omitempty"`
	Limit        int                 `json:"limit,omitempty"`
	NFTokenTaxon uint32              `json:"nft_taxon,omitempty"`
}
