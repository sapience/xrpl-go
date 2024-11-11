package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type DirectoryNode struct {
	Flags             uint
	Indexes           []types.Hash256
	IndexNext         string `json:",omitempty"`
	IndexPrevious     string `json:",omitempty"`
	LedgerEntryType   EntryType
	Owner             types.Address `json:",omitempty"`
	RootIndex         types.Hash256
	TakerGetsCurrency string `json:",omitempty"`
	TakerGetsIssuer   string `json:",omitempty"`
	TakerPaysCurrency string `json:",omitempty"`
	TakerPaysIssuer   string `json:",omitempty"`
}

func (*DirectoryNode) EntryType() EntryType {
	return DirectoryNodeEntry
}
