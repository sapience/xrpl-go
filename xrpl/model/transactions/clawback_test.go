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

func TestClawbackFlatten(t *testing.T) {
	s := Clawback{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: ClawbackTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			Currency: "USD",
			Value:    "1",
		},
	}

	flattened := s.Flatten()

	expected := FlatTransaction{
		"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "Clawback",
		"Fee":             "1",
		"Sequence":        uint(1234),
		"Amount": map[string]interface{}{
			"issuer":   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			"currency": "USD",
			"value":    "1",
		},
	}

	if !reflect.DeepEqual(flattened, expected) {
		t.Errorf("Flatten result differs from expected: %v, %v", flattened, expected)
	}
}
func TestClawbackValidate(t *testing.T) {
	// Valid Clawback transaction
	validClawback := Clawback{
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

	valid, err := validClawback.Validate()
	if err != nil {
		t.Errorf("Validation failed for valid Clawback transaction: %s", err.Error())
	}
	if !valid {
		t.Error("Validation should pass for valid Clawback transaction")
	}

	// Clawback transaction with missing Amount field
	missingAmount := Clawback{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: ClawbackTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
	}

	valid, err = missingAmount.Validate()
	if err == nil || valid {
		t.Error("Validation should fail for Clawback transaction with missing Amount field")
	}

	// Clawback transaction with invalid Amount
	invalidAmount := Clawback{
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
			Value:    "invalid",
		},
	}

	valid, err = invalidAmount.Validate()
	if err == nil || valid {
		t.Error("Validation should fail for Clawback transaction with invalid Amount")
	}

	// Clawback transaction with Account same as the issuer
	invalidAccount := Clawback{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: ClawbackTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   "abcdef",
			Currency: "USD",
			Value:    "1",
		},
	}

	valid, err = invalidAccount.Validate()
	if err == nil || valid {
		t.Error("Validation should fail for Clawback transaction with Account same as the issuer")
	}
}
