package transaction

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
)

type EntryRequest struct {
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	TxHash      string                 `json:"tx_hash"`
}

func (*EntryRequest) Method() string {
	return "transaction_entry"
}

func (t *EntryRequest) UnmarshalJSON(data []byte) error {
	type terHelper struct {
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		TxHash      string            `json:"tx_hash"`
	}
	var h terHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*t = EntryRequest{
		LedgerHash: h.LedgerHash,
		TxHash:     h.TxHash,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	t.LedgerIndex = i
	return nil
}

type EntryResponse struct {
	LedgerIndex common.LedgerIndex           `json:"ledger_index"`
	LedgerHash  common.LedgerHash            `json:"ledger_hash,omitempty"`
	Metadata    transaction.TxObjMeta       `json:"metadata"`
	Tx          transaction.FlatTransaction `json:"tx_json"`
}
