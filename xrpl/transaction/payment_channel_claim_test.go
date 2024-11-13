package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestPaymentChannelClaim_TxType(t *testing.T) {
	tx := &PaymentChannelClaim{}
	assert.Equal(t, PaymentChannelClaimTx, tx.TxType())
}

func TestPaymentChannelClaimFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*PaymentChannelClaim)
		expected uint32
	}{
		{
			name: "pass - SetRenewFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetRenewFlag()
			},
			expected: tfRenew,
		},
		{
			name: "pass - SetCloseFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetCloseFlag()
			},
			expected: tfClose,
		},
		{
			name: "pass - SetRenewFlag and SetCloseFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetRenewFlag()
				p.SetCloseFlag()
			},
			expected: tfRenew | tfClose,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaymentChannelClaim{}
			tt.setter(p)
			if p.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, p.Flags)
			}
		})
	}
}

func TestPaymentChannelClaim_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		claim    PaymentChannelClaim
		expected string
	}{
		{
			name: "pass - PaymentChannelClaim with Channel",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel: types.Hash256("ABC123"),
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelClaim",
				"Channel": "ABC123"
			}`,
		},
		{
			name: "PaymentChannelClaim with Balance and Amount",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Balance: types.XRPCurrencyAmount(1000),
				Amount:  types.XRPCurrencyAmount(2000),
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelClaim",
				"Balance": "1000",
				"Amount": "2000"
			}`,
		},
		{
			name: "PaymentChannelClaim with Signature and PublicKey",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Signature: "ABCDEF",
				PublicKey: "123456",
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelClaim",
				"Signature": "ABCDEF",
				"PublicKey": "123456"
			}`,
		},
		{
			name: "PaymentChannelClaim with all fields",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel:   types.Hash256("ABC123"),
				Balance:   types.XRPCurrencyAmount(1000),
				Amount:    types.XRPCurrencyAmount(2000),
				Signature: "ABCDEF",
				PublicKey: "123456",
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelClaim",
				"Channel": "ABC123",
				"Balance": "1000",
				"Amount": "2000",
				"Signature": "ABCDEF",
				"PublicKey": "123456"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.claim.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestPaymentChannelClaim_Validate(t *testing.T) {
	tests := []struct {
		name        string
		claim       PaymentChannelClaim
		wantValid   bool
		wantErr     bool
		expectedErr error
	}{
		{
			name: "pass - all fields valid",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Balance:   types.XRPCurrencyAmount(1000),
				Channel:   "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				Signature: "ABCDEF",
				PublicKey: "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "fail - missing Account in BaseTx",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					TransactionType: PaymentChannelClaimTx,
				},
				Balance:   types.XRPCurrencyAmount(1000),
				Channel:   "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				Signature: "ABCDEF",
				PublicKey: "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidAccount,
		},
		{
			name: "fail - empty Channel",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidChannel,
		},
		{
			name: "fail - invalid Signature",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel:   "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				Signature: "INVALID_SIGNATURE",
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidSignature,
		},
		{
			name: "pass - no Signature",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel: "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "fail - invalid PublicKey",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel:   "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				PublicKey: "INVALID",
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidPublicKey,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.claim.Validate()
			if valid != tt.wantValid {
				t.Errorf("Validate() valid = %v, want %v", valid, tt.wantValid)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err != tt.expectedErr {
				t.Errorf("Validate() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}
