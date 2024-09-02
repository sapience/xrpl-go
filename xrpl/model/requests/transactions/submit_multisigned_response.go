package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
)

type SubmitMultisignedResponse struct {
	EngineResult        string                       `json:"engine_result"`
	EngineResultCode    int                          `json:"engine_result_code"`
	EngineResultMessage string                       `json:"engine_result_message"`
	TxBlob              string                       `json:"tx_blob"`
	Tx                  transactions.FlatTransaction `json:"tx_json"`
}
