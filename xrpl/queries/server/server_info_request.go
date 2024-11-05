package server

type InfoRequest struct {
}

func (*InfoRequest) Method() string {
	return "server_info"
}

// TODO: Implement
func (*InfoRequest) Validate() error {
	return nil
}
