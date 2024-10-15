package transaction

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestAMMBidTransaction(t *testing.T) {
	s := AMMBid{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: AMMBidTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Asset:  ledger.Asset{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"},
		Asset2: ledger.Asset{Currency: "XRP"},
		BidMin: types.XRPCurrencyAmount(100),
		BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
		AuthAccounts: []ledger.AuthAccounts{
			{
				AuthAccount: ledger.AuthAccount{
					Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
				},
			},
			{
				AuthAccount: ledger.AuthAccount{
					Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE",
				},
			},
		},
	}
	j := `{
	"Account": "abcdef",
	"TransactionType": "AMMBid",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"Asset": {
		"currency": "USD",
		"issuer": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"
	},
	"Asset2": {
		"currency": "XRP"
	},
	"BidMin": "100",
	"BidMax": {
		"issuer": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
		"currency": "USD",
		"value": "200"
	},
	"AuthAccounts": [
		{
			"AuthAccount": {
				"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ"
			}
		},
		{
			"AuthAccount": {
				"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE"
			}
		}
	]
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

	tx, err := UnmarshalTx(json.RawMessage(j))
	if err != nil {
		t.Errorf("UnmarshalTx error: %s", err.Error())
	}
	if !reflect.DeepEqual(tx, &s) {
		t.Error("UnmarshalTx result differs from expected")
	}
}

func TestAMMBidFlatten(t *testing.T) {
	s := AMMBid{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: AMMCreateTx,
			Fee:             types.XRPCurrencyAmount(10),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Asset:  ledger.Asset{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"},
		Asset2: ledger.Asset{Currency: "XRP"},
		BidMin: types.XRPCurrencyAmount(100),
		BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
		AuthAccounts: []ledger.AuthAccounts{
			{
				AuthAccount: ledger.AuthAccount{
					Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
				},
			},
			{
				AuthAccount: ledger.AuthAccount{
					Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
				}},
		},
	}

	flattened := s.Flatten()

	expected := `{
	"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	"TransactionType": "AMMBid",
	"Fee":             "10",
	"Sequence":        1234,
	"SigningPubKey":   "ghijk",
	"TxnSignature":    "A1B2C3D4E5F6",
	"Asset": {
		"currency": "USD",
		"issuer": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"
	},
	"Asset2": {
		"currency": "XRP"
	},
	"BidMin": "100",
	"BidMax": {
		"currency": "USD",
		"issuer": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
		"value": "200"
	},
	"AuthAccounts": [
		{
			"AuthAccount": {
				"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ"
			}
		},
		{
			"AuthAccount": {
				"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ"
			}
		}
	]
}`

	err := testutil.CompareFlattenAndExpected(flattened, []byte(expected))
	if err != nil {
		t.Error(err)
	}
}
