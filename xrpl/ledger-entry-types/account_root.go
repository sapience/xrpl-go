package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type AccountRoot struct {
	Account           types.Address
	AccountTxnID      types.Hash256           `json:",omitempty"`
	Balance           types.XRPCurrencyAmount `json:",omitempty"`
	BurnedNFTokens    uint32                  `json:",omitempty"`
	Domain            string                  `json:",omitempty"`
	EmailHash         types.Hash128           `json:",omitempty"`
	Flags             uint64
	LedgerEntryType   EntryType
	MessageKey        string        `json:",omitempty"`
	MintedNFTokens    uint32        `json:",omitempty"`
	NFTokenMinter     types.Address `json:",omitempty"`
	OwnerCount        uint64
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint64
	RegularKey        types.Address `json:",omitempty"`
	Sequence          uint64
	TicketCount       uint32 `json:",omitempty"`
	TickSize          uint8  `json:",omitempty"`
	TransferRate      uint32 `json:",omitempty"`
	// TODO: determine if this is a required field
	// Index             types.Hash256 `json:"index,omitempty"`
}

func (*AccountRoot) EntryType() EntryType {
	return AccountRootEntry
}
