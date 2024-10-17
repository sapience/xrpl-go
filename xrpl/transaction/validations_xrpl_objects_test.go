package transaction

import (
	"testing"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestIsSigner(t *testing.T) {
	t.Run("Valid Signer object", func(t *testing.T) {
		validSigner := SignerData{
			Account:       "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
			TxnSignature:  "0123456789abcdef",
			SigningPubKey: "abcdef0123456789",
		}

		if ok, err := IsSigner(validSigner); !ok {
			t.Errorf("Expected IsSigner to return true, but got false with error: %v", err)
		}
	})

	t.Run("Signer object with missing fields", func(t *testing.T) {
		invalidSigner := SignerData{
			Account:       "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
			SigningPubKey: "abcdef0123456789",
		}

		if ok, _ := IsSigner(invalidSigner); ok {
			t.Errorf("Expected IsSigner to return false, but got true")
		}
	})

	t.Run("Nil object", func(t *testing.T) {
		invalidSigner := SignerData{}
		if ok, _ := IsSigner(invalidSigner); ok {
			t.Errorf("Expected IsSigner to return false, but got true")
		}
	})
}
func TestIsIssuedCurrency(t *testing.T) {
	tests := []struct {
		name     string
		input    types.IssuedCurrencyAmount
		expected bool
	}{
		{
			name: "Valid IssuedCurrency object",
			input: types.IssuedCurrencyAmount{
				Value:    "100",
				Issuer:   "r1234567890",
				Currency: "USD",
			},
			expected: true,
		},
		{
			name: "IssuedCurrency object with missing currency and issuer fields",
			input: types.IssuedCurrencyAmount{
				Value: "100",
			},
			expected: false,
		},
		{
			name: "IssuedCurrency object with missing issuer and value fields",
			input: types.IssuedCurrencyAmount{
				Currency: "USD",
			},
			expected: false,
		},
		{
			name: "IssuedCurrency object with missing currency and value fields",
			input: types.IssuedCurrencyAmount{
				Issuer: "r1234567890",
			},
			expected: false,
		},
		{
			name:     "Empty object",
			input:    types.IssuedCurrencyAmount{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ok, err := IsIssuedCurrency(tt.input); ok != tt.expected {
				t.Errorf("Expected IsIssuedCurrency to return %v, but got %v with error: %v", tt.expected, ok, err)
			}
		})
	}
}

func TestIsMemo(t *testing.T) {
	t.Run("Valid Memo object with all fields", func(t *testing.T) {
		obj := Memo{
			MemoData:   "0123456789abcdef",
			MemoFormat: "abcdef0123456789",
			MemoType:   "abcdef0123456789",
		}

		ok, _ := IsMemo(obj)

		if !(ok) {
			t.Errorf("Expected IsMemo to return true, but got false")
		}
	})

	t.Run("Valid memo object with missing fields", func(t *testing.T) {
		obj := Memo{
			MemoData: "0123456789abcdef",
		}

		ok, err := IsMemo(obj)

		if !ok {
			t.Errorf("Expected IsMemo to return true, but got false with error: %v", err)
		}
	})

	t.Run("Memo object with MemoData non hex value", func(t *testing.T) {
		obj := Memo{
			MemoData: "bob",
		}

		if ok, _ := IsMemo(obj); ok {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})

	t.Run("Memo object with MemoFormat non hex value", func(t *testing.T) {
		obj := Memo{
			MemoData:   "0123456789abcdef",
			MemoFormat: "non-hex",
		}

		if ok, _ := IsMemo(obj); ok {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})

	t.Run("Memo object with MemoType non hex value", func(t *testing.T) {
		obj := Memo{
			MemoData:   "0123456789abcdef",
			MemoFormat: "0123456789abcdef",
			MemoType:   "non-hex",
		}

		if ok, _ := IsMemo(obj); ok {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})

	t.Run("Empty object", func(t *testing.T) {
		obj := Memo{}
		if ok, _ := IsMemo(obj); ok {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})
}
func TestIsAsset(t *testing.T) {
	t.Run("Valid Asset object with currency XRP only", func(t *testing.T) {
		obj := ledger.Asset{
			Currency: "xrP", // will be converted to XRP in the Validate function
		}

		ok, err := IsAsset(obj)

		if !ok {
			t.Errorf("Expected IsAsset to return true, but got false with error: %v", err)
		}
	})

	t.Run("Invalid Asset object with currency only and different than XRP", func(t *testing.T) {
		obj := ledger.Asset{
			Currency: "USD", // missing issuer
		}

		ok, err := IsAsset(obj)

		if ok {
			t.Errorf("Expected IsAsset to return true, but got false with error: %v", err)
		}
	})

	t.Run("Valid Asset object with currency and issuer", func(t *testing.T) {
		obj := ledger.Asset{
			Currency: "USD",
			Issuer:   "r1234567890",
		}

		ok, err := IsAsset(obj)

		if !ok {
			t.Errorf("Expected IsAsset to return true, but got false with error: %v", err)
		}
	})

	t.Run("Asset object with missing currency", func(t *testing.T) {
		obj := ledger.Asset{
			Issuer: "r1234567890",
		}

		ok, err := IsAsset(obj)

		if ok {
			t.Errorf("Expected IsAsset to return false, but got true")
		} else if err == nil {
			t.Errorf("Expected an error, but got nil")
		}
	})

	t.Run("Empty Asset object", func(t *testing.T) {
		obj := ledger.Asset{}

		ok, err := IsAsset(obj)

		if ok {
			t.Errorf("Expected IsAsset to return false, but got true")
		} else if err == nil {
			t.Errorf("Expected an error, but got nil")
		}
	})
}
