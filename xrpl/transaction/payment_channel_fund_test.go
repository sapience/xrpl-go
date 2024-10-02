package transaction

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestPaymentChannelFund(t *testing.T) {
	s := PaymentChannelFund{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: PaymentChannelFundTx,
		},
		Channel: "ABACAD",
		Amount:  types.XRPCurrencyAmount(2000),
	}

	j := `{
	"Account": "abc",
	"TransactionType": "PaymentChannelFund",
	"Channel": "ABACAD",
	"Amount": "2000"
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
