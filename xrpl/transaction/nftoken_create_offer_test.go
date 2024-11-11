package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestNFTokenCreateOffer_TxType(t *testing.T) {
	tx := &NFTokenCreateOffer{}
	assert.Equal(t, NFTokenCreateOfferTx, tx.TxType())
}

func TestNFTokenCreateOffer_Flags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*NFTokenCreateOffer)
		expected uint
	}{
		{
			name: "pass - SetSellNFTokenFlag",
			setter: func(n *NFTokenCreateOffer) {
				n.SetSellNFTokenFlag()
			},
			expected: tfSellNFToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NFTokenCreateOffer{}
			tt.setter(n)
			if n.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, n.Flags)
			}
		})
	}
}

func TestNFTokenCreateOffer_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		input    *NFTokenCreateOffer
		expected string
	}{
		{
			name: "all fields set",
			input: &NFTokenCreateOffer{
				BaseTx: BaseTx{
					Account:         "rs8jBmmfpwgmrSPgwMsh7CvKRmRt1JTVSX",
					TransactionType: NFTokenCreateOfferTx,
				},
				Owner:       "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				NFTokenID:   "000100001E962F495F07A990F4ED55ACCFEEF365DBAA76B6A048C0A200000007",
				Amount:      types.XRPCurrencyAmount(1000000),
				Expiration:  600000000,
				Destination: "r3G8r9hV1J8r9hV1J8r9hV1J8r9hV1J8r9",
			},
			expected: `{
				"TransactionType": "NFTokenCreateOffer",
				"Account": "rs8jBmmfpwgmrSPgwMsh7CvKRmRt1JTVSX",
				"Owner": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				"NFTokenID": "000100001E962F495F07A990F4ED55ACCFEEF365DBAA76B6A048C0A200000007",
				"Amount": "1000000",
				"Expiration": 600000000,
				"Destination": "r3G8r9hV1J8r9hV1J8r9hV1J8r9hV1J8r9"
			}`,
		},
		{
			name: "optional fields omitted",
			input: &NFTokenCreateOffer{
				BaseTx: BaseTx{
					Account:         "rs8jBmmfpwgmrSPgwMsh7CvKRmRt1JTVSX",
					TransactionType: NFTokenCreateOfferTx,
				},
				NFTokenID: "000100001E962F495F07A990F4ED55ACCFEEF365DBAA76B6A048C0A200000007",
				Amount:    types.XRPCurrencyAmount(1000000),
			},
			expected: `{
				"TransactionType": "NFTokenCreateOffer",
				"Account": "rs8jBmmfpwgmrSPgwMsh7CvKRmRt1JTVSX",
				"NFTokenID": "000100001E962F495F07A990F4ED55ACCFEEF365DBAA76B6A048C0A200000007",
				"Amount": "1000000"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.input.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}
