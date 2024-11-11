package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type Hashes struct {
	FirstLedgerSequence uint
	Flags               uint
	Hashes              []types.Hash256
	LastLedgerSequence  uint
	LedgerEntryType     EntryType
}

func (*Hashes) EntryType() EntryType {
	return LedgerHashesEntry
}
