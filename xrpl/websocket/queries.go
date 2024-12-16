package websocket

import (
	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/queries/channel"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/ledger"
	"github.com/Peersyst/xrpl-go/xrpl/queries/nft"
	"github.com/Peersyst/xrpl-go/xrpl/queries/path"
	"github.com/Peersyst/xrpl-go/xrpl/queries/server"
	"github.com/Peersyst/xrpl-go/xrpl/queries/utility"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// Account queries

// GetAccountInfo retrieves information about an account on the XRP Ledger.
// It takes an AccountInfoRequest as input and returns an AccountInfoResponse,
// along with the raw XRPL response and any error encountered.
func (c *Client) GetAccountInfo(req *account.InfoRequest) (*account.InfoResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var air account.InfoResponse
	err = res.GetResult(&air)
	if err != nil {
		return nil, err
	}
	return &air, nil
}

// GetAccountObjects retrieves a list of objects owned by an account on the XRP Ledger.
// It takes an AccountObjectsRequest as input and returns an AccountObjectsResponse,
// along with any error encountered.
func (c *Client) GetAccountObjects(req *account.ObjectsRequest) (*account.ObjectsResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var acr account.ObjectsResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

// GetXrpBalance retrieves the XRP balance of a given account address.
// It returns the balance as a string in XRP (not drops) and any error encountered.
func (c *Client) GetXrpBalance(address types.Address) (string, error) {
	res, err := c.GetAccountInfo(&account.InfoRequest{
		Account: address,
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

// GetAccountLines retrieves the lines associated with an account on the XRP Ledger.
// It takes an AccountLinesRequest as input and returns an AccountLinesResponse,
// along with any error encountered.
func (c *Client) GetAccountLines(req *account.LinesRequest) (*account.LinesResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var acr account.LinesResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

func (c *Client) GetAccountNFTs(req *account.NFTsRequest) (*account.NFTsResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var acr account.NFTsResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

func (c *Client) GetAccountCurrencies(req *account.CurrenciesRequest) (*account.CurrenciesResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var acr account.CurrenciesResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

func (c *Client) GetAccountOffers(req *account.OffersRequest) (*account.OffersResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var acr account.OffersResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

func (c *Client) GetAccountTransactions(req *account.TransactionsRequest) (*account.TransactionsResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var acr account.TransactionsResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

// Channel queries

func (c *Client) GetChannelVerify(req *channel.VerifyRequest) (*channel.VerifyResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var acr channel.VerifyResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

// Ledger queries

// Returns the index of the most recently validated ledger.
func (c *Client) GetLedgerIndex() (common.LedgerIndex, error) {
	res, err := c.Request(&ledger.Request{
		LedgerIndex: common.LedgerTitle("validated"),
	})
	if err != nil {
		return 0, err
	}

	var lr ledger.Response
	err = res.GetResult(&lr)
	if err != nil {
		return 0, err
	}
	return lr.LedgerIndex, err
}

func (c *Client) GetClosedLedger() (*ledger.ClosedResponse, error) {
	res, err := c.Request(&ledger.ClosedRequest{})
	if err != nil {
		return nil, err
	}
	var lr ledger.ClosedResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetCurrentLedger() (*ledger.CurrentResponse, error) {
	res, err := c.Request(&ledger.CurrentRequest{})
	if err != nil {
		return nil, err
	}
	var lr ledger.CurrentResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetLedgerData(req *ledger.DataRequest) (*ledger.DataResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr ledger.DataResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetLedger(req *ledger.Request) (*ledger.Response, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr ledger.Response
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

// NFT queries

func (c *Client) GetNFTBuyOffers(req *nft.NFTokenBuyOffersRequest) (*nft.NFTokenBuyOffersResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr nft.NFTokenBuyOffersResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetNFTSellOffers(req *nft.NFTokenSellOffersRequest) (*nft.NFTokenSellOffersResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr nft.NFTokenSellOffersResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

// Path queries

func (c *Client) GetBookOffers(req *path.BookOffersRequest) (*path.BookOffersResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr path.BookOffersResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetDepositAuthorized(req *path.DepositAuthorizedRequest) (*path.DepositAuthorizedResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr path.DepositAuthorizedResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) FindPathCreate(req *path.FindCreateRequest) (*path.FindResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr path.FindResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) FindPathClose(req *path.FindCloseRequest) (*path.FindResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr path.FindResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) FindPathStatus(req *path.FindStatusRequest) (*path.FindResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr path.FindResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetRipplePathFind(req *path.RipplePathFindRequest) (*path.RipplePathFindResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr path.RipplePathFindResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

// Server queries

// GetServerInfo retrieves information about the server.
// It takes a ServerInfoRequest as input and returns a ServerInfoResponse,
// along with any error encountered.
func (c *Client) GetServerInfo(req *server.InfoRequest) (*server.InfoResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var sir server.InfoResponse
	err = res.GetResult(&sir)
	if err != nil {
		return nil, err
	}
	return &sir, err
}

func (c *Client) GetAllFeatures(req *server.FeatureAllRequest) (*server.FeatureAllResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr server.FeatureAllResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetFeature(req *server.FeatureOneRequest) (*server.FeatureResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr server.FeatureResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetFee(req *server.FeeRequest) (*server.FeeResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr server.FeeResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetManifest(req *server.ManifestRequest) (*server.ManifestResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr server.ManifestResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetServerState(req *server.StateRequest) (*server.StateResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr server.StateResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

// Utility queries

func (c *Client) Ping(req *utility.PingRequest) (*utility.PingResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr utility.PingResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (c *Client) GetRandom(req *utility.RandomRequest) (*utility.RandomResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var lr utility.RandomResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}
