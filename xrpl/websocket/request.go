package websocket

type WebsocketXRPLRequest interface {
	Method() string
	Validate() error
}
