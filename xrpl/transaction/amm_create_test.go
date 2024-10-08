package transaction

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestAMMCreateTransaction(t *testing.T) {
	s := AMMCreate{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: AMMCreateTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Amount: types.XRPCurrencyAmount(100),
		Amount2: types.IssuedCurrencyAmount{
			Currency: "USD",
			Value:    "200",
			Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
		},
		TradingFee: 10,
	}

	j := `{
	"Account": "abcdef",
	"TransactionType": "AMMCreate",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"Amount": "100",
	"Amount2": {
		"issuer": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
		"currency": "USD",
		"value": "200"
	},
	"TradingFee": 10
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

	tx, err := UnmarshalTx(json.RawMessage(j))
	if err != nil {
		t.Errorf("UnmarshalTx error: %s", err.Error())
	}
	if !reflect.DeepEqual(tx, &s) {
		t.Error("UnmarshalTx result differs from expected")
	}
}
func TestAMMCreateFlatten(t *testing.T) {
	s := AMMCreate{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: AMMCreateTx,
			Fee:             types.XRPCurrencyAmount(10),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Amount: types.XRPCurrencyAmount(100),
		Amount2: types.IssuedCurrencyAmount{
			Currency: "USD",
			Value:    "200",
			Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
		},
		TradingFee: 10,
	}

	flattened := s.Flatten()

	expected := `{
		"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "AMMCreate",
		"Fee":             "10",
		"Sequence":        1234,
		"SigningPubKey":   "ghijk",
		"TxnSignature":    "A1B2C3D4E5F6",
		"Amount":          "100",
		"Amount2":         {
			"currency": "USD",
			"value":    "200",
			"issuer":   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"
		},
		"TradingFee":      10
	}`

	// Convert flattened to JSON
	flattenedJSON, err := json.Marshal(flattened)
	if err != nil {
		t.Errorf("Error marshaling payment flattened, error: %v", err)
	}

	// Normalize expected JSON
	var expectedMap map[string]interface{}
	if err := json.Unmarshal([]byte(expected), &expectedMap); err != nil {
		t.Errorf("Error unmarshaling expected, error: %v", err)
	}
	expectedJSON, err := json.Marshal(expectedMap)
	if err != nil {
		t.Errorf("Error marshaling expected payment object: %v", err)
	}

	// Compare JSON strings
	if string(flattenedJSON) != string(expectedJSON) {
		t.Errorf("The flattened and expected Payment JSON are not equal.\nGot:      %v\nExpected: %v", string(flattenedJSON), string(expectedJSON))
	}
}
