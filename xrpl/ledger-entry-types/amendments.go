package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type Amendments struct {
	Amendments      []types.Hash256 `json:",omitempty"`
	Flags           uint32
	LedgerEntryType EntryType
	Majorities      []MajorityEntry `json:",omitempty"`
}

func (*Amendments) EntryType() EntryType {
	return AmendmentsEntry
}

type MajorityEntry struct {
	Majority Majority
}

type Majority struct {
	Amendment types.Hash256
	CloseTime uint32
}
