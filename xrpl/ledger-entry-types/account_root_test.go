package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestAccountRoot(t *testing.T) {
	var s Object = &AccountRoot{
		Account:           "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		AccountTxnID:      "0D5FB50FA65C9FE1538FD7E398FFFE9D1908DFA4576D8D7A020040686F93C77D",
		Balance:           types.XRPCurrencyAmount(148446663),
		Domain:            "6D64756F31332E636F6D",
		EmailHash:         "98B4375E1D753E5B91627516F6D70977",
		Flags:             8388608,
		LedgerEntryType:   AccountRootEntry,
		MessageKey:        "0000000000000000000000070000000300",
		OwnerCount:        3,
		PreviousTxnID:     "0D5FB50FA65C9FE1538FD7E398FFFE9D1908DFA4576D8D7A020040686F93C77D",
		PreviousTxnLgrSeq: 14091160,
		Sequence:          336,
		TransferRate:      1004999999,
	}
	j := `{
	"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
	"AccountTxnID": "0D5FB50FA65C9FE1538FD7E398FFFE9D1908DFA4576D8D7A020040686F93C77D",
	"Balance": "148446663",
	"Domain": "6D64756F31332E636F6D",
	"EmailHash": "98B4375E1D753E5B91627516F6D70977",
	"Flags": 8388608,
	"LedgerEntryType": "AccountRoot",
	"MessageKey": "0000000000000000000000070000000300",
	"OwnerCount": 3,
	"PreviousTxnID": "0D5FB50FA65C9FE1538FD7E398FFFE9D1908DFA4576D8D7A020040686F93C77D",
	"PreviousTxnLgrSeq": 14091160,
	"Sequence": 336,
	"TransferRate": 1004999999
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
