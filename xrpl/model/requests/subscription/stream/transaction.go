package stream

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
)

type TransactionStream struct {
	Type                StreamType                   `json:"type"`
	EngineResult        string                       `json:"engine_result"`
	EngineResultCode    int                          `json:"engine_result_code"`
	EngineResultMessage string                       `json:"engine_result_message"`
	LedgerCurrentIndex  common.LedgerIndex           `json:"ledger_current_index,omitempty"`
	LedgerHash          common.LedgerHash            `json:"ledger_hash,omitempty"`
	LedgerIndex         common.LedgerIndex           `json:"ledger_index,omitempty"`
	Meta                transactions.TxObjMeta       `json:"meta,omitempty"`
	Transaction         transactions.FlatTransaction `json:"transaction"`
	Validated           bool                         `json:"validated"`
}
