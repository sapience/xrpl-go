package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

func main() {

	wsClient:= websocket.NewWebsocketClient(&websocket.WebsocketConfig{URL: "wss://s.altnet.rippletest.net"})

	payment := transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: types.Address("rhKy9bFVTTZh7TAVvqnbULUZRdtH9dWZBr"),
		},
		Amount: types.XRPCurrencyAmount(100),
		Destination: types.Address("rwMEfPmJSCauyu4N3XWEc3XKCMwi5uYQiW"),
	}

	tx := payment.Flatten()

	fmt.Println("Transaction before autofill", tx)

	err := wsClient.Autofill(&tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction after autofill", tx)
}