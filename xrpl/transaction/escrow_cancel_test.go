package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/stretchr/testify/assert"
)

func TestEscrowCancel_TxType(t *testing.T) {
	entry := &EscrowCancel{}
	assert.Equal(t, EscrowCancelTx, entry.TxType())
}

func TestEscrowCancel_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		escrow   *EscrowCancel
		expected string
	}{
		{
			name: "complete EscrowCancel",
			escrow: &EscrowCancel{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
				Owner:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				OfferSequence: 7,
			},
			expected: `{
				"TransactionType": "EscrowCancel",
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Owner":           "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"OfferSequence":   7
			}`,
		},
		{
			name: "EscrowCancel without Owner",
			escrow: &EscrowCancel{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
				OfferSequence: 7,
			},
			expected: `{
				"TransactionType": "EscrowCancel",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"OfferSequence": 7
			}`,
		},
		{
			name: "EscrowCancel without OfferSequence",
			escrow: &EscrowCancel{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
				Owner: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			},
			expected: `{
				"TransactionType": "EscrowCancel",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Owner": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn"
			}`,
		},
		{
			name: "EscrowCancel with only BaseTx",
			escrow: &EscrowCancel{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
			},
			expected: `{
				"TransactionType": "EscrowCancel",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.escrow.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestEscrowCancel_Validate(t *testing.T) {
	tests := []struct {
		name     string
		escrow   *EscrowCancel
		expected bool
	}{
		{
			name: "valid EscrowCancel",
			escrow: &EscrowCancel{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCancelTx,
				},
				Owner:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				OfferSequence: 7,
			},
			expected: true,
		},
		{
			name: "Invalid EscrowCancel BaseTx",
			escrow: &EscrowCancel{
				BaseTx: BaseTx{
					TransactionType: EscrowCancelTx,
				},
				Owner:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				OfferSequence: 7,
			},
			expected: false,
		},
		{
			name: "invalid Owner address",
			escrow: &EscrowCancel{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCancelTx,
				},
				Owner:         "invalidAddress",
				OfferSequence: 7,
			},
			expected: false,
		},
		{
			name: "missing OfferSequence",
			escrow: &EscrowCancel{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCancelTx,
				},
				Owner: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, _ := tt.escrow.Validate()
			assert.Equal(t, tt.expected, valid)
		})
	}
}
