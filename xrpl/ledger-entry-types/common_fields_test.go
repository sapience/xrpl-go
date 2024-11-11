package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

func TestLedgerEntryCommonFields_SerializeAndDeserialize(t *testing.T) {
	ledgerEntryCommonFields := EntryCommonFields{
		Index:           "1",
		LedgerEntryType: "AccountRoot",
		Flags:           123,
	}

	json := `{
	"index": "1",
	"LedgerEntryType": "AccountRoot",
	"Flags": 123
}`

	if err := testutil.SerializeAndDeserialize(t, ledgerEntryCommonFields, json); err != nil {
		t.Error(err)
	}
}
