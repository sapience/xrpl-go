package utils

import (
	"testing"
)

func TestIsMemo(t *testing.T) {
	t.Run("Valid Memo object", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"Memo": map[string]interface{}{
				"MemoData":   "Hello World",
				"MemoFormat": "text/plain",
				"MemoType":   "general",
			},
		}
		if !IsMemo(obj1) {
			t.Errorf("Expected IsMemo to return true, but got false")
		}
	})

	t.Run("Memo object with missing fields", func(t *testing.T) {
		obj2 := map[string]interface{}{
			"Memo": map[string]interface{}{
				"MemoData": "Hello World",
			},
		}
		if !IsMemo(obj2) {
			t.Errorf("Expected IsMemo to return true, but got false")
		}
	})

	t.Run("Memo object with invalid field types", func(t *testing.T) {
		obj3 := map[string]interface{}{
			"Memo": map[string]interface{}{
				"MemoData":   12345,
				"MemoFormat": 12345,
				"MemoType":   12345,
			},
		}
		if IsMemo(obj3) {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})

	t.Run("Memo object with extra fields", func(t *testing.T) {
		obj4 := map[string]interface{}{
			"Memo": map[string]interface{}{
				"MemoData":   "Hello World",
				"MemoFormat": "text/plain",
				"MemoType":   "general",
				"ExtraField": "Extra Value",
			},
		}
		if IsMemo(obj4) {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})

	t.Run("Nil object", func(t *testing.T) {
		obj5 := map[string]interface{}{}
		if IsMemo(obj5) {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})
}
func TestIsSigner(t *testing.T) {
	t.Run("Valid Signer object", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"Signer": map[string]interface{}{
				"Account":       "r1234567890",
				"TxnSignature":  "0123456789abcdef",
				"SigningPubKey": "abcdef0123456789",
			},
		}
		if !IsSigner(obj1) {
			t.Errorf("Expected IsSigner to return true, but got false")
		}
	})

	t.Run("Signer object with missing fields", func(t *testing.T) {
		obj2 := map[string]interface{}{
			"Signer": map[string]interface{}{
				"Account": "r1234567890",
			},
		}
		if IsSigner(obj2) {
			t.Errorf("Expected IsSigner to return false, but got true")
		}
	})

	t.Run("Signer object with invalid field types", func(t *testing.T) {
		obj3 := map[string]interface{}{
			"Signer": map[string]interface{}{
				"Account":       12345,
				"TxnSignature":  12345,
				"SigningPubKey": 12345,
			},
		}
		if IsSigner(obj3) {
			t.Errorf("Expected IsSigner to return false, but got true")
		}
	})

	t.Run("Signer object with extra fields", func(t *testing.T) {
		obj4 := map[string]interface{}{
			"Signer": map[string]interface{}{
				"Account":       "r1234567890",
				"TxnSignature":  "0123456789abcdef",
				"SigningPubKey": "abcdef0123456789",
				"ExtraField":    "Extra Value",
			},
		}
		if IsSigner(obj4) {
			t.Errorf("Expected IsSigner to return false, but got true")
		}
	})

	t.Run("Nil object", func(t *testing.T) {
		obj5 := map[string]interface{}{}
		if IsSigner(obj5) {
			t.Errorf("Expected IsSigner to return false, but got true")
		}
	})
}
func TestIsIssuedCurrency(t *testing.T) {
	t.Run("Valid IssuedCurrency object", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"value":    "100",
			"issuer":   "r1234567890",
			"currency": "USD",
		}
		if !IsIssuedCurrency(obj1) {
			t.Errorf("Expected IsIssuedCurrency to return true, but got false")
		}
	})

	t.Run("IssuedCurrency object with missing fields", func(t *testing.T) {
		obj2 := map[string]interface{}{
			"value": "100",
		}
		if IsIssuedCurrency(obj2) {
			t.Errorf("Expected IsIssuedCurrency to return false, but got true")
		}
	})

	t.Run("IssuedCurrency object with invalid field types", func(t *testing.T) {
		obj3 := map[string]interface{}{
			"value":    100,
			"issuer":   12345,
			"currency": 12345,
		}
		if IsIssuedCurrency(obj3) {
			t.Errorf("Expected IsIssuedCurrency to return false, but got true")
		}
	})

	t.Run("IssuedCurrency object with extra fields", func(t *testing.T) {
		obj4 := map[string]interface{}{
			"value":    "100",
			"issuer":   "r1234567890",
			"currency": "USD",
			"extra":    "extra field",
		}
		if IsIssuedCurrency(obj4) {
			t.Errorf("Expected IsIssuedCurrency to return false, but got true")
		}
	})

	t.Run("Nil object", func(t *testing.T) {
		obj5 := map[string]interface{}{}
		if IsIssuedCurrency(obj5) {
			t.Errorf("Expected IsIssuedCurrency to return false, but got true")
		}
	})
}
