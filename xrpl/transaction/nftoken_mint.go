package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type NFTokenMint struct {
	BaseTx
	NFTokenTaxon uint32
	Issuer       types.Address    `json:",omitempty"`
	TransferFee  uint16           `json:",omitempty"`
	URI          types.NFTokenURI `json:",omitempty"`
}

func (*NFTokenMint) TxType() TxType {
	return NFTokenMintTx
}

// TODO: Implement flatten
func (s *NFTokenMint) Flatten() FlatTransaction {
	return nil
}
