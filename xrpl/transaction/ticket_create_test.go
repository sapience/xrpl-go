package transaction

import (
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

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
