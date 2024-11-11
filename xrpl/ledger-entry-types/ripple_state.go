package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

// TODO flags

type RippleState struct {
	Balance           types.IssuedCurrencyAmount
	Flags             uint
	HighLimit         types.IssuedCurrencyAmount
	HighNode          string
	HighQualityIn     uint `json:",omitempty"`
	HighQualityOut    uint `json:",omitempty"`
	LedgerEntryType   EntryType
	LowLimit          types.IssuedCurrencyAmount
	LowNode           string
	LowQualityIn      uint `json:",omitempty"`
	LowQualityOut     uint `json:",omitempty"`
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint
}

func (*RippleState) EntryType() EntryType {
	return RippleStateEntry
}
