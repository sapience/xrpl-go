package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type AccountNFTsResponse struct {
	Account            types.Address      `json:"account"`
	AccountNFTs        []NFT              `json:"account_nfts"`
	LedgerIndex        common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index,omitempty"`
	Validated          bool               `json:"validated"`
}
