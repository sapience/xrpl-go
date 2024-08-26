package transactions

type AMMDeposit struct {
	BaseTx
}

func (*AMMDeposit) TxType() TxType {
	return AMMDepositTx
}
