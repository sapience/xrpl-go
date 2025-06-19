package types

import "testing"

func TestIssuedCurrencyAmount_IsZero(t *testing.T) {
	tests := []struct {
		name string
		ica  IssuedCurrencyAmount
		want bool
	}{
		{
			name: "Zero value",
			ica:  IssuedCurrencyAmount{},
			want: true,
		},
		{
			name: "Non-zero value",
			ica: IssuedCurrencyAmount{
				Issuer:   "rEXAMPLE",
				Currency: "USD",
				Value:    "100",
			},
			want: false,
		},
		{
			name: "Non-zero value, invalid only with issuer",
			ica: IssuedCurrencyAmount{
				Issuer: "rEXAMPLE",
			},
			want: false,
		},
		{
			name: "Non-zero value, invalid only with value",
			ica: IssuedCurrencyAmount{
				Value: "100",
			},
			want: false,
		},
		{
			name: "Non-zero value, invalid only with currency",
			ica: IssuedCurrencyAmount{
				Currency: "USD",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ica.IsZero(); got != tt.want {
				t.Errorf("IssuedCurrencyAmount.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMPTCurrencyAmount_Kind(t *testing.T) {
	mpt := MPTCurrencyAmount{}
	if kind := mpt.Kind(); kind != MPT {
		t.Errorf("MPTCurrencyAmount.Kind() = %v, want %v", kind, MPT)
	}
}

func TestMPTCurrencyAmount_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		mpt      MPTCurrencyAmount
		expected map[string]interface{}
	}{
		{
			name:     "Empty MPT amount",
			mpt:      MPTCurrencyAmount{},
			expected: map[string]interface{}{},
		},
		{
			name: "With MPT issuance ID only",
			mpt: MPTCurrencyAmount{
				MPTIssuanceID: "00000000000000000000000000000000",
			},
			expected: map[string]interface{}{
				"mpt_issuance_id": "00000000000000000000000000000000",
			},
		},
		{
			name: "With value only",
			mpt: MPTCurrencyAmount{
				Value: "100",
			},
			expected: map[string]interface{}{
				"value": "100",
			},
		},
		{
			name: "With both issuance ID and value",
			mpt: MPTCurrencyAmount{
				MPTIssuanceID: "00000000000000000000000000000000",
				Value:         "100",
			},
			expected: map[string]interface{}{
				"mpt_issuance_id": "00000000000000000000000000000000",
				"value":           "100",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.mpt.Flatten()
			flattened, ok := result.(map[string]interface{})
			if !ok {
				t.Fatalf("Expected map[string]interface{}, got %T", result)
			}

			if len(flattened) != len(tt.expected) {
				t.Errorf("Expected map of length %d, got %d", len(tt.expected), len(flattened))
			}

			for key, expectedValue := range tt.expected {
				actualValue, exists := flattened[key]
				if !exists {
					t.Errorf("Expected key %q not found in flattened output", key)
				} else if actualValue != expectedValue {
					t.Errorf("For key %q, expected value %v, got %v", key, expectedValue, actualValue)
				}
			}
		})
	}
}

func TestUnmarshalCurrencyAmount_MPT(t *testing.T) {
	raw := `{"mpt_issuance_id":"issuance","value":"42"}`
	amt, err := UnmarshalCurrencyAmount([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}
	m, ok := amt.(MPTCurrencyAmount)
	if !ok {
		t.Fatalf("expected MPTCurrencyAmount, got %T", amt)
	}
	if m.MPTIssuanceID != "issuance" || m.Value != "42" {
		t.Errorf("unmarshaled %#v", m)
	}
}
