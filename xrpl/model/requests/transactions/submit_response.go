package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
)

type SubmitResponse struct {
	EngineResult             string                       `json:"engine_result"`
	EngineResultCode         int                          `json:"engine_result_code"`
	EngineResultMessage      string                       `json:"engine_result_message"`
	TxBlob                   string                       `json:"tx_blob"`
	Tx                       transactions.FlatTransaction `json:"tx_json"`
	Accepted                 bool                         `json:"accepted"`
	AccountSequenceAvailable uint                         `json:"account_sequence_available"`
	AccountSequenceNext      uint                         `json:"account_sequence_next"`
	Applied                  bool                         `json:"applied"`
	Broadcast                bool                         `json:"broadcast"`
	Kept                     bool                         `json:"kept"`
	Queued                   bool                         `json:"queued"`
	OpenLedgerCost           string                       `json:"open_ledger_cost"`
	ValidatedLedgerIndex     common.LedgerIndex           `json:"validated_ledger_index"`
}
