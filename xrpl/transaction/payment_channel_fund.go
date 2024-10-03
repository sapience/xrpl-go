package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type PaymentChannelFund struct {
	BaseTx
	Channel    types.Hash256
	Amount     types.XRPCurrencyAmount
	Expiration uint `json:",omitempty"`
}

func (*PaymentChannelFund) TxType() TxType {
	return PaymentChannelFundTx
}

// TODO: Implement flatten
func (s *PaymentChannelFund) Flatten() FlatTransaction {
	return nil
}
