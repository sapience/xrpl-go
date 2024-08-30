package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

func main() {
	client := websocket.NewWebsocketClient(&websocket.WebsocketConfig{URL: "wss://s.altnet.rippletest.net"})

	payment := transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: "rLY96NyP8Wq5yX5NQ3XdeZdUyUFBRbWNgd",
		},
		Destination: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		Amount:      types.XRPCurrencyAmount(100000),
	}
	payment.SetPartialPaymentFlag()
	payment.SetLimitQualityFlag()
	payment.SetRippleNotDirectFlag()

	paymentFlat := payment.Flatten()

	err := client.Autofill(&paymentFlat)
	if err != nil {
		panic(err)
	}

	fmt.Println(paymentFlat)
}
