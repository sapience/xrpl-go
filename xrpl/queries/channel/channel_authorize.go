package channel

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type AuthorizeRequest struct {
	ChannelID  string                  `json:"channel_id"`
	Secret     string                  `json:"secret,omitempty"`
	Seed       string                  `json:"seed,omitempty"`
	SeedHex    string                  `json:"seed_hex,omitempty"`
	Passphrase string                  `json:"passphrase,omitempty"`
	KeyType    string                  `json:"key_type,omitempty"`
	Amount     types.XRPCurrencyAmount `json:"amount"`
}

func (*AuthorizeRequest) Method() string {
	return "channel_authorize"
}

// do not allow secrets to be printed
func (c *AuthorizeRequest) Format(s fmt.State, v rune) {
	type fHelper struct {
		ChannelID string                  `json:"channel_id"`
		KeyType   string                  `json:"key_type,omitempty"`
		Amount    types.XRPCurrencyAmount `json:"amount"`
	}
	h := fHelper{
		ChannelID: c.ChannelID,
		KeyType:   c.KeyType,
		Amount:    c.Amount,
	}
	fmt.Fprintf(s, "%"+string(v), h)
}

type AuthorizeResponse struct {
	Signature string `json:"signature"`
}
