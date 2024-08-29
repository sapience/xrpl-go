package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type NFTokenCancelOffer struct {
	BaseTx
	NFTokenOffers []types.Hash256
}

func (*NFTokenCancelOffer) TxType() TxType {
	return NFTokenCancelOfferTx
}

// TODO: Implement flatten
func (s *NFTokenCancelOffer) Flatten() map[string]interface{} {
	return nil
}
