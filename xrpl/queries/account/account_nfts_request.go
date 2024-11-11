package account

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type NFTsRequest struct {
	Account     types.Address          `json:"account"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
}

func (*NFTsRequest) Method() string {
	return "account_nfts"
}

func (r *NFTsRequest) UnmarshalJSON(data []byte) error {
	type anrHelper struct {
		Account     types.Address     `json:"account"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		Limit       int               `json:"limit,omitempty"`
		Marker      any               `json:"marker,omitempty"`
	}
	var h anrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = NFTsRequest{
		Account:    h.Account,
		LedgerHash: h.LedgerHash,
		Limit:      h.Limit,
		Marker:     h.Marker,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
