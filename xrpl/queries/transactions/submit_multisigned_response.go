package transaction

import "github.com/Peersyst/xrpl-go/xrpl/transaction"

type SubmitMultisignedResponse struct {
	EngineResult        string                       `json:"engine_result"`
	EngineResultCode    int                          `json:"engine_result_code"`
	EngineResultMessage string                       `json:"engine_result_message"`
	TxBlob              string                       `json:"tx_blob"`
	Tx                  transaction.FlatTransaction `json:"tx_json"`
}
