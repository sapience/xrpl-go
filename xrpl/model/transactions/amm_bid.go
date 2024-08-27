package transactions

// TODO: Implement AMMBid
type AMMBid struct {
	BaseTx
}

func (*AMMBid) TxType() TxType {
	return AMMBidTx
}

// TODO: Implement flatten
func (s *AMMBid) Flatten() map[string]interface{} {
	return nil
}
