package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestCheckCreate_TxType(t *testing.T) {
	tx := &CheckCreate{}
	assert.Equal(t, CheckCreateTx, tx.TxType())
}

func TestCheckCreate_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		tx       *CheckCreate
		expected FlatTransaction
	}{
		{
			name: "pass - All fields",
			tx: &CheckCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: CheckCreateTx,
				},
				Destination:    "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				SendMax:        types.XRPCurrencyAmount(10000),
				DestinationTag: 23480,
				Expiration:     533257958,
				InvoiceID:      "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
			},
			expected: FlatTransaction{
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "CheckCreate",
				"Destination":     "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"SendMax":         "10000",
				"DestinationTag":  uint32(23480),
				"Expiration":      uint32(533257958),
				"InvoiceID":       "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
			},
		},
		{
			name: "pass - Optional fields omitted",
			tx: &CheckCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: CheckCreateTx,
				},
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				SendMax:     types.XRPCurrencyAmount(10000),
			},
			expected: FlatTransaction{
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "CheckCreate",
				"Destination":     "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"SendMax":         "10000",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.tx.Flatten())
		})
	}
}
