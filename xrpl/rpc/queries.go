package rpc

import (
	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/ledger"
	"github.com/Peersyst/xrpl-go/xrpl/queries/server"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func (r *JsonRpcClient) GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, error) {
	res, err := r.SendRequest(req)
	if err != nil {
		return nil, err
	}
	var acr account.AccountChannelsResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

// GetAccountInfo retrieves information about an account on the XRP Ledger.
// It takes an AccountInfoRequest as input and returns an AccountInfoResponse,
// along with the raw XRPL response and any error encountered.
func (r *JsonRpcClient) GetAccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, error) {
	res, err := r.SendRequest(req)
	if err != nil {
		return nil, err
	}
	var air account.AccountInfoResponse
	err = res.GetResult(&air)
	if err != nil {
		return nil, err
	}
	return &air, nil
}

// GetAccountObjects retrieves a list of objects owned by an account on the XRP Ledger.
// It takes an AccountObjectsRequest as input and returns an AccountObjectsResponse,
// along with any error encountered.
func (r *JsonRpcClient) GetAccountObjects(req *account.AccountObjectsRequest) (*account.AccountObjectsResponse, error) {
	res, err := r.SendRequest(req)
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

// GetAccountLines retrieves the lines associated with an account on the XRP Ledger.
// It takes an AccountLinesRequest as input and returns an AccountLinesResponse,
// along with any error encountered.
func (r *JsonRpcClient) GetAccountLines(req *account.AccountLinesRequest) (*account.AccountLinesResponse, error) {
	res, err := r.SendRequest(req)
	if err != nil {
		return nil, err
	}
	var alr account.AccountLinesResponse
	err = res.GetResult(&alr)
	if err != nil {
		return nil, err
	}
	return &alr, nil
}

// GetXrpBalance retrieves the XRP balance of a given account address.
// It returns the balance as a string in XRP (not drops) and any error encountered.
func (r *JsonRpcClient) GetXrpBalance(address string) (string, error) {
	res, err := r.GetAccountInfo(&account.AccountInfoRequest{
		Account: types.Address(address),
	})
	if err != nil {
		return "", err
	}
	xrpBalance, err := currency.DropsToXrp(res.AccountData.Balance.String())
	if err != nil {
		return "", err
	}
	return xrpBalance, nil
}

// Returns the index of the most recently validated ledger.
func (r *JsonRpcClient) GetLedgerIndex() (*common.LedgerIndex, error) {
	res, err := r.SendRequest(&ledger.LedgerRequest{
		LedgerIndex: common.LedgerTitle("validated"),
	})
	if err != nil {
		return nil, err
	}
	var lr ledger.LedgerResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr.LedgerIndex, nil
}

// GetServerInfo retrieves information about the server.
// It takes a ServerInfoRequest as input and returns a ServerInfoResponse,
// along with any error encountered.
func (r *JsonRpcClient) GetServerInfo(req *server.ServerInfoRequest) (*server.ServerInfoResponse, error) {
	res, err := r.SendRequest(req)
	if err != nil {
		return nil, err
	}
	var sir server.ServerInfoResponse
	err = res.GetResult(&sir)
	if err != nil {
		return nil, err
	}
	return &sir, nil
}
