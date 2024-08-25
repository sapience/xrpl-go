package ledger

import "github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"

type LedgerHashes struct {
	FirstLedgerSequence uint
	Flags               uint
	Hashes              []types.Hash256
	LastLedgerSequence  uint
	LedgerEntryType     LedgerEntryType
}

func (*LedgerHashes) EntryType() LedgerEntryType {
	return LedgerHashesEntry
}
