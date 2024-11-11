package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type NFTokenCreateOffer struct {
	BaseTx
	Owner       types.Address `json:",omitempty"`
	NFTokenID   types.NFTokenID
	Amount      types.CurrencyAmount
	Expiration  uint          `json:",omitempty"`
	Destination types.Address `json:",omitempty"`
}

func (*NFTokenCreateOffer) TxType() TxType {
	return NFTokenCreateOfferTx
}

// TODO: Implement flatten
func (n *NFTokenCreateOffer) Flatten() FlatTransaction {
	return nil
}
