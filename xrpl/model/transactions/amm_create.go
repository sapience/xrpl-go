package transactions

type AMMCreate struct {
	BaseTx
}

func (*AMMCreate) TxType() TxType {
	return AMMCreateTx
}
