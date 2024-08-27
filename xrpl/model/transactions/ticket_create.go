package transactions

type TicketCreate struct {
	BaseTx
	TicketCount uint
}

func (*TicketCreate) TxType() TxType {
	return TicketCreateTx
}

func (t *TicketCreate) Flatten() map[string]interface{} {
	flattened := t.BaseTx.Flatten()

	flattened["TransactionType"] = "TicketCreate"

	if t.TicketCount != 0 {
		flattened["TicketCount"] = t.TicketCount
	}

	return flattened
}
