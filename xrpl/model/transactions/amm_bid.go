package transactions

// TODO: Implement AMMBid
type AMMBid struct {
	BaseTx
}

func (*AMMBid) TxType() TxType {
	return AMMBidTx
}
