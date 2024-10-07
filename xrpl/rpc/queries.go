package rpc

import "github.com/Peersyst/xrpl-go/xrpl/queries/account"

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
