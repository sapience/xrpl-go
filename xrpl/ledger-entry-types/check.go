package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

// TODO verify format of SendMax
type Check struct {
	Account           types.Address
	Destination       types.Address
	DestinationNode   string `json:",omitempty"`
	DestinationTag    uint32 `json:",omitempty"`
	Expiration        uint32 `json:",omitempty"`
	Flags             uint32
	InvoiceID         types.Hash256 `json:",omitempty"`
	LedgerEntryType   EntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
	SendMax           string
	Sequence          uint32
	SourceTag         uint32 `json:",omitempty"`
}

func (*Check) EntryType() EntryType {
	return CheckEntry
}
