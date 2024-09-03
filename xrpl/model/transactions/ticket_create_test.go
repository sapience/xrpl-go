package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/test"
)

func TestTicketCreateTx(t *testing.T) {
	s := TicketCreate{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: TicketCreateTx,
			Fee:             types.XRPCurrencyAmount(10),
			Sequence:        50,
		},
		TicketCount: 5,
	}

	j := `{
	"Account": "abc",
	"TransactionType": "TicketCreate",
	"Fee": "10",
	"Sequence": 50,
	"TicketCount": 5
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
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

func TestTicketCreateFlatten(t *testing.T) {
	s := TicketCreate{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: TicketCreateTx,
			Fee:             types.XRPCurrencyAmount(10),
			Sequence:        50,
		},
		TicketCount: 5,
	}

	flattened := s.Flatten()

	expected := FlatTransaction{
		"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "TicketCreate",
		"Fee":             "10",
		"Sequence":        uint(50),
		"TicketCount":     uint32(5),
	}

	if !reflect.DeepEqual(flattened, expected) {
		t.Errorf("Flatten result differs from expected: %v, %v", flattened, expected)
	}
}
