package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestOfferCancel_TxType(t *testing.T) {
	tx := &OfferCancel{}
	assert.Equal(t, OfferCancelTx, tx.TxType())
}

func TestOfferCancel_Flatten(t *testing.T) {
	tx := &OfferCancel{
		BaseTx: BaseTx{
			Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType:    OfferCancelTx,
			Fee:                types.XRPCurrencyAmount(10),
			Flags:              123,
			LastLedgerSequence: 7108629,
			Sequence:           7,
		},
		OfferSequence: 6,
	}

	expected := `{
		"Account":            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType":    "OfferCancel",
		"Fee":                "10",
		"Flags":              123,
		"LastLedgerSequence": 7108629,
		"Sequence":           7,
		"OfferSequence":      6
	}`

	err := testutil.CompareFlattenAndExpected(tx.Flatten(), []byte(expected))
	if err != nil {
		t.Error(err)
	}
}
