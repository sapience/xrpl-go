package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type Escrow struct {
	Account           types.Address
	Amount            types.XRPCurrencyAmount
	CancelAfter       uint32 `json:",omitempty"`
	Condition         string `json:",omitempty"`
	Destination       types.Address
	DestinationNode   string `json:",omitempty"`
	DestinationTag    uint32 `json:",omitempty"`
	FinishAfter       uint32 `json:",omitempty"`
	Flags             uint32
	LedgerEntryType   EntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
	SourceTag         uint32 `json:",omitempty"`
}

func (*Escrow) EntryType() EntryType {
	return EscrowEntry
}
