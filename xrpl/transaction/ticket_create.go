package transaction

// A TicketCreate transaction sets aside one or more sequence numbers as Tickets.
type TicketCreate struct {
	// Base transaction fields
	BaseTx

	//How many Tickets to create. This must be a positive number and cannot cause
	// the account to own more than 250 Tickets after executing this transaction.
	TicketCount uint32
}

func (*TicketCreate) TxType() TxType {
	return TicketCreateTx
}

func (t *TicketCreate) Flatten() FlatTransaction {
	flattened := t.BaseTx.Flatten()

	flattened["TransactionType"] = "TicketCreate"

	if t.TicketCount != 0 {
		flattened["TicketCount"] = t.TicketCount
	}

	return flattened
}
