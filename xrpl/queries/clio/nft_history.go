package clio

import "github.com/Peersyst/xrpl-go/xrpl/transaction"

// ############################################################################
// Request
// ############################################################################

type NFTHistoryRequest struct {
	NFTokenID      string `json:"nft_id"`
	LedgerIndexMin uint   `json:"ledger_index_min,omitempty"`
	LedgerIndexMax uint   `json:"ledger_index_max,omitempty"`
	Binary         bool   `json:"binary,omitempty"`
	Forward        bool   `json:"forward,omitempty"`
	Limit          uint   `json:"limit,omitempty"`
	Marker         any    `json:"marker,omitempty"`
}

func (*NFTHistoryRequest) Method() string {
	return "nft_history"
}

// TODO: Implement V2
func (*NFTHistoryRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type NFTHistoryTransactions struct {
	LedgerIndex uint                        `json:"ledger_index"`
	Meta        transaction.TxObjMeta       `json:"meta"`
	Tx          transaction.FlatTransaction `json:"tx,omitempty"`
	TxBlob      string                      `json:"tx_blob,omitempty"`
	Validated   bool                        `json:"validated"`
}

type NFTHistoryResponse struct {
	NFTokenID      string                   `json:"nft_id"`
	LedgerIndexMin uint                     `json:"ledger_index_min,omitempty"`
	LedgerIndexMax uint                     `json:"ledger_index_max,omitempty"`
	Limit          uint                     `json:"limit,omitempty"`
	Marker         any                      `json:"marker,omitempty"`
	Transactions   []NFTHistoryTransactions `json:"transactions"`
	Validated      bool                     `json:"validated,omitempty"`
}
