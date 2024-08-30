package transactions

type AMMVote struct {
	BaseTx
}

func (*AMMVote) TxType() TxType {
	return AMMVoteTx
}

// TODO: Implement flatten
func (s *AMMVote) Flatten() map[string]interface{} {
	return nil
}
