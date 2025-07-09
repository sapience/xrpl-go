package transaction

import (
	"encoding/json"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNFTokenMintMetadata_JSONMarshal(t *testing.T) {
	tests := []struct {
		name     string
		metadata *NFTokenMintMetadata
		expected string
	}{
		{
			name: "pass - with both NFTokenID and OfferID",
			metadata: &NFTokenMintMetadata{
				NFTokenID: func() *types.NFTokenID {
					id := types.NFTokenID("000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65")
					return &id
				}(),
				OfferID: func() *types.Hash256 {
					hash := types.Hash256("68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77")
					return &hash
				}(),
			},
			expected: `{"nftoken_id":"000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65","offer_id":"68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77"}`,
		},
		{
			name:     "pass - empty metadata",
			metadata: &NFTokenMintMetadata{},
			expected: `{}`,
		},
		{
			name: "pass - with full metadata including TxObjMeta",
			metadata: &NFTokenMintMetadata{
				TxObjMeta: TxObjMeta{
					TransactionIndex:  123,
					TransactionResult: "tesSUCCESS",
				},
				NFTokenID: func() *types.NFTokenID {
					id := types.NFTokenID("000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65")
					return &id
				}(),
				OfferID: func() *types.Hash256 {
					hash := types.Hash256("68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77")
					return &hash
				}(),
			},
			expected: `{"TransactionIndex":123,"TransactionResult":"tesSUCCESS","nftoken_id":"000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65","offer_id":"68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77"}`,
		},
		{
			name: "pass - nil pointers should omit fields",
			metadata: &NFTokenMintMetadata{
				NFTokenID: nil,
				OfferID:   nil,
			},
			expected: `{}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, err := json.Marshal(tt.metadata)
			require.NoError(t, err)
			assert.JSONEq(t, tt.expected, string(jsonBytes))
		})
	}
}

func TestNFTokenMintMetadata_JSONUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		expected *NFTokenMintMetadata
	}{
		{
			name: "pass - with both NFTokenID and OfferID",
			json: `{"nftoken_id":"000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65","offer_id":"68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77"}`,
			expected: &NFTokenMintMetadata{
				NFTokenID: func() *types.NFTokenID {
					id := types.NFTokenID("000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65")
					return &id
				}(),
				OfferID: func() *types.Hash256 {
					hash := types.Hash256("68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77")
					return &hash
				}(),
			},
		},
		{
			name:     "pass - empty metadata",
			json:     `{}`,
			expected: &NFTokenMintMetadata{},
		},
		{
			name: "pass - with full metadata including TxObjMeta fields",
			json: `{"TransactionIndex":123,"TransactionResult":"tesSUCCESS","nftoken_id":"000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65","offer_id":"68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77"}`,
			expected: &NFTokenMintMetadata{
				TxObjMeta: TxObjMeta{
					TransactionIndex:  123,
					TransactionResult: "tesSUCCESS",
				},
				NFTokenID: func() *types.NFTokenID {
					id := types.NFTokenID("000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65")
					return &id
				}(),
				OfferID: func() *types.Hash256 {
					hash := types.Hash256("68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77")
					return &hash
				}(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var metadata NFTokenMintMetadata
			err := json.Unmarshal([]byte(tt.json), &metadata)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, &metadata)
		})
	}
}
