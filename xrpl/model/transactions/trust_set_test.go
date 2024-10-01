package transactions

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/test"
)

func TestTrustSetTx(t *testing.T) {
	s := TrustSet{
		BaseTx: BaseTx{
			Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType:    TrustSetTx,
			Fee:                types.XRPCurrencyAmount(12),
			Flags:              262144,
			Sequence:           12,
			LastLedgerSequence: 8007750,
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Issuer:   "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
			Currency: "USD",
			Value:    "100",
		},
	}

	j := `{
	"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	"TransactionType": "TrustSet",
	"Fee": "12",
	"Sequence": 12,
	"Flags": 262144,
	"LastLedgerSequence": 8007750,
	"LimitAmount": {
		"issuer": "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
		"currency": "USD",
		"value": "100"
	}
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

func TestTrustSetFlatten(t *testing.T) {
	s := TrustSet{
		BaseTx: BaseTx{
			Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType:    TrustSetTx,
			Fee:                types.XRPCurrencyAmount(12),
			Flags:              262144,
			Sequence:           12,
			LastLedgerSequence: 8007750,
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Issuer:   "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
			Currency: "USD",
			Value:    "100",
		},
	}

	flattened := s.Flatten()

	expected := FlatTransaction{
		"Account":            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType":    "TrustSet",
		"Fee":                "12",
		"Flags":              int(262144),
		"Sequence":           uint(12),
		"LastLedgerSequence": uint(8007750),
		"LimitAmount": map[string]interface{}{
			"issuer":   "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
			"currency": "USD",
			"value":    "100",
		},
	}

	// Existing DeepEqual check
	if !reflect.DeepEqual(flattened, expected) {
		t.Errorf("Flatten result differs from expected: %v, %v", flattened, expected)
	}
}

func TestTrustSetFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*TrustSet)
		expected uint
	}{
		{
			name: "SetSetAuthFlag",
			setter: func(ts *TrustSet) {
				ts.SetSetAuthFlag()
			},
			expected: tfSetAuth,
		},
		{
			name: "SetSetNoRippleFlag",
			setter: func(ts *TrustSet) {
				ts.SetSetNoRippleFlag()
			},
			expected: tfSetNoRipple,
		},
		{
			name: "SetClearNoRippleFlag",
			setter: func(ts *TrustSet) {
				ts.SetClearNoRippleFlag()
			},
			expected: tfClearNoRipple,
		},
		{
			name: "SetSetfAuthFlag and SetSetNoRippleFlag",
			setter: func(ts *TrustSet) {
				ts.SetSetAuthFlag()
				ts.SetSetNoRippleFlag()
			},
			expected: tfSetAuth | tfSetNoRipple,
		},
		{
			name: "SetSetfAuthFlag and SetClearNoRippleFlag",
			setter: func(ts *TrustSet) {
				ts.SetSetAuthFlag()
				ts.SetClearNoRippleFlag()
			},
			expected: tfSetAuth | tfClearNoRipple,
		},
		{
			name: "All flags",
			setter: func(ts *TrustSet) {
				ts.SetSetAuthFlag()
				ts.SetSetNoRippleFlag()
				ts.SetClearNoRippleFlag()
			},
			expected: tfSetAuth | tfSetNoRipple | tfClearNoRipple,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TrustSet{}
			tt.setter(ts)
			if ts.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, ts.Flags)
			}
		})
	}
}
func TestTrustSetValidate(t *testing.T) {
	tests := []struct {
		name     string
		trustSet *TrustSet
		valid    bool
		err      error
	}{
		{
			name: "ValidTrustSet",
			trustSet: &TrustSet{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    TrustSetTx,
					Fee:                types.XRPCurrencyAmount(12),
					Flags:              262144,
					Sequence:           12,
					LastLedgerSequence: 8007750,
				},
				LimitAmount: types.IssuedCurrencyAmount{
					Issuer:   "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
					Currency: "USD",
					Value:    "100",
				},
				QualityIn:  100,
				QualityOut: 200,
			},
			valid: true,
			err:   nil,
		},
		{
			name: "MissingLimitAmount",
			trustSet: &TrustSet{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    TrustSetTx,
					Fee:                types.XRPCurrencyAmount(12),
					Flags:              262144,
					Sequence:           12,
					LastLedgerSequence: 8007750,
				},
				QualityIn:  100,
				QualityOut: 200,
			},
			valid: false,
			err:   errors.New("trustSet: missing field LimitAmount"),
		},
		{
			name: "InvalidLimitAmount",
			trustSet: &TrustSet{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    TrustSetTx,
					Fee:                types.XRPCurrencyAmount(12),
					Flags:              262144,
					Sequence:           12,
					LastLedgerSequence: 8007750,
				},
				LimitAmount: types.IssuedCurrencyAmount{
					Issuer:   "r123",
					Currency: "USD",
				},
				QualityIn:  100,
				QualityOut: 200,
			},
			valid: false,
			err:   errors.New("trustSet: invalid LimitAmount"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.trustSet.Validate()
			if valid != tt.valid {
				t.Errorf("Expected valid to be %v, got %v", tt.valid, valid)
			}
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Expected error to be %v, got %v", tt.err, err)
			}
		})
	}
}
