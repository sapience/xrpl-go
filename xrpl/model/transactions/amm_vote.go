package transactions

type AMMVote struct {
	BaseTx
}

func (*AMMVote) TxType() TxType {
	return AMMVoteTx
}
