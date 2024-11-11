package server

type StateRequest struct {
}

func (*StateRequest) Method() string {
	return "server_state"
}
