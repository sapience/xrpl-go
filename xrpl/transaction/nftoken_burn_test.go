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

func TestNFTokenBurn_Validate(t *testing.T) {
	tests := []struct {
		name      string
		nftBurn   *NFTokenBurn
		wantValid bool
		wantErr   bool
	}{
		{
			name: "pass - Valid NFTokenBurn without Owner",
			nftBurn: &NFTokenBurn{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenBurnTx,
				},
				NFTokenID: "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65",
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "fail - Invalid BaseTx, missing Account",
			nftBurn: &NFTokenBurn{
				BaseTx: BaseTx{
					TransactionType: NFTokenBurnTx,
				},
				NFTokenID: "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65",
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "pass - Valid NFTokenBurn with Owner",
			nftBurn: &NFTokenBurn{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenBurnTx,
				},
				NFTokenID: "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65",
				Owner:     "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "fail - Invalid NFTokenBurn with invalid Owner",
			nftBurn: &NFTokenBurn{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenBurnTx,
				},
				NFTokenID: "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65",
				Owner:     "invalidOwnerAddress",
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "fail - Invalid NFTokenBurn with invalid NFTokenID",
			nftBurn: &NFTokenBurn{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenBurnTx,
				},
				NFTokenID: "invalidNFTokenID",
			},
			wantValid: false,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.nftBurn.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if valid != tt.wantValid {
				t.Errorf("Validate() valid = %v, wantValid %v", valid, tt.wantValid)
			}
		})
	}
}
