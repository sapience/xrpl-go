package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestAmm(t *testing.T) {
	amm := AMM{
		LedgerEntryCommonFields: LedgerEntryCommonFields{
			Flags: 0,
		},
		Account: "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
		Asset: Asset{
			Currency: "XRP",
		},
		Asset2: Asset{
			Currency: "TST",
			Issuer:   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
		},
		AuctionSlot: AuctionSlot{
			Account: "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
			AuthAccounts: []AuthAccounts{
				{
					AuthAccount: AuthAccount{
						Account: "rMKXGCbJ5d8LbrqthdG46q3f969MVK2Qeg",
					},
				},
				{
					AuthAccount: AuthAccount{
						Account: "rBepJuTLFJt3WmtLXYAxSjtBWAeQxVbncv",
					},
				},
			},
			DiscountedFee: 60,
			Expiration:    721870180,
			Price: types.IssuedCurrencyAmount{
				Currency: "039C99CD9AB0B70B32ECDA51EAAE471625608EA2",
				Issuer:   "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
				Value:    "0.8696263565463045",
			},
		},
		LPTokenBalance: types.IssuedCurrencyAmount{
			Currency: "039C99CD9AB0B70B32ECDA51EAAE471625608EA2",
			Issuer:   "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
			Value:    "71150.53584131501",
		},
		TradingFee: 600,
		VoteSlots: []VoteSlots{
			{
				VoteEntry: VoteEntry{
					Account:    "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TradingFee: 600,
					VoteWeight: 100000,
				},
			},
		},
	}

	json := `{
	"Flags": 0,
	"Account": "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
	"Asset": {
		"currency": "XRP"
	},
	"Asset2": {
		"currency": "TST",
		"issuer": "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
	},
	"AuctionSlot": {
		"Account": "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
		"AuthAccounts": [
			{
				"AuthAccount": {
					"Account": "rMKXGCbJ5d8LbrqthdG46q3f969MVK2Qeg"
				}
			},
			{
				"AuthAccount": {
					"Account": "rBepJuTLFJt3WmtLXYAxSjtBWAeQxVbncv"
				}
			}
		],
		"DiscountedFee": 60,
		"Price": {
			"issuer": "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
			"currency": "039C99CD9AB0B70B32ECDA51EAAE471625608EA2",
			"value": "0.8696263565463045"
		},
		"Expiration": 721870180
	},
	"LPTokenBalance": {
		"issuer": "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
		"currency": "039C99CD9AB0B70B32ECDA51EAAE471625608EA2",
		"value": "71150.53584131501"
	},
	"TradingFee": 600,
	"VoteSlots": [
		{
			"VoteEntry": {
				"Account": "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
				"TradingFee": 600,
				"VoteWeight": 100000
			}
		}
	]
}`

	if err := testutil.SerializeAndDeserialize(t, amm, json); err != nil {
		t.Error(err)
	}
}

func TestAssetFlatten(t *testing.T) {
	asset := Asset{
		Currency: "USD",
		Issuer:   "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
	}

	json := `{
	"currency": "USD",
	"issuer": "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm"
}`

	if err := testutil.CompareFlattenAndExpected(asset.Flatten(), []byte(json)); err != nil {
		t.Error(err)
	}

	// 2nd test with issuer empty
	asset2 := Asset{
		Currency: "XRP",
	}

	json2 := `{
	"currency": "XRP"
}`

	if err := testutil.CompareFlattenAndExpected(asset2.Flatten(), []byte(json2)); err != nil {
		t.Error(err)
	}
}
func TestUnmarshalAsset(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		want    Asset
		wantErr bool
	}{
		{
			name: "Valid Asset with Issuer",
			json: `{"currency": "USD", "issuer": "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm"}`,
			want: Asset{
				Currency: "USD",
				Issuer:   "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
			},
			wantErr: false,
		},
		{
			name: "Valid Asset without Issuer and XRP Currency",
			json: `{"currency": "XRP"}`,
			want: Asset{
				Currency: "XRP",
			},
			wantErr: false,
		},
		{
			name:    "Empty JSON",
			json:    `{}`,
			want:    Asset{},
			wantErr: false,
		},
		{
			name:    "Invalid JSON",
			json:    `{"currency": "USD", "issuer": 12345}`,
			want:    Asset{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalAsset([]byte(tt.json))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalAsset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("UnmarshalAsset() = %v, want %v", got, tt.want)
			}
		})
	}
}
