package transaction

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestSetRegularKeyTx(t *testing.T) {
	s := SetRegularKey{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: SetRegularKeyTx,
			Fee:             types.XRPCurrencyAmount(10),
		},
		RegularKey: "def",
	}

	j := `{
	"Account": "abc",
	"TransactionType": "SetRegularKey",
	"Fee": "10",
	"RegularKey": "def"
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
