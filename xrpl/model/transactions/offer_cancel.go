package transactions

type OfferCancel struct {
	BaseTx
	OfferSequence uint
}

func (*OfferCancel) TxType() TxType {
	return OfferCancelTx
}

// TODO: Implement flatten
func (s *OfferCancel) Flatten() FlatTransaction {
	return nil
}
