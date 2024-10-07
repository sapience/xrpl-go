package transaction

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestNFTokenCancelOfferTx(t *testing.T) {
	s := NFTokenCancelOffer{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: NFTokenCancelOfferTx,
		},
		NFTokenOffers: []types.Hash256{
			"ABC",
			"DEF",
		},
	}
	j := `{
	"Account": "abcdef",
	"TransactionType": "NFTokenCancelOffer",
	"NFTokenOffers": [
		"ABC",
		"DEF"
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
