package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type Signer struct {
	SignerData SignerData `json:"Signer"`
}

type FlatSigner map[string]interface{}

func (s *Signer) Flatten() FlatSigner {
	flattened := make(FlatSigner)
	if s.SignerData != (SignerData{}) {
		flattened["SignerData"] = s.SignerData.Flatten()
	}
	return flattened
}

type SignerData struct {
	Account       types.Address
	TxnSignature  string
	SigningPubKey string
}

type FlatSignerData map[string]interface{}

func (sd *SignerData) Flatten() FlatSignerData {
	flattened := make(FlatSignerData)
	if sd.Account != "" {
		flattened["Account"] = sd.Account.String()
	}
	if sd.TxnSignature != "" {
		flattened["TxnSignature"] = sd.TxnSignature
	}
	if sd.SigningPubKey != "" {
		flattened["SigningPubKey"] = sd.SigningPubKey
	}
	return flattened
}
