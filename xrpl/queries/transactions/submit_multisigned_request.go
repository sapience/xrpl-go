package transaction

import "github.com/Peersyst/xrpl-go/xrpl/transaction"

type SubmitMultisignedRequest struct {
	Tx       transaction.FlatTransaction `json:"tx_json"`
	FailHard bool                         `json:"fail_hard"`
}

func (*SubmitMultisignedRequest) Method() string {
	return "submit_multisigned"
}
