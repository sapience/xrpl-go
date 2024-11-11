package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestNFTokenMint_TxType(t *testing.T) {
	tx := &NFTokenMint{}
	assert.Equal(t, NFTokenMintTx, tx.TxType())
}

func TestNFTokenMint_Flags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*NFTokenMint)
		expected uint
	}{
		{
			name: "pass - SetBurnableFlag",
			setter: func(n *NFTokenMint) {
				n.SetBurnableFlag()
			},
			expected: tfBurnable,
		},
		{
			name: "pass - SetOnlyXRPFlag",
			setter: func(n *NFTokenMint) {
				n.SetOnlyXRPFlag()
			},
			expected: tfOnlyXRP,
		},
		{
			name: "pass - SetTrustlineFlag",
			setter: func(n *NFTokenMint) {
				n.SetTrustlineFlag()
			},
			expected: tfTrustLine,
		},
		{
			name: "pass - SetTransferableFlag",
			setter: func(n *NFTokenMint) {
				n.SetTransferableFlag()
			},
			expected: tfTransferable,
		},
		{
			name: "pass - SetBurnableFlag and SetTransferableFlag",
			setter: func(n *NFTokenMint) {
				n.SetBurnableFlag()
				n.SetTransferableFlag()
			},
			expected: tfBurnable | tfTransferable,
		},
		{
			name: "pass - SetBurnableFlag and SetTransferableFlag and SetOnlyXRPFlag",
			setter: func(n *NFTokenMint) {
				n.SetBurnableFlag()
				n.SetTransferableFlag()
				n.SetOnlyXRPFlag()
			},
			expected: tfBurnable | tfTransferable | tfOnlyXRP,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &NFTokenMint{}
			tt.setter(p)
			if p.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, p.Flags)
			}
		})
	}
}

func TestNFTokenMint_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		nft      *NFTokenMint
		expected string
	}{
		{
			name: "Flatten with all fields",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Fee:     types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
				Issuer:       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				TransferFee:  314,
				URI:          "697066733A2F2F62616679626569676479727A74357366703775646D37687537367568377932366E6634646675796C71616266336F636C67747179353566627A6469",
			},
			expected: `{
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "NFTokenMint",
				"Fee":             "10",
				"NFTokenTaxon":    12345,
				"Issuer":          "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"TransferFee":     314,
				"URI":             "697066733A2F2F62616679626569676479727A74357366703775646D37687537367568377932366E6634646675796C71616266336F636C67747179353566627A6469"
			}`,
		},
		{
			name: "Flatten with minimal fields",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: NFTokenMintTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
			},
			expected: `{
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "NFTokenMint",
				"Fee":             "10",
				"NFTokenTaxon":    12345
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					err := testutil.CompareFlattenAndExpected(tt.nft.Flatten(), []byte(tt.expected))
					if err != nil {
						t.Error(err)
					}
				})
			}
		})
	}
}
