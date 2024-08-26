package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type PathStep struct {
	Account  types.Address `json:"account,omitempty"`
	Currency string        `json:"currency,omitempty"`
	Issuer   types.Address `json:"issuer,omitempty"`
}
