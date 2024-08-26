package transactions

type AMMWithdraw struct {
	BaseTx
}

func (*AMMWithdraw) TxType() TxType {
	return AMMWithdrawTx
}
