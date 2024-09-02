package transactions

// TODO: Implement AMMBid
type AMMBid struct {
	BaseTx
}

func (*AMMBid) TxType() TxType {
	return AMMBidTx
}

// TODO: Implement flatten
func (s *AMMBid) Flatten() FlatTransaction {
	return nil
}
