package transaction

import (
	"testing"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestAMMVote_TxType(t *testing.T) {
	tx := &AMMVote{}
	assert.Equal(t, AMMVoteTx, tx.TxType())
}
func TestAMMVote_Flatten(t *testing.T) {
	tx := &AMMVote{
		BaseTx: BaseTx{
			Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
			TransactionType: "AMMVote",
			Fee:             types.XRPCurrencyAmount(10),
			Flags:           2147483648,
			Sequence:        8,
		},
		Asset: ledger.Asset{
			Currency: "XRP",
		},
		Asset2: ledger.Asset{
			Currency: "TST",
			Issuer:   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
		},
		TradingFee: 600,
	}

	flattened := tx.Flatten()

	expected := `{
		"Account":         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
		"Fee":             "10",
		"Flags":           2147483648,
		"Sequence":        8,
		"TransactionType": "AMMVote",
		"Asset": {
			"currency": "XRP"
		},
		"Asset2": {
			"currency": "TST",
			"issuer":   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
		},
		"TradingFee": 600
	}`

	err := testutil.CompareFlattenAndExpected(flattened, []byte(expected))
	if err != nil {
		t.Error(err)
	}
}
