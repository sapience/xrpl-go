package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type Ticket struct {
	Account           types.Address
	Flags             uint32
	LedgerEntryType   EntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
	TicketSequence    uint32
}

func (*Ticket) EntryType() EntryType {
	return TicketEntry
}
