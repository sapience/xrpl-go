package transaction

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/test"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
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

	expected := FlatTransaction{
		"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "Payment",
		"Fee":             "1000",
		"Flags":           int(262144),
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

func TestPaymentFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*Payment)
		expected uint
	}{
		{
			name: "SetRippleNotDirectFlag",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
			},
			expected: tfRippleNotDirect,
		},
		{
			name: "SetPartialPaymentFlag",
			setter: func(p *Payment) {
				p.SetPartialPaymentFlag()
			},
			expected: tfPartialPayment,
		},
		{
			name: "SetLimitQualityFlag",
			setter: func(p *Payment) {
				p.SetLimitQualityFlag()
			},
			expected: tfLimitQuality,
		},
		{
			name: "SetRippleNotDirectFlag and SetPartialPaymentFlag",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
				p.SetPartialPaymentFlag()
			},
			expected: tfRippleNotDirect | tfPartialPayment,
		},
		{
			name: "SetRippleNotDirectFlag and SetLimitQualityFlag",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
				p.SetLimitQualityFlag()
			},
			expected: tfRippleNotDirect | tfLimitQuality,
		},
		{
			name: "SetPartialPaymentFlag and SetLimitQualityFlag",
			setter: func(p *Payment) {
				p.SetPartialPaymentFlag()
				p.SetLimitQualityFlag()
			},
			expected: tfPartialPayment | tfLimitQuality,
		},
		{
			name: "All flags",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
				p.SetPartialPaymentFlag()
				p.SetLimitQualityFlag()
			},
			expected: tfRippleNotDirect | tfPartialPayment | tfLimitQuality,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Payment{}
			tt.setter(p)
			if p.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, p.Flags)
			}
		})
	}
}
