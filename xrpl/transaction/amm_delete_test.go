package transaction

import (
	"testing"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestAMMDelete_TxType(t *testing.T) {
	tx := &AMMDelete{}
	assert.Equal(t, AMMDeleteTx, tx.TxType())
}
func TestAMMDelete_Flatten(t *testing.T) {
	tx := &AMMDelete{
		BaseTx: BaseTx{
			Account:  "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
			Fee:      types.XRPCurrencyAmount(10),
			Sequence: 9,
		},
		Asset: ledger.Asset{
			Currency: "XRP",
		},
		Asset2: ledger.Asset{
			Currency: "TST",
			Issuer:   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
		},
	}

	flattened := tx.Flatten()

	expected := `{
	"Account": "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
	"Fee": "10",
	"Sequence": 9,
	"TransactionType": "AMMDelete",
	"Asset": {
		"currency": "XRP"
	},
	"Asset2": {
		"currency": "TST",
		"issuer": "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
	}
}`

	err := testutil.CompareFlattenAndExpected(flattened, []byte(expected))
	if err != nil {
		t.Error(err)
	}
}
