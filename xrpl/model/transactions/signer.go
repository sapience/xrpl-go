package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type Signer struct {
	SignerData SignerData `json:"Signer"`
}

type SignerData struct {
	Account       types.Address
	TxnSignature  string
	SigningPubKey string
}
