package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/account"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

func main() {
	wallet, err := xrpl.NewWalletFromSeed("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := websocket.NewWebsocketClient(
		websocket.NewWebsocketClientConfig().
			WithHost("wss://s.altnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)

	accountInfo, _, err := client.GetAccountInfo(&account.AccountInfoRequest{
		Account: types.Address(wallet.GetAddress()),
		SignerList: true,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Account address: ", accountInfo.AccountData.Account)
	fmt.Println("Account sequence: ", accountInfo.AccountData.Sequence)
	fmt.Println("Account balance: ", accountInfo.AccountData.Balance)
}
