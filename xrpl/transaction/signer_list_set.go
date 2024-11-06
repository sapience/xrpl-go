package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
)

type SignerListSet struct {
	BaseTx
	SignerQuorum  uint32
	SignerEntries []ledger.SignerEntryWrapper
}

func (*SignerListSet) TxType() TxType {
	return SignerListSetTx
}

// TODO: Implement flatten
func (s *SignerListSet) Flatten() FlatTransaction {
	return nil
}
