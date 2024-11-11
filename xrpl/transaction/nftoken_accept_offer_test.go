package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestNFTokenAcceptOffer_TxType(t *testing.T) {
	tx := &NFTokenAcceptOffer{}
	assert.Equal(t, NFTokenAcceptOfferTx, tx.TxType())
}

func TestNFTokenAcceptOffer_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		tx       *NFTokenAcceptOffer
		expected string
	}{
		{
			name: "BaseTx only NFTokenAcceptOffer",
			tx: &NFTokenAcceptOffer{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenAcceptOfferTx,
				},
			},
			expected: `{
				"Account": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
				"TransactionType": "NFTokenAcceptOffer"
			}`,
		},
		{
			name: "NFTokenAcceptOffer with Sell Offer",
			tx: &NFTokenAcceptOffer{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenAcceptOfferTx,
				},
				NFTokenSellOffer: "68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77",
			},
			expected: `{
				"Account": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
				"TransactionType": "NFTokenAcceptOffer",
				"NFTokenSellOffer": "68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77"
			}`,
		},
		{
			name: "NFTokenAcceptOffer with Buy Offer",
			tx: &NFTokenAcceptOffer{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenAcceptOfferTx,
				},
				NFTokenBuyOffer: "68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77",
			},
			expected: `{
				"Account": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
				"TransactionType": "NFTokenAcceptOffer",
				"NFTokenBuyOffer": "68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77"
			}`,
		},
		{
			name: "NFTokenAcceptOffer with Broker Fee",
			tx: &NFTokenAcceptOffer{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenAcceptOfferTx,
				},
				NFTokenBrokerFee: types.XRPCurrencyAmount(1000),
			},
			expected: `{
				"Account": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
				"TransactionType": "NFTokenAcceptOffer",
				"NFTokenBrokerFee": "1000"
			}`,
		},
		{
			name: "NFTokenAcceptOffer with all fields",
			tx: &NFTokenAcceptOffer{
				BaseTx: BaseTx{
					Account:         "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
					TransactionType: NFTokenAcceptOfferTx,
				},
				NFTokenSellOffer: "68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77",
				NFTokenBuyOffer:  "68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77",
				NFTokenBrokerFee: types.XRPCurrencyAmount(1000),
			},
			expected: `{
				"Account": "rNCFjv8Ek5oDrNiMJ3pw6eLLFtMjZLJnf2",
				"TransactionType": "NFTokenAcceptOffer",
				"NFTokenSellOffer": "68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77",
				"NFTokenBuyOffer": "68CD1F6F906494EA08C9CB5CAFA64DFA90D4E834B7151899B73231DE5A0C3B77",
				"NFTokenBrokerFee": "1000"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.tx.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}
