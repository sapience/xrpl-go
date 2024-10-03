package channel

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type ChannelVerifyRequest struct {
	Amount    types.XRPCurrencyAmount `json:"amount"`
	ChannelID string                  `json:"channel_id"`
	PublicKey string                  `json:"public_key"`
	Signature string                  `json:"signature"`
}

func (*ChannelVerifyRequest) Method() string {
	return "channel_verify"
}
