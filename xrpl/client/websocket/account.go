package websocket

import (
	"github.com/Peersyst/xrpl-go/xrpl/client"
	"github.com/Peersyst/xrpl-go/xrpl/model/client/account"
)

func (c *WebsocketClient) GetAccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, client.XRPLResponse, error) {
	res, err := c.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var air account.AccountInfoResponse
	err = res.GetResult(&air)
	if err != nil {
		return nil, nil, err
	}
	return &air, res, nil
}

func (c *WebsocketClient) GetAccountObjects(req *account.AccountObjectsRequest) (*account.AccountObjectsResponse, error) {
	res, err := c.SendRequest(req)
	if err != nil {
		return nil, err
	}
	var acr account.AccountObjectsResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}
