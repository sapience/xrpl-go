package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type DepositPreauthObj struct {
	Account           types.Address
	Authorize         types.Address
	Flags             uint32
	LedgerEntryType   EntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
}

func (*DepositPreauthObj) EntryType() EntryType {
	return DepositPreauthObjEntry
}
