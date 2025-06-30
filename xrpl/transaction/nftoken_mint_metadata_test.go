package transaction

import (
	"encoding/json"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNFTokenMintMetadata_TxMeta(t *testing.T) {
	// Test that NFTokenMintMetadata implements TxMeta interface
	var _ TxMeta = NFTokenMintMetadata{}
	var _ TxMeta = &NFTokenMintMetadata{}
}

func TestNFTokenMintMetadata_EmbeddedFields(t *testing.T) {
	tests := []struct {
		name                      string
		metadata                  *NFTokenMintMetadata
		expectedTransactionIndex  uint64
		expectedTransactionResult string
		expectedNFTokenID         string
		expectedOfferID           string
	}{
		{
			name: "pass - embedded fields are accessible",
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
			expectedTransactionIndex:  123,
			expectedTransactionResult: "tesSUCCESS",
			expectedNFTokenID:         "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65",
			expectedOfferID:           "68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77",
		},
		{
			name: "pass - embedded fields with different values",
			metadata: &NFTokenMintMetadata{
				TxObjMeta: TxObjMeta{
					TransactionIndex:  456,
					TransactionResult: "tecPATH_PARTIAL",
				},
				NFTokenID: func() *types.NFTokenID {
					id := types.NFTokenID("111A013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E11111111")
					return &id
				}(),
				OfferID: func() *types.Hash256 {
					hash := types.Hash256("22221F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B22")
					return &hash
				}(),
			},
			expectedTransactionIndex:  456,
			expectedTransactionResult: "tecPATH_PARTIAL",
			expectedNFTokenID:         "111A013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E11111111",
			expectedOfferID:           "22221F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B22",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedTransactionIndex, tt.metadata.TransactionIndex)
			assert.Equal(t, tt.expectedTransactionResult, tt.metadata.TransactionResult)
			assert.Equal(t, tt.expectedNFTokenID, tt.metadata.NFTokenID.String())
			assert.Equal(t, tt.expectedOfferID, tt.metadata.OfferID.String())
		})
	}
}

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
			name: "pass - with only NFTokenID",
			metadata: &NFTokenMintMetadata{
				NFTokenID: func() *types.NFTokenID {
					id := types.NFTokenID("000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65")
					return &id
				}(),
			},
			expected: `{"nftoken_id":"000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65"}`,
		},
		{
			name: "pass - with only OfferID",
			metadata: &NFTokenMintMetadata{
				OfferID: func() *types.Hash256 {
					hash := types.Hash256("68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77")
					return &hash
				}(),
			},
			expected: `{"offer_id":"68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77"}`,
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
			expected: `{
				"TransactionIndex": 123,
				"TransactionResult": "tesSUCCESS",
				"nftoken_id": "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65",
				"offer_id": "68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77"
			}`,
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
			name: "pass - with only NFTokenID",
			json: `{"nftoken_id":"000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65"}`,
			expected: &NFTokenMintMetadata{
				NFTokenID: func() *types.NFTokenID {
					id := types.NFTokenID("000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65")
					return &id
				}(),
			},
		},
		{
			name: "pass - with only OfferID",
			json: `{"offer_id":"68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77"}`,
			expected: &NFTokenMintMetadata{
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

func TestNFTokenMintMetadata_PointerFields(t *testing.T) {
	tests := []struct {
		name     string
		metadata *NFTokenMintMetadata
		expected string
	}{
		{
			name: "pass - nil pointers should omit fields",
			metadata: &NFTokenMintMetadata{
				NFTokenID: nil,
				OfferID:   nil,
			},
			expected: `{}`,
		},
		{
			name: "pass - valid pointers should include fields",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, err := json.Marshal(tt.metadata)
			require.NoError(t, err)
			assert.JSONEq(t, tt.expected, string(jsonBytes))
		})
	}
}
