package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type SignerListFlags uint32

const (
	LsfOneOwnerCount SignerListFlags = 0x00010000
)

type SignerList struct {
	LedgerEntryType   EntryType
	Flags             SignerListFlags
	PreviousTxnID     string
	PreviousTxnLgrSeq uint32
	OwnerNode         string
	SignerEntries     []SignerEntryWrapper
	SignerListID      uint32
	SignerQuorum      uint32
}

type SignerEntryWrapper struct {
	SignerEntry SignerEntry
}

type SignerEntry struct {
	Account       types.Address
	SignerWeight  uint16
	WalletLocator types.Hash256 `json:",omitempty"`
}

func (*SignerList) EntryType() EntryType {
	return SignerListEntry
}
