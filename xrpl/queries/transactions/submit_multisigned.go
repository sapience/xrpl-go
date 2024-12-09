package transaction

import "github.com/Peersyst/xrpl-go/xrpl/transaction"

type SubmitMultisignedRequest struct {
	Tx       transaction.FlatTransaction `json:"tx_json"`
	FailHard bool                        `json:"fail_hard"`
}

func (*SubmitMultisignedRequest) Method() string {
	return "submit_multisigned"
}

type SubmitMultisignedResponse struct {
	EngineResult        string                      `json:"engine_result"`
	EngineResultCode    int                         `json:"engine_result_code"`
	EngineResultMessage string                      `json:"engine_result_message"`
	TxBlob              string                      `json:"tx_blob"`
	Tx                  transaction.FlatTransaction `json:"tx_json"`
}
