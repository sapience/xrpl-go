package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type NFTokenAcceptOffer struct {
	BaseTx
	NFTokenSellOffer types.Hash256        `json:",omitempty"`
	NFTokenBuyOffer  types.Hash256        `json:",omitempty"`
	NFTokenBrokerFee types.CurrencyAmount `json:",omitempty"`
}

func (*NFTokenAcceptOffer) TxType() TxType {
	return NFTokenAcceptOfferTx
}

// TODO: Implement flatten
func (s *NFTokenAcceptOffer) Flatten() FlatTransaction {
	return nil
}
