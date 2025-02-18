package websocket

import (
	"log"

	streamtypes "github.com/Peersyst/xrpl-go/xrpl/queries/subscription/types"
)

func (c *Client) SubscribeWS(msgChan chan []byte) error {
	// Start a goroutine to read messages
	go func() {
		defer close(msgChan)
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			// Send the message to the channel
			msgChan <- message
		}
	}()
	return nil
}

func (c *Client) OnLedgerClosed(
	func (ledger *streamtypes.LedgerStream),
) {
}