package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

const (
	TfRenew uint = 65536
	TfClose uint = 131072
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

// TODO: Implement flatten
func (s *PaymentChannelClaim) Flatten() map[string]interface{} {
	return nil
}

func (s *PaymentChannelClaim) SetRenewFlag(enabled bool) {
	if enabled {
		s.Flags |= TfRenew
	} else {
		s.Flags &= ^TfRenew
	}
}

func (s *PaymentChannelClaim) SetCloseFlag(enabled bool) {
	if enabled {
		s.Flags |= TfClose
	} else {
		s.Flags &= ^TfClose
	}
}
