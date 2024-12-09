package data

type LogrotateRequest struct {
}

func (*LogrotateRequest) Method() string {
	return "logrotate"
}

type LogrotateResponse struct {
	Message string `json:"message"`
}
