package peer

type ConnectRequest struct {
	IP   string `json:"ip"`
	Port int    `json:"port,omitempty"`
}

func (*ConnectRequest) Method() string {
	return "connect"
}

type ConnectResponse struct {
	Message string `json:"message"`
}
