package websocket

type XRPLRequest interface {
	Method() string
	Validate() error
}
