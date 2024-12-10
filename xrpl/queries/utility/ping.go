package utility

// ############################################################################
// Request
// ############################################################################

type PingRequest struct{}

func (*PingRequest) Method() string {
	return "ping"
}

func (*PingRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type PingResponse struct {
	Role      string `json:"role,omitempty"`
	Unlimited bool   `json:"unlimited,omitempty"`
}
