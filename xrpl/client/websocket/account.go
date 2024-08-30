package websocket

import (
	"github.com/Peersyst/xrpl-go/xrpl/client"
	"github.com/Peersyst/xrpl-go/xrpl/model/client/account"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/utils"
)

// GetAccountInfo retrieves information about an account on the XRP Ledger.
// It takes an AccountInfoRequest as input and returns an AccountInfoResponse,
// along with the raw XRPL response and any error encountered.
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

// GetAccountObjects retrieves a list of objects owned by an account on the XRP Ledger.
// It takes an AccountObjectsRequest as input and returns an AccountObjectsResponse,
// along with any error encountered.
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

// GetXrpBalance retrieves the XRP balance of a given account address.
// It returns the balance as a string in XRP (not drops) and any error encountered.
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
