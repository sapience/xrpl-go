package ledger

type CurrentRequest struct {
}

func (*CurrentRequest) Method() string {
	return "ledger_current"
}
