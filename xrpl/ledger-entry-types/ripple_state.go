package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

// TODO flags

type RippleState struct {
	Balance           types.IssuedCurrencyAmount
	Flags             uint32
	HighLimit         types.IssuedCurrencyAmount
	HighNode          string
	HighQualityIn     uint32 `json:",omitempty"`
	HighQualityOut    uint32 `json:",omitempty"`
	LedgerEntryType   EntryType
	LowLimit          types.IssuedCurrencyAmount
	LowNode           string
	LowQualityIn      uint32 `json:",omitempty"`
	LowQualityOut     uint32 `json:",omitempty"`
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
}

func (*RippleState) EntryType() EntryType {
	return RippleStateEntry
}
