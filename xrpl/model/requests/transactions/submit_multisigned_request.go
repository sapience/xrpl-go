package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
)

type SubmitMultisignedRequest struct {
	Tx       transactions.FlatTransaction `json:"tx_json"`
	FailHard bool                         `json:"fail_hard"`
}

func (*SubmitMultisignedRequest) Method() string {
	return "submit_multisigned"
}
