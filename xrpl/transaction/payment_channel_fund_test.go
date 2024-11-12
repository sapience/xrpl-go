package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestPaymentChannelFund_TxType(t *testing.T) {
	tx := &PaymentChannelFund{}
	assert.Equal(t, PaymentChannelFundTx, tx.TxType())
}

func TestPaymentChannelFund_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		tx       *PaymentChannelFund
		expected string
	}{
		{
			name: "pass - Without Expiration",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelFundTx,
				},
				Channel: "ABC123",
				Amount:  types.XRPCurrencyAmount(200000),
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelFund",
				"Channel": "ABC123",
				"Amount":  "200000"
			}`,
		},
		{
			name: "With Expiration",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelFundTx,
				},
				Channel:    "DEF456",
				Amount:     types.XRPCurrencyAmount(300000),
				Expiration: 543171558,
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelFund",
				"Channel": "DEF456",
				"Amount": "300000",
				"Expiration": 543171558
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
