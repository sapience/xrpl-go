package rpc

type JsonRpcRequest struct {
	Method string         `json:"method"`
	Params [1]interface{} `json:"params,omitempty"`
}

type XRPLRequest interface {
	Method() string
	Validate() error
}
