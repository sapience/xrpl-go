package client

import "github.com/Peersyst/xrpl-go/xrpl/model/transactions"

type Client interface {
	SendRequest(req XRPLRequest) (XRPLResponse, error)
	Autofill(tx *transactions.FlatTransaction) error
}

type XRPLClient struct {
	client  Client
	Account Account
}

type XRPLRequest interface {
	Method() string
	Validate() error
}

type XRPLResponse interface {
	GetResult(v any) error
}

type XRPLResponseWarning struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		client:  cl,
		Account: &accountImpl{client: cl},
	}
}

func (c *XRPLClient) Client() Client {
	return c.client
}
