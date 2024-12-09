package channel

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type VerifyRequest struct {
	Amount    types.XRPCurrencyAmount `json:"amount"`
	ChannelID string                  `json:"channel_id"`
	PublicKey string                  `json:"public_key"`
	Signature string                  `json:"signature"`
}

func (*VerifyRequest) Method() string {
	return "channel_verify"
}

type VerifyResponse struct {
	SignatureVerified bool `json:"signature_verified"`
}
