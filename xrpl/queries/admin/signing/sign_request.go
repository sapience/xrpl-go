package signing

import "github.com/Peersyst/xrpl-go/xrpl/transaction"

type SignRequest struct {
	TxJSON     transaction.FlatTransaction `json:"tx_json"`
	Secret     string                      `json:"secret,omitempty"`
	Seed       string                      `json:"seed,omitempty"`
	SeedHex    string                      `json:"seed_hex,omitempty"`
	Passphrase string                      `json:"passphrase,omitempty"`
	KeyType    string                      `json:"key_type,omitempty"`
	Offline    bool                        `json:"offline,omitempty"`
	BuildPath  bool                        `json:"build_path,omitempty"`
	FeeMultMax int                         `json:"fee_mult_max,omitempty"`
	FeeDivMax  int                         `json:"fee_div_max,omitempty"`
}

func (*SignRequest) Method() string {
	return "sign"
}
