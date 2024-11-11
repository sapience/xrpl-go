package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/stretchr/testify/assert"
)

func TestNFTokenBurn_TxType(t *testing.T) {
	tx := &NFTokenBurn{}
	assert.Equal(t, NFTokenBurnTx, tx.TxType())
}

func TestNFTokenBurn_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		nftBurn  *NFTokenBurn
		expected string
	}{
		{
			name: "Without Owner",
			nftBurn: &NFTokenBurn{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenBurnTx,
				},
				NFTokenID: "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65",
			},
			expected: `{
				"Account": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
				"TransactionType": "NFTokenBurn",
				"NFTokenID": "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65"
			}`,
		},
		{
			name: "With Owner",
			nftBurn: &NFTokenBurn{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenBurnTx,
				},
				NFTokenID: "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65",
				Owner:     "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
			},
			expected: `{
				"Account": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
				"TransactionType": "NFTokenBurn",
				"NFTokenID": "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65",
				"Owner": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2"
			}`,
		},
		{
			name: "Without NFTokenID",
			nftBurn: &NFTokenBurn{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenBurnTx,
				},
			},
			expected: `{
				"Account": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
				"TransactionType": "NFTokenBurn"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.nftBurn.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}
