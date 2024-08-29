package websocket

import (
	"github.com/Peersyst/xrpl-go/xrpl/client"
	"github.com/Peersyst/xrpl-go/xrpl/model/client/account"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/utils"
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

func (c *WebsocketClient) GetXrpBalance(address string) (string, error) {
	res, _, err := c.GetAccountInfo(&account.AccountInfoRequest{
		Account: types.Address(address),
	})
	if err != nil {
		return "", err 
	}

	xrpBalance, err := utils.DropsToXrp(res.AccountData.Balance.String())
	if err != nil {
		return "", err
	}

	return xrpBalance, nil
}
