package transaction

import (
	"testing"
	"time"

	"github.com/Peersyst/xrpl-go/xrpl"
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
func TestPaymentChannelFund_Validate(t *testing.T) {
	tests := []struct {
		name             string
		tx               *PaymentChannelFund
		expirationSetter func(tx *PaymentChannelFund)
		wantValid        bool
		wantErr          bool
		expectedErr      error
	}{
		{
			name: "pass - Valid Transaction",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelFundTx,
				},
				Channel: "ABC123",
				Amount:  types.XRPCurrencyAmount(200000),
			},
			expirationSetter: func(tx *PaymentChannelFund) {
				tx.Expiration = uint32(time.Now().Unix()) + 5000
			},
			wantValid:   true,
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name: "fail - Valid Transaction",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					TransactionType: PaymentChannelFundTx,
				},
				Channel: "ABC123",
				Amount:  types.XRPCurrencyAmount(200000),
			},
			expirationSetter: func(tx *PaymentChannelFund) {
				tx.Expiration = uint32(xrpl.UnixTimeToRippleTime(time.Now().Unix()) + 5000)
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidAccount,
		},
		{
			name: "fail - Invalid Expiration",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelFundTx,
				},
				Channel:    "DEF456",
				Amount:     types.XRPCurrencyAmount(300000),
				Expiration: 1, // Invalid expiration time
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidExpiration,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.tx.Validate()
			if tt.expirationSetter != nil {
				tt.expirationSetter(tt.tx)
			}

			if valid != tt.wantValid {
				t.Errorf("Validate() valid = %v, want %v", valid, tt.wantValid)
			}
			if (err != nil) && err != tt.expectedErr {
				t.Errorf("Validate() got error message = %v, want error message %v", err, tt.expectedErr)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
