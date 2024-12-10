package channel

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

// ############################################################################
// Request
// ############################################################################

type VerifyRequest struct {
	Amount    types.XRPCurrencyAmount `json:"amount"`
	ChannelID string                  `json:"channel_id"`
	PublicKey string                  `json:"public_key"`
	Signature string                  `json:"signature"`
}

func (*VerifyRequest) Method() string {
	return "channel_verify"
}

// TODO: Implement V2
func (*VerifyRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type VerifyResponse struct {
	SignatureVerified bool `json:"signature_verified"`
}
