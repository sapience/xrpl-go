package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/stretchr/testify/assert"
)

func TestEscrowFinish_TxType(t *testing.T) {
	entry := &EscrowFinish{}
	assert.Equal(t, EscrowFinishTx, entry.TxType())
}

func TestEscrowFinish_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		entry    *EscrowFinish
		expected string
	}{
		{
			name: "All fields set",
			entry: &EscrowFinish{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowFinishTx,
				},
				Owner:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				OfferSequence: 7,
				Condition:     "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
				Fulfillment:   "A0028000",
			},
			expected: `{
				"TransactionType": "EscrowFinish",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Owner": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"OfferSequence":   7,
				"Condition": "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
				"Fulfillment": "A0028000"
			}`,
		},
		{
			name: "Optional fields omitted",
			entry: &EscrowFinish{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowFinishTx,
				},
				Owner:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				OfferSequence: 7,
			},
			expected: `{
				"TransactionType": "EscrowFinish",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Owner": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"OfferSequence": 7
			}`,
		},
		{
			name: "Only BaseTx fields",
			entry: &EscrowFinish{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowFinishTx,
				},
			},
			expected: `{
				"TransactionType": "EscrowFinish",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn"
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
