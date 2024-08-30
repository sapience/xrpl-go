package websocket

import (
	"bytes"
	"encoding/json"
	"errors"
	"sync/atomic"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/xrpl/client"
	requests "github.com/Peersyst/xrpl-go/xrpl/model/requests/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/gorilla/websocket"
)

var _ client.Client = (*WebsocketClient)(nil)

var ErrIncorrectId = errors.New("incorrect id")

type WebsocketConfig struct {
	URL string
}

type WebsocketClient struct {
	cfg       *WebsocketConfig
	idCounter atomic.Uint32
}

func (c *WebsocketClient) SendRequest(req client.XRPLRequest) (client.XRPLResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id := c.idCounter.Add(1)

	conn, _, err := websocket.DefaultDialer.Dial(c.cfg.URL, nil)
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

func (c *WebsocketClient) SubmitRequest(req *requests.SubmitRequest) (*requests.SubmitResponse, error) {
	res, err := c.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var sr requests.SubmitResponse
	err = res.GetResult(&sr)
	if err != nil {
		return nil, err
	}
	return &sr, nil
}

func (c *WebsocketClient) SubmitTransactionBlob(txBlob string, failHard bool) (*requests.SubmitResponse, error) {
	tx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return nil, err
	}

	_, okTxSig := tx["TxSignature"].(string)
	_, okSigPubKey := tx["SigningPubKey"].(string)
	signers, okSigners := tx["Signers"].([]transactions.Signer)

	if okSigners {
		for _, signer := range signers {
			if signer.SignerData.SigningPubKey == "" && signer.SignerData.TxnSignature == "" {
				return nil, errors.New("transaction is not signed")
			}
		}
	} else if !okTxSig && !okSigPubKey {
		return nil, errors.New("transaction is not signed")
	}

	submitRequest := &requests.SubmitRequest{
		TxBlob:   txBlob,
		FailHard: failHard,
	}

	return c.SubmitRequest(submitRequest)
}

// Creates a new websocket client with cfg.
//
// This client will open and close a websocket connection for each request.
func NewWebsocketClient(cfg *WebsocketConfig) *WebsocketClient {
	return &WebsocketClient{
		cfg: cfg,
	}
}

func NewClient(cfg *WebsocketConfig) *client.XRPLClient {
	wcl := &WebsocketClient{
		cfg: cfg,
	}
	return client.NewXRPLClient(wcl)
}
