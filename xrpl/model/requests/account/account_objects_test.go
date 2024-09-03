package account

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/test"
)

func TestAccountObjectsRequest(t *testing.T) {
	s := AccountObjectsRequest{
		Account:     "rsuHaTvJh1bDmDoxX9QcKP7HEBSBt4XsHx",
		Type:        SignerListObject,
		LedgerIndex: common.LedgerIndex(123),
	}

	j := `{
	"account": "rsuHaTvJh1bDmDoxX9QcKP7HEBSBt4XsHx",
	"type": "signer_list",
	"ledger_index": 123
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

