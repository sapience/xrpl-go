package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type PaymentChannelClaim struct {
	BaseTx
	Channel   types.Hash256
	Balance   types.XRPCurrencyAmount `json:",omitempty"`
	Amount    types.XRPCurrencyAmount `json:",omitempty"`
	Signature string                  `json:",omitempty"`
	PublicKey string                  `json:",omitempty"`
}

func (*PaymentChannelClaim) TxType() TxType {
	return PaymentChannelClaimTx
}
