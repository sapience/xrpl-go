package hash

import (
	"testing"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/stretchr/testify/require"
)

func TestHashTxBlob(t *testing.T) {
	testCases := []struct {
		name    string
		txBlob  string
		want    string
		wantErr bool
	}{
		{
			name:    "valid tx blob",
			txBlob:  "120000220000000024001B733261400000000000000F68400000000000000C7321ED90ADC33C2BBD9B4A0D94223DBE30D34227B82F587C5909A857B3AB7DE8D6E2EF74402754D4EE7EBDA0A073488904E8A55CECAEDA13EA2829AF5C0EB2CC201C4B4E2AB72D20D308EE12C5D1C112BCFCAFEBDA6C8198D92C0C57F15D8A25B5BFBF200E811474E4DD74B588FA412F0993B8E7E07C2FA92109B48314858233827B488ECB8D0EB940E7AC85CE41E343CF",
			want:    "BE76FC0ABE8BE83F91219D2371FF5199F0271ACF0E12794D2EA5DE77AC49E877",
			wantErr: false,
		},
		{
			name:    "invalid tx blob",
			txBlob:  "120000220000000024001B733261400000000000001268400000000000000C811474E4DD74B588FA412F0993B8E7E07C2FA92109B48314D708DAB02885BA68A48EBCC4EE3551CF1AF7B267",
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SignTxBlob(tc.txBlob)
			if (err != nil) != tc.wantErr {
				t.Errorf("HashSignedTx() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("HashSignedTx() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSignTxBlob(t *testing.T) {
	tests := []struct {
		name     string
		tx       map[string]interface{}
		expected bool
	}{
		{
			name: "pass - has TxnSignature",
			tx: map[string]interface{}{
				"TransactionType": "Payment",
				"Account":         "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
				"TxnSignature":    "30440220702ABC11419AD4940969CC32EB4D1BFDBFCA651F064F30D6E1646D74FBFC493902204E5B451B447B0F69904127F04FE71634BD825A8970B9467871DA89EEC4B021F8",
				"Flags":           uint32(0),
			},
			expected: true,
		},
		{
			name: "pass - has SigningPubKey",
			tx: map[string]interface{}{
				"TransactionType": "Payment",
				"Account":         "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
				"SigningPubKey":   "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA06DC66A",
				"Flags":           uint32(0),
			},
			expected: true,
		},
		{
			name: "pass - has Signers",
			tx: map[string]interface{}{
				"TransactionType": "Payment",
				"Account":         "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
				"Signers": []interface{}{
					map[string]interface{}{
						"Signer": map[string]interface{}{
							"Account":       "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
							"SigningPubKey": "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA06DC66A",
							"TxnSignature":  "30440220702ABC11419AD4940969CC32EB4D1BFDBFCA651F064F30D6E1646D74FBFC493902204E5B451B447B0F69904127F04FE71634BD825A8970B9467871DA89EEC4B021F8",
						},
					},
				},
				"Flags": uint32(0),
			},
			expected: true,
		},
		{
			name: "fail - no signature fields",
			tx: map[string]interface{}{
				"TransactionType": "Payment",
				"Account":         "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
				"Destination":     "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
				"Amount":          "1000000",
				"Flags":           uint32(0),
			},
			expected: false,
		},
		{
			name: "pass - inner-batch skips signature",
			tx: map[string]interface{}{
				"TransactionType": "Payment",
				"Account":         "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
				"Destination":     "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
				"Amount":          "1000000",
				"Flags":           uint32(transaction.TfInnerBatchTxn),
			},
			expected: true,
		},
		{
			name: "pass - inner-batch + other flags skips signature",
			tx: map[string]interface{}{
				"TransactionType": "Payment",
				"Account":         "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
				"Destination":     "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
				"Amount":          "1000000",
				"Flags":           uint32(transaction.TfInnerBatchTxn | 0x00010000),
			},
			expected: true,
		},
		{
			name: "fail - other flags without inner-batch",
			tx: map[string]interface{}{
				"TransactionType": "Payment",
				"Account":         "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
				"Flags":           uint32(0x00010000),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode to blob
			blob, err := binarycodec.Encode(tt.tx)
			require.NoError(t, err, "Encode failed")

			hash, err := SignTxBlob(blob)

			if tt.expected {
				require.NoError(t, err)
				require.NotEmpty(t, hash)
			} else {
				require.Error(t, err)
				require.Empty(t, hash)
			}
		})
	}
}
