package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/test"
)

func TestPaymentTx(t *testing.T) {
	s := Payment{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: PaymentTx,
			Fee:             types.XRPCurrencyAmount(1000),
			Flags:           262144,
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   "def",
			Currency: "USD",
			Value:    "1",
		},
		Destination: "hij",
	}

	j := `{
	"Account": "abc",
	"TransactionType": "Payment",
	"Fee": "1000",
	"Flags": 262144,
	"Amount": {
		"issuer": "def",
		"currency": "USD",
		"value": "1"
	},
	"Destination": "hij"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
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

func TestPaymentFlatten(t *testing.T) {
	s := Payment{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: PaymentTx,
			Fee:             types.XRPCurrencyAmount(1000),
			Flags:           262144,
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			Currency: "USD",
			Value:    "1",
		},
		Destination: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	}

	flattened := s.Flatten()

	expected := map[string]interface{}{
		"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "Payment",
		"Fee":             uint64(1000),
		"Flags":           uint(262144),
		"Amount": map[string]interface{}{
			"issuer":   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			"currency": "USD",
			"value":    "1",
		},
		"Destination": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	}

	if !reflect.DeepEqual(flattened, expected) {
		t.Errorf("Flatten result differs from expected: %v, %v", flattened, expected)
	}
}
