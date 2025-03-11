package transaction

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestCredentialAccept_TxType(t *testing.T) {
	tx := &CredentialAccept{}
	assert.Equal(t, CredentialAcceptTx, tx.TxType())
}

func TestCredentialAccept_Flatten(t *testing.T) {
	s := CredentialAccept{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: CredentialAcceptTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
		},
		Issuer:         "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
		CredentialType: "6D795F63726564656E7469616C", // "my_credential" in hex
	}

	flattened := s.Flatten()

	expected := FlatTransaction{
		"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "CredentialAccept",
		"Fee":             "1",
		"Sequence":        uint32(1234),
		"Issuer":          "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
		"CredentialType":  "6D795F63726564656E7469616C",
	}

	if !reflect.DeepEqual(flattened, expected) {
		t.Errorf("Flatten result differs from expected: %v, %v", flattened, expected)
	}
}

func TestCredentialAccept_Validate(t *testing.T) {
	tests := []struct {
		name     string
		input    *CredentialAccept
		expected bool
	}{
		{
			name: "pass - valid CredentialAccept",
			input: &CredentialAccept{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: "AMMWithdraw",
					Fee:             types.XRPCurrencyAmount(10),
					Flags:           1048576,
					Sequence:        10,
				},
				Issuer:         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				CredentialType: "6D795F63726564656E7469616C",
			},
			expected: true,
		},
		{
			name: "fail - CredentialAccept with an invalid Issuer",
			input: &CredentialAccept{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: "AMMWithdraw",
					Fee:             types.XRPCurrencyAmount(10),
					Flags:           1048576,
					Sequence:        10,
				},
				Issuer:         "invalid_address",
				CredentialType: types.CredentialType("6D795F63726564656E7469616C"),
			},
			expected: false,
		},
		{
			name: "fail - CredentialAccept with an invalid CredentialType (empty)",
			input: &CredentialAccept{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: "AMMWithdraw",
					Fee:             types.XRPCurrencyAmount(10),
					Flags:           1048576,
					Sequence:        10,
				},
				Issuer:         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				CredentialType: types.CredentialType(""),
			},
			expected: false,
		},
		{
			name: "fail - CredentialAccept with an invalid CredentialType (not hex)",
			input: &CredentialAccept{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: "AMMWithdraw",
					Fee:             types.XRPCurrencyAmount(10),
					Flags:           1048576,
					Sequence:        10,
				},
				Issuer:         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				CredentialType: types.CredentialType("not hexadecimal value"),
			},
			expected: false,
		},
		{
			name: "fail - CredentialCreate with an invalid CredentialType (too long)",
			input: &CredentialAccept{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: "AMMWithdraw",
					Fee:             types.XRPCurrencyAmount(10),
					Flags:           1048576,
					Sequence:        10,
				},
				Issuer:         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				CredentialType: types.CredentialType(strings.Repeat("0", common.MaxCredentialTypeLength+1)),
			},
			expected: false,
		},
		{
			name: "fail - CredentialCreate with an invalid CredentialType (too short)",
			input: &CredentialAccept{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: "AMMWithdraw",
					Fee:             types.XRPCurrencyAmount(10),
					Flags:           1048576,
					Sequence:        10,
				},
				Issuer:         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				CredentialType: types.CredentialType(strings.Repeat("0", common.MinCredentialTypeLength-1)),
			},
			expected: false,
		},
		{
			name: "fail - CredentialAccept with an invalid BaseTx",
			input: &CredentialAccept{
				BaseTx: BaseTx{
					Account:         "invalid",
					TransactionType: "AMMWithdraw",
					Fee:             types.XRPCurrencyAmount(10),
					Flags:           1048576,
					Sequence:        10,
				},
				Issuer:         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				CredentialType: types.CredentialType("6D795F63726564656E7469616C"),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.input.Validate()
			if valid != tt.expected {
				t.Errorf("Expected validation result to be %v, got %v", tt.expected, valid)
			}
			if (err != nil) != !tt.expected {
				t.Errorf("Expected error presence to be %v, got %v", !tt.expected, err != nil)
			}
		})
	}
}
