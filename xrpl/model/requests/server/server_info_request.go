package server

type ServerInfoRequest struct {
}

func (*ServerInfoRequest) Method() string {
	return "server_info"
}

// TODO: Implement
func (*ServerInfoRequest) Validate() error {
	return nil
}
