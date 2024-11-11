package ledger

type ClosedRequest struct {
}

func (*ClosedRequest) Method() string {
	return "ledger_closed"
}
