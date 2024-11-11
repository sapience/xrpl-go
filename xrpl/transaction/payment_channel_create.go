package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type PaymentChannelCreate struct {
	BaseTx
	Amount         types.XRPCurrencyAmount
	Destination    types.Address
	SettleDelay    uint32
	PublicKey      string
	CancelAfter    uint32 `json:",omitempty"`
	DestinationTag uint32 `json:",omitempty"`
}

func (*PaymentChannelCreate) TxType() TxType {
	return PaymentChannelCreateTx
}

// TODO: Implement flatten
func (s *PaymentChannelCreate) Flatten() FlatTransaction {
	return nil
}
