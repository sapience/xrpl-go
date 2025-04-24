package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNFTokenModify_TxType(t *testing.T) {
	tx := &NFTokenModify{}
	assert.Equal(t, NFTokenModifyTx, tx.TxType())
}

func TestNFTokenModify_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		nft      *NFTokenModify
		expected FlatTransaction
	}{
		{
			name: "pass - all fields",
			nft: &NFTokenModify{
				BaseTx: BaseTx{
					TransactionType: NFTokenMintTx,
					Account:         "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				},
				Owner:     "rogue5HnPRSszD9CWGSUz8UGHMVwSSKF6",
				NFTokenID: "0008C350C182B4F213B82CCFA4C6F59AD76F0AFCFBDF04D5A048C0A300000007",
				URI:       "697066733A2F2F62616679626569636D6E73347A736F6C686C6976346C746D6E356B697062776373637134616C70736D6C6179696970666B73746B736D3472746B652F5665742E706E67",
			},
			expected: FlatTransaction{
				"TransactionType": "NFTokenModify",
				"Account":         "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				"Owner":           types.Address("rogue5HnPRSszD9CWGSUz8UGHMVwSSKF6"),
				"NFTokenID":       types.NFTokenID("0008C350C182B4F213B82CCFA4C6F59AD76F0AFCFBDF04D5A048C0A300000007"),
				"URI":             types.NFTokenURI("697066733A2F2F62616679626569636D6E73347A736F6C686C6976346C746D6E356B697062776373637134616C70736D6C6179696970666B73746B736D3472746B652F5665742E706E67"),
			},
		},
		{
			name: "pass - minimum required fields",
			nft: &NFTokenModify{
				BaseTx: BaseTx{
					TransactionType: NFTokenMintTx,
					Account:         "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				},
				NFTokenID: "0008C350C182B4F213B82CCFA4C6F59AD76F0AFCFBDF04D5A048C0A300000007",
			},
			expected: FlatTransaction{
				"TransactionType": "NFTokenModify",
				"Account":         "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				"NFTokenID":       types.NFTokenID("0008C350C182B4F213B82CCFA4C6F59AD76F0AFCFBDF04D5A048C0A300000007"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.nft.Flatten())
		})
	}
}
