package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestPaymentChannelCreate_TxType(t *testing.T) {
	tx := &PaymentChannelCreate{}
	assert.Equal(t, PaymentChannelCreateTx, tx.TxType())
}

func TestPaymentChannelCreate_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		tx       *PaymentChannelCreate
		expected string
	}{
		{
			name: "pass - All fields set",
			tx: &PaymentChannelCreate{
				BaseTx: BaseTx{
					Account:         "r2UeJh4HhYc5VtYc8U2YpZfQzY5Lw8kZV",
					TransactionType: PaymentChannelCreateTx,
				},
				Amount:         types.XRPCurrencyAmount(10000),
				Destination:    types.Address("rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn"),
				SettleDelay:    86400,
				PublicKey:      "32D2471DB72B27E3310F355BB33E339BF26F8392D5A93D3BC0FC3B566612DA0F0A",
				CancelAfter:    533171558,
				DestinationTag: 23480,
			},
			expected: `{
				"Account":       "r2UeJh4HhYc5VtYc8U2YpZfQzY5Lw8kZV",
				"TransactionType": "PaymentChannelCreate",
				"Amount":        "10000",
				"Destination":   "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"SettleDelay":   86400,
				"PublicKey":     "32D2471DB72B27E3310F355BB33E339BF26F8392D5A93D3BC0FC3B566612DA0F0A",
				"CancelAfter":   533171558,
				"DestinationTag": 23480
			}`,
		},
		{
			name: "pass - Optional fields omitted",
			tx: &PaymentChannelCreate{
				BaseTx: BaseTx{
					Account:         "r2UeJh4HhYc5VtYc8U2YpZfQzY5Lw8kZV",
					TransactionType: PaymentChannelCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: types.Address("rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn"),
				SettleDelay: 86400,
				PublicKey:   "32D2471DB72B27E3310F355BB33E339BF26F8392D5A93D3BC0FC3B566612DA0F0A",
			},
			expected: `{
				"Account":     "r2UeJh4HhYc5VtYc8U2YpZfQzY5Lw8kZV",
				"TransactionType": "PaymentChannelCreate",
				"Amount":      "10000",
				"Destination": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"SettleDelay": 86400,
				"PublicKey":   "32D2471DB72B27E3310F355BB33E339BF26F8392D5A93D3BC0FC3B566612DA0F0A"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.tx.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}
