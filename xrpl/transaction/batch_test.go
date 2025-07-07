package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestBatch_TxType(t *testing.T) {
	tx := &Batch{}
	assert.Equal(t, BatchTx, tx.TxType())
}

func TestBatchFlatten(t *testing.T) {
	tests := []struct {
		name     string
		input    Batch
		expected string
	}{
		{
			name: "pass - batch transaction with payment",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(12),
					Flags:           tfAllOrNothing,
				},
				RawTransactions: []RawTransactionWrapper{
					{
						RawTransaction: FlatTransaction{
							"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee":             "0",
							"Flags":           uint32(tfInnerBatchTxn),
							"SigningPubKey":   "",
						},
					},
				},
			},
			expected: `{
				"Account": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
				"TransactionType": "Batch",
				"Fee": "12",
				"Flags": 65536,
				"RawTransactions": [
					{
						"RawTransaction": {
							"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee": "0",
							"Flags": 1073741824,
							"SigningPubKey": ""
						}
					}
				]
			}`,
		},
		{
			name: "pass - batch with offer create and payment transactions",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rUserBSM7T3b6nHX3Jjua62wgX9unH8s9b",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(40),
					Flags:           tfAllOrNothing,
					Sequence:        3,
					SigningPubKey:   "022D40673B44C82DEE1DDB8B9BB53DCCE4F97B27404DB850F068DD91D685E337EA",
					TxnSignature:    "3045022100EC5D367FAE2B461679AD446FBBE7BA260506579AF4ED5EFC3EC25F4DD1885B38022018C2327DB281743B12553C7A6DC0E45B07D3FC6983F261D7BCB474D89A0EC5B8",
				},
				RawTransactions: []RawTransactionWrapper{
					{
						RawTransaction: FlatTransaction{
							"TransactionType": "OfferCreate",
							"Flags":           uint32(tfInnerBatchTxn),
							"Account":         "rUserBSM7T3b6nHX3Jjua62wgX9unH8s9b",
							"TakerGets":       "6000000",
							"TakerPays": map[string]interface{}{
								"currency": "GKO",
								"issuer":   "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
								"value":    "2",
							},
							"Sequence":      uint32(4),
							"Fee":           "0",
							"SigningPubKey": "",
						},
					},
					{
						RawTransaction: FlatTransaction{
							"TransactionType": "Payment",
							"Flags":           uint32(tfInnerBatchTxn),
							"Account":         "rUserBSM7T3b6nHX3Jjua62wgX9unH8s9b",
							"Destination":     "rDEXfrontEnd23E44wKL3S6dj9FaXv",
							"Amount":          "1000",
							"Sequence":        uint32(5),
							"Fee":             "0",
							"SigningPubKey":   "",
						},
					},
				},
			},
			expected: `{
				"Account": "rUserBSM7T3b6nHX3Jjua62wgX9unH8s9b",
				"TransactionType": "Batch",
				"Fee": "40",
				"Flags": 65536,
				"Sequence": 3,
				"SigningPubKey": "022D40673B44C82DEE1DDB8B9BB53DCCE4F97B27404DB850F068DD91D685E337EA",
				"TxnSignature": "3045022100EC5D367FAE2B461679AD446FBBE7BA260506579AF4ED5EFC3EC25F4DD1885B38022018C2327DB281743B12553C7A6DC0E45B07D3FC6983F261D7BCB474D89A0EC5B8",
				"RawTransactions": [
					{
						"RawTransaction": {
							"TransactionType": "OfferCreate",
							"Flags": 1073741824,
							"Account": "rUserBSM7T3b6nHX3Jjua62wgX9unH8s9b",
							"TakerGets": "6000000",
							"TakerPays": {
								"currency": "GKO",
								"issuer": "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
								"value": "2"
							},
							"Sequence": 4,
							"Fee": "0",
							"SigningPubKey": ""
						}
					},
					{
						"RawTransaction": {
							"TransactionType": "Payment",
							"Flags": 1073741824,
							"Account": "rUserBSM7T3b6nHX3Jjua62wgX9unH8s9b",
							"Destination": "rDEXfrontEnd23E44wKL3S6dj9FaXv",
							"Amount": "1000",
							"Sequence": 5,
							"Fee": "0",
							"SigningPubKey": ""
						}
					}
				]
			}`,
		},
		{
			name: "pass - batch with batch signers",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(12),
					Flags:           tfAllOrNothing,
				},
				RawTransactions: []RawTransactionWrapper{
					{
						RawTransaction: FlatTransaction{
							"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee":             "0",
							"Flags":           uint32(tfInnerBatchTxn),
							"SigningPubKey":   "",
						},
					},
				},
				BatchSigners: []BatchSigner{
					{
						BatchSigner: BatchSignerData{
							Account:       "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
							SigningPubKey: "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA06DC66A",
							TxnSignature:  "C4E2834B9C0E7519DC47E4C48F19B4B2C5C92FB4F8C5C8F8C8C8C8C8C8C8C8",
						},
					},
				},
			},
			expected: `{
				"Account": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
				"TransactionType": "Batch",
				"Fee": "12",
				"Flags": 65536,
				"RawTransactions": [
					{
						"RawTransaction": {
							"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee": "0",
							"Flags": 1073741824,
							"SigningPubKey": ""
						}
					}
				],
				"BatchSigners": [
					{
						"BatchSigner": {
							"Account": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
							"SigningPubKey": "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA06DC66A",
							"TxnSignature": "C4E2834B9C0E7519DC47E4C48F19B4B2C5C92FB4F8C5C8F8C8C8C8C8C8C8C8"
						}
					}
				]
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Flatten()
			err := testutil.CompareFlattenAndExpected(result, []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestBatch_Validate(t *testing.T) {
	tests := []struct {
		name     string
		input    Batch
		expected bool
	}{
		{
			name: "pass - valid batch transaction with payments",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(12),
					Flags:           tfAllOrNothing,
				},
				RawTransactions: []RawTransactionWrapper{
					{
						RawTransaction: FlatTransaction{
							"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee":             "0",
							"Flags":           uint32(tfInnerBatchTxn),
							"SigningPubKey":   "",
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "pass - valid batch with multiple payment transactions",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(12),
					Flags:           tfIndependent,
				},
				RawTransactions: []RawTransactionWrapper{
					{
						RawTransaction: FlatTransaction{
							"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee":             "0",
							"Flags":           uint32(tfInnerBatchTxn),
							"SigningPubKey":   "",
						},
					},
					{
						RawTransaction: FlatTransaction{
							"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee":             "0",
							"Flags":           uint32(tfInnerBatchTxn),
							"SigningPubKey":   "",
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "fail - empty raw transactions",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				RawTransactions: []RawTransactionWrapper{},
			},
			expected: false,
		},
		{
			name: "fail - inner transaction missing tfInnerBatchTxn flag",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				RawTransactions: []RawTransactionWrapper{
					{
						RawTransaction: FlatTransaction{
							"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee":             "0",
							"Flags":           uint32(0), // Missing tfInnerBatchTxn
							"SigningPubKey":   "",
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "fail - inner transaction with nested batch",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				RawTransactions: []RawTransactionWrapper{
					{
						RawTransaction: FlatTransaction{
							"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Batch", // Nested batch not allowed
							"Fee":             "0",
							"Flags":           uint32(tfInnerBatchTxn),
							"SigningPubKey":   "",
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "fail - inner transaction with non-zero fee",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				RawTransactions: []RawTransactionWrapper{
					{
						RawTransaction: FlatTransaction{
							"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee":             "12", // Non-zero fee not allowed
							"Flags":           uint32(tfInnerBatchTxn),
							"SigningPubKey":   "",
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "fail - inner transaction with non-empty SigningPubKey",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				RawTransactions: []RawTransactionWrapper{
					{
						RawTransaction: FlatTransaction{
							"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee":             "0",
							"Flags":           uint32(tfInnerBatchTxn),
							"SigningPubKey":   "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA06DC66A", // Non-empty not allowed
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "fail - batch signer with empty account",
			input: Batch{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: BatchTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				RawTransactions: []RawTransactionWrapper{
					{
						RawTransaction: FlatTransaction{
							"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"TransactionType": "Payment",
							"Fee":             "0",
							"Flags":           uint32(tfInnerBatchTxn),
							"SigningPubKey":   "",
						},
					},
				},
				BatchSigners: []BatchSigner{
					{
						BatchSigner: BatchSignerData{
							Account: "", // Empty account not allowed
						},
					},
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.input.Validate()
			if valid != tt.expected {
				t.Errorf("expected %v, got %v, error: %v", tt.expected, valid, err)
			}
		})
	}
}

func TestBatch_Flags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*Batch)
		expected uint32
	}{
		{
			name: "pass - SetAllOrNothingFlag",
			setter: func(b *Batch) {
				b.SetAllOrNothingFlag()
			},
			expected: tfAllOrNothing,
		},
		{
			name: "pass - SetOnlyOneFlag",
			setter: func(b *Batch) {
				b.SetOnlyOneFlag()
			},
			expected: tfOnlyOne,
		},
		{
			name: "pass - SetUntilFailureFlag",
			setter: func(b *Batch) {
				b.SetUntilFailureFlag()
			},
			expected: tfUntilFailure,
		},
		{
			name: "pass - SetIndependentFlag",
			setter: func(b *Batch) {
				b.SetIndependentFlag()
			},
			expected: tfIndependent,
		},
		{
			name: "pass - SetAllOrNothingFlag and SetOnlyOneFlag",
			setter: func(b *Batch) {
				b.SetAllOrNothingFlag()
				b.SetOnlyOneFlag()
			},
			expected: tfAllOrNothing | tfOnlyOne,
		},
		{
			name: "pass - all flags",
			setter: func(b *Batch) {
				b.SetAllOrNothingFlag()
				b.SetOnlyOneFlag()
				b.SetUntilFailureFlag()
				b.SetIndependentFlag()
			},
			expected: tfAllOrNothing | tfOnlyOne | tfUntilFailure | tfIndependent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Batch{}
			tt.setter(b)
			if b.Flags != tt.expected {
				t.Errorf("Expected Batch Flags to be %d, got %d", tt.expected, b.Flags)
			}
		})
	}
}

func TestBatchSigner_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		input    BatchSigner
		expected string
	}{
		{
			name: "pass - basic batch signer",
			input: BatchSigner{
				BatchSigner: BatchSignerData{
					Account:       "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
					SigningPubKey: "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA06DC66A",
					TxnSignature:  "C4E2834B9C0E7519DC47E4C48F19B4B2C5C92FB4F8C5C8F8C8C8C8C8C8C8C8",
				},
			},
			expected: `{
				"BatchSigner": {
					"Account": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
					"SigningPubKey": "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA06DC66A",
					"TxnSignature": "C4E2834B9C0E7519DC47E4C48F19B4B2C5C92FB4F8C5C8F8C8C8C8C8C8C8C8"
				}
			}`,
		},
		{
			name: "pass - batch signer with inner signers",
			input: BatchSigner{
				BatchSigner: BatchSignerData{
					Account: "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
					Signers: []Signer{
						{
							SignerData: SignerData{
								Account:       "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
								SigningPubKey: "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA06DC66A",
								TxnSignature:  "C4E2834B9C0E7519DC47E4C48F19B4B2C5C92FB4F8C5C8F8C8C8C8C8C8C8C8",
							},
						},
					},
				},
			},
			expected: `{
				"BatchSigner": {
					"Account": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
					"Signers": [
						{
							"Signer": {
								"Account": "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
								"SigningPubKey": "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA06DC66A",
								"TxnSignature": "C4E2834B9C0E7519DC47E4C48F19B4B2C5C92FB4F8C5C8F8C8C8C8C8C8C8C8"
							}
						}
					]
				}
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Flatten()
			err := testutil.CompareFlattenAndExpected(result, []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}
