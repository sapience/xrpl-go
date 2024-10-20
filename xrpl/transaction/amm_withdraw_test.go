package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAMMWithdraw_TxType(t *testing.T) {
	tx := &AMMWithdraw{}
	assert.Equal(t, AMMWithdrawTx, tx.TxType())
}

func TestAMMWithdraw_Flags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*AMMWithdraw)
		expected uint
	}{
		{
			name: "SetLPTokentFlag",
			setter: func(a *AMMWithdraw) {
				a.SetLPTokentFlag()
			},
			expected: tfLPToken,
		},
		{
			name: "SetWithdrawAllFlag",
			setter: func(a *AMMWithdraw) {
				a.SetWithdrawAllFlag()
			},
			expected: tfWithdrawAll,
		},
		{
			name: "SetOneAssetWithdrawAllFlag",
			setter: func(a *AMMWithdraw) {
				a.SetOneAssetWithdrawAllFlag()
			},
			expected: tfOneAssetWithdrawAll,
		},
		{
			name: "SetSingleAssetFlag",
			setter: func(a *AMMWithdraw) {
				a.SetSingleAssetFlag()
			},
			expected: tfSingleAsset,
		},
		{
			name: "SetTwoAssetFlag",
			setter: func(a *AMMWithdraw) {
				a.SetTwoAssetFlag()
			},
			expected: tfTwoAsset,
		},
		{
			name: "SetOneAssetLPTokenFlag",
			setter: func(a *AMMWithdraw) {
				a.SetOneAssetLPTokenFlag()
			},
			expected: tfOneAssetLPToken,
		},
		{
			name: "SetLimitLPTokenFlag",
			setter: func(a *AMMWithdraw) {
				a.SetLimitLPTokenFlag()
			},
			expected: tfLimitLPToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AMMWithdraw{}
			tt.setter(a)
			if a.Flags != tt.expected {
				t.Errorf("Expected AMMWithdraw Flags to be %d, got %d", tt.expected, a.Flags)
			}
		})
	}
}
