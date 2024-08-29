package websocket

import (
	"bytes"
	"encoding/json"
	"errors"
	"sync/atomic"

	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/client"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/gorilla/websocket"
)

const (
	DEFAULT_FEE_CUSHION float32 = 1.2
	DEFAULT_MAX_FEE_XRP float32 = 2
)

var _ client.Client = (*WebsocketClient)(nil)

var ErrIncorrectId = errors.New("incorrect id")

type WebsocketClient struct {
	cfg WebsocketClientConfig

	idCounter atomic.Uint32
	NetworkId uint32
}

// Creates a new websocket client with cfg.
// This client will open and close a websocket connection for each request.
func NewWebsocketClient(cfg WebsocketClientConfig) *WebsocketClient {
	return &WebsocketClient{
		cfg: cfg,
	}
}

func NewClient(host string) *client.XRPLClient {
	wcl := NewWebsocketClient(
		NewWebsocketClientConfig().
			WithHost(host),
	)
	return client.NewXRPLClient(wcl)
}

func (c *WebsocketClient) Autofill(tx *map[string]interface{}) error {
	if err := c.setValidTransactionAddresses(tx); err != nil {
		return err
	}

	err := c.setTransactionFlags(tx)
	if err != nil {
		return err
	}

	if _, ok := (*tx)["NetworkID"]; !ok {
		if c.NetworkId != 0 {
			(*tx)["NetworkID"] = c.NetworkId
		}
	}
	if _, ok := (*tx)["Sequence"]; !ok {
		err := c.setTransactionNextValidSequenceNumber(tx)
		if err != nil {
			return err
		}
	}
	if _, ok := (*tx)["Fee"]; !ok {
		err := c.calculateFeePerTransactionType(tx)
		if err != nil {
			return err
		}
	}
	if _, ok := (*tx)["LastLedgerSequence"]; !ok {
		err := c.setLastLedgerSequence(tx)
		if err != nil {
			return err
		}
	}
	if txType, ok := (*tx)["TransactionType"].(transactions.TxType); ok {
		if acc, ok := (*tx)["Account"].(types.Address); txType == transactions.AccountDeleteTx && ok {
			err := c.checkAccountDeleteBlockers(acc)
			if err != nil {
				return err
			}
		}
		if txType == transactions.PaymentTx {
			err := c.checkPaymentAmounts(tx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *WebsocketClient) FundWallet(wallet *xrpl.Wallet) error {
	if wallet.ClassicAddress == "" {
		return errors.New("fund wallet: cannot fund a wallet without a classic address")
	}

	err := c.cfg.faucetProvider.FundWallet(wallet.ClassicAddress)
	if err != nil {
		return err
	}

	return nil
}

func (c *WebsocketClient) SendRequest(req client.XRPLRequest) (client.XRPLResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id := c.idCounter.Add(1)

	conn, _, err := websocket.DefaultDialer.Dial(c.cfg.host, nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	msg, err := c.formatRequest(req, int(id), nil)
	if err != nil {
		return nil, err
	}

	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return nil, err
	}

	_, v, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	jDec := json.NewDecoder(bytes.NewReader(json.RawMessage(v)))
	jDec.UseNumber()
	var res WebSocketClientXrplResponse
	err = jDec.Decode(&res)
	if err != nil {
		return nil, err
	}

	if res.ID != int(id) {
		return nil, ErrIncorrectId
	}
	if err := res.CheckError(); err != nil {
		return nil, err
	}

	return &res, nil
}
