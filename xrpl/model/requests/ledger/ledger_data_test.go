package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/test"
)

func TestLedgerDataRequest(t *testing.T) {
	s := LedgerDataRequest{
		LedgerIndex: common.CLOSED,
		Binary:      true,
		Limit:       5,
	}
	j := `{
	"ledger_index": "closed",
	"binary": true,
	"limit": 5
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
