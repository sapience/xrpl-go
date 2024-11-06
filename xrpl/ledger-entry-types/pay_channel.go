package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type PayChannel struct {
	Account           types.Address
	Amount            types.XRPCurrencyAmount
	Balance           types.XRPCurrencyAmount
	CancelAfter       uint32 `json:",omitempty"`
	Destination       types.Address
	DestinationTag    uint32 `json:",omitempty"`
	DestinationNode   string `json:",omitempty"`
	Expiration        uint32 `json:",omitempty"`
	Flags             uint32
	LedgerEntryType   EntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
	PublicKey         string
	SettleDelay       uint32
	SourceTag         uint32 `json:",omitempty"`
}

func (*PayChannel) EntryType() EntryType {
	return PayChannelEntry
}
