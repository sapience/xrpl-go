package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type PaymentChannelCreate struct {
	BaseTx
	Amount         types.XRPCurrencyAmount
	Destination    types.Address
	SettleDelay    uint
	PublicKey      string
	CancelAfter    uint `json:",omitempty"`
	DestinationTag uint `json:",omitempty"`
}

func (*PaymentChannelCreate) TxType() TxType {
	return PaymentChannelCreateTx
}

// TODO: Implement flatten
func (s *PaymentChannelCreate) Flatten() map[string]interface{} {
	return nil
}
