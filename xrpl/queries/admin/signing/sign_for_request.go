package signing

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type SignForRequest struct {
	Account    types.Address               `json:"account"`
	TxJson     transaction.FlatTransaction `json:"tx_json"`
	Secret     string                      `json:"secret,omitempty"`
	Seed       string                      `json:"seed,omitempty"`
	SeedHex    string                      `json:"seed_hex,omitempty"`
	Passphrase string                      `json:"passphrase,omitempty"`
	KeyType    string                      `json:"key_type,omitempty"`
}

func (*SignForRequest) Method() string {
	return "sign_for"
}
