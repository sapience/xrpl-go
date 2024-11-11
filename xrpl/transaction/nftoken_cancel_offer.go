package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type NFTokenCancelOffer struct {
	BaseTx
	NFTokenOffers []types.Hash256
}

func (*NFTokenCancelOffer) TxType() TxType {
	return NFTokenCancelOfferTx
}

// TODO: Implement flatten
func (s *NFTokenCancelOffer) Flatten() FlatTransaction {
	return nil
}
