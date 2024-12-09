package server

type StopRequest struct {
}

func (*StopRequest) Method() string {
	return "stop"
}

type StopResponse struct {
	Message string `json:"message"`
}
