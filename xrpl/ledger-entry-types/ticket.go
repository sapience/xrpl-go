package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type Ticket struct {
	Account           types.Address
	Flags             uint
	LedgerEntryType   EntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint
	TicketSequence    uint
}

func (*Ticket) EntryType() EntryType {
	return TicketEntry
}
