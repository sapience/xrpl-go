package transactions

type AMMDeposit struct {
	BaseTx
}

func (*AMMDeposit) TxType() TxType {
	return AMMDepositTx
}

// TODO: Implement flatten
func (s *AMMDeposit) Flatten() FlatTransaction {
	return nil
}
