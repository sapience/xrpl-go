package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/account"
)

func main() {
	client := websocket.NewClient("wss://s.altnet.rippletest.net")

	acr, _, err := client.Account.GetAccountInfo(&account.AccountInfoRequest{Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59"})
	if err != nil {
		panic(err)
	}
	fmt.Println(acr)
	// Do stuff

}
