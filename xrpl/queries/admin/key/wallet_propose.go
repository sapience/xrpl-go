package key

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type WalletProposeRequest struct {
	KeyType    string `json:"key_type,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
	Seed       string `json:"seed,omitempty"`
	SeedHex    string `json:"seed_hex,omitempty"`
}

func (*WalletProposeRequest) Method() string {
	return "wallet_propose"
}

type WalletProposeResponse struct {
	KeyType       string        `json:"key_type"`
	MasterSeed    string        `json:"master_seed"`
	MasterSeedHex string        `json:"master_seed_hex"`
	AccountID     types.Address `json:"account_id"`
	PublicKey     string        `json:"public_key"`
	PublicKeyHex  string        `json:"public_key_hex"`
	Warning       string        `json:"warning,omitempty"`
}
