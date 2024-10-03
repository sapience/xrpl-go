package transaction

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestPaymentChannelCreate(t *testing.T) {
	s := PaymentChannelCreate{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: PaymentChannelCreateTx,
		},
		Amount:      types.XRPCurrencyAmount(1000),
		Destination: "def",
		SettleDelay: 10,
		PublicKey:   "abcd",
	}

	j := `{
	"Account": "abc",
	"TransactionType": "PaymentChannelCreate",
	"Amount": "1000",
	"Destination": "def",
	"SettleDelay": 10,
	"PublicKey": "abcd"
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
