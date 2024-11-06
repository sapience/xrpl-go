package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type Hashes struct {
	FirstLedgerSequence uint32
	Flags               uint32
	Hashes              []types.Hash256
	LastLedgerSequence  uint32
	LedgerEntryType     EntryType
}

func (*Hashes) EntryType() EntryType {
	return LedgerHashesEntry
}
