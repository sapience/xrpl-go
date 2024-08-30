package websocket

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/ledger"
)

const (
	// Expire unconfirmed transactions after 20 ledger versions, approximately 1 minute, by default
	LEDGER_OFFSET uint32 = 20
)

// Returns the index of the most recently validated ledger.
func (c *WebsocketClient) GetLedgerIndex() (*common.LedgerIndex, error) {
	res, err := c.SendRequest(&ledger.LedgerRequest{
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
	return &lr.LedgerIndex, err
}
