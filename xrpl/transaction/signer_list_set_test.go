package transaction

import (
	"testing"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestSignerListSet_TxType(t *testing.T) {
	entry := &SignerListSet{}
	assert.Equal(t, SignerListSetTx, entry.TxType())
}

func TestSignerListSet_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		entry    *SignerListSet
		expected string
	}{
		{
			name: "With SignerEntries",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Fee:     types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 3,
				SignerEntries: []ledger.SignerEntryWrapper{
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
							SignerWeight: 2,
						},
					},
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
							SignerWeight: 1,
						},
					},
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "raKEEVSGnKSD9Zyvxu4z6Pqpm4ABH8FS6n",
							SignerWeight: 1,
						},
					},
				},
			},
			expected: `{
				"TransactionType": "SignerListSet",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Fee": "12",
				"SignerQuorum": 3,
				"SignerEntries": [
					{
						"SignerEntry": {
							"Account": "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
							"SignerWeight": 2
						}
					},
					{
						"SignerEntry": {
							"Account": "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
							"SignerWeight": 1
						}
					},
					{
						"SignerEntry": {
							"Account": "raKEEVSGnKSD9Zyvxu4z6Pqpm4ABH8FS6n",
							"SignerWeight": 1
						}
					}
				]
			}`,
		},
		{
			name: "Without SignerEntries",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Fee:     types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 0,
			},
			expected: `{
				"TransactionType": "SignerListSet",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Fee": "12",
				"SignerQuorum": 0
			}`,
		},
		{
			name: "Without SignerEntries and SignerQuorum",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Fee:     types.XRPCurrencyAmount(12),
				},
			},
			expected: `{
				"TransactionType": "SignerListSet",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Fee": "12",
				"SignerQuorum": 0
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.entry.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}
