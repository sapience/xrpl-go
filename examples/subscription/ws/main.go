package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	subscribe "github.com/Peersyst/xrpl-go/xrpl/queries/subscription"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

func main() {
	fmt.Println("⏳ Connecting to testnet...")
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.altnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	defer client.Disconnect()

	if err := client.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	if !client.IsConnected() {
		fmt.Println("❌ Failed to connect to testnet")
		return
	}

	fmt.Println("✅ Connected to testnet")
	fmt.Println()


	res, err := client.Subscribe(&subscribe.Request{
		Streams: []string{"ledger"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	msgChan := make(chan []byte)

	err = client.SubscribeWS(msgChan)
	if err != nil {
		fmt.Println(err)
		return
	}

	for msg := range msgChan {
		fmt.Println(string(msg))
	}
}