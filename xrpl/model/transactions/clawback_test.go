package transactions

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/test"
)

func TestClawbackTransaction(t *testing.T) {
	clawbackTx := Clawback{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: ClawbackTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   "def",
			Currency: "USD",
			Value:    "1",
		},
	}

	clawbackJSON := `{
	"Account": "abcdef",
	"TransactionType": "Clawback",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"Amount": {
		"issuer": "def",
		"currency": "USD",
		"value": "1"
	}
}`

	if err := test.SerializeAndDeserialize(t, clawbackTx, clawbackJSON); err != nil {
		t.Error(err)
	}

	tx, err := UnmarshalTx(json.RawMessage(clawbackJSON))
	if err != nil {
		t.Errorf("UnmarshalTx error: %s", err.Error())
	}
	if !reflect.DeepEqual(tx, &clawbackTx) {
		fmt.Println("Expected: ", clawbackTx)
		fmt.Println("Actual: ", tx)
		t.Error("UnmarshalTx result differs from expected")
	}
}
