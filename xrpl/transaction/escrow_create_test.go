package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestEscrowCreate_TxType(t *testing.T) {
	entry := &EscrowCreate{}
	assert.Equal(t, EscrowCreateTx, entry.TxType())
}
func TestEscrowCreate_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		entry    *EscrowCreate
		expected string
	}{
		{
			name: "All fields set",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
				Amount:         types.XRPCurrencyAmount(10000),
				Destination:    "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				CancelAfter:    533257958,
				FinishAfter:    533171558,
				Condition:      "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
				DestinationTag: 23480,
			},
			expected: `{
				"TransactionType": "EscrowCreate",
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Amount":          "10000",
				"Destination":     "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"CancelAfter":     533257958,
				"FinishAfter":     533171558,
				"Condition":       "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
				"DestinationTag":  23480
			}`,
		},
		{
			name: "Optional fields omitted",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
			},
			expected: `{
				"TransactionType": "EscrowCreate",
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Amount":          "10000",
				"Destination":     "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.entry.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}
func TestEscrowCreate_Validate(t *testing.T) {
	tests := []struct {
		name     string
		entry    *EscrowCreate
		expected bool
	}{
		{
			name: "Valid transaction",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				CancelAfter: 533257958,
				Condition:   "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
			},
			expected: true,
		},
		{
			name: "Invalid BaseTx, missing TransactionType",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				CancelAfter: 533257958,
			},
			expected: false,
		},
		{
			name: "Valid transaction with FinishAfter",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				FinishAfter: 533171558,
			},
			expected: true,
		},
		{
			name: "Valid transaction with Condition",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				FinishAfter: 533171558,
				Condition:   "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
			},
			expected: true,
		},
		{
			name: "Invalid transaction with no CancelAfter or FinishAfter",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
			},
			expected: false,
		},
		{
			name: "Invalid transaction with no Condition and FinishAfter",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				CancelAfter: 533257958,
			},
			expected: false,
		},
		{
			name: "Invalid transaction with invalid destination address",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "invalidAddress",
				CancelAfter: 533257958,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, _ := tt.entry.Validate()
			assert.Equal(t, tt.expected, valid)
		})
	}
}
