package results

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTxResult_String(t *testing.T) {
	tests := []struct {
		name     string
		txResult TxResult
		expected string
	}{
		{
			name:     "TesSUCCESS",
			txResult: TesSUCCESS,
			expected: "tesSUCCESS",
		},
		{
			name:     "TecEXPIRED",
			txResult: TecEXPIRED,
			expected: "tecEXPIRED",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.txResult.String()
			require.Equal(t, test.expected, result)
		})
	}
}
