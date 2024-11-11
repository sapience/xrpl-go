package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type NFTokenPage struct {
	LedgerEntryType   EntryType
	NextPageMin       types.Hash256 `json:",omitempty"`
	PreviousPageMin   types.Hash256 `json:",omitempty"`
	PreviousTxnID     types.Hash256 `json:",omitempty"`
	PreviousTxnLgrSeq uint32        `json:",omitempty"`
	NFTokens          []types.NFToken
}

func (*NFTokenPage) EntryType() EntryType {
	return NFTokenPageEntry
}
