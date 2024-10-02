package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

func TestLedgerEntryRequest(t *testing.T) {
	s := LedgerEntryRequest{
		LedgerIndex: common.VALIDATED,
		Directory: &DirectoryEntryReq{
			Owner: "abc",
		},
	}
	j := `{
	"ledger_index": "validated",
	"directory": {
		"owner": "abc"
	}
}`
	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
