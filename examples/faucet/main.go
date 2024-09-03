package main

import (
	"fmt"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
)

func main() {

	fmt.Println("Funding wallet on testnet:")

	testnetFaucet := faucet.NewTestnetFaucetProvider()
	testnetClientCfg := websocket.NewWebsocketClientConfig().
		WithHost("wss://s.altnet.rippletest.net:51233").
		WithFaucetProvider(testnetFaucet)
	testnetClient := websocket.NewWebsocketClient(testnetClientCfg)

	wallet, err := xrpl.NewWallet(addresscodec.ED25519)
	if err != nil {
		fmt.Println(err)
		return
	}

	balance, err := testnetClient.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		balance = "0"
	}

	fmt.Println("Balance", wallet.ClassicAddress, balance)

	fmt.Println("Funding wallet", wallet.ClassicAddress)
	err = testnetClient.FundWallet(&wallet)
	if err != nil {
		fmt.Println(err)
		return
	}

	balance, err = testnetClient.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Balance", wallet.ClassicAddress, balance)

	fmt.Println("Funding wallet on devnet:")

	devnetFaucet := faucet.NewDevnetFaucetProvider()

	clientCfg := websocket.NewWebsocketClientConfig().
		WithHost("wss://s.devnet.rippletest.net:51233").
		WithFaucetProvider(devnetFaucet)

	devnetClient := websocket.NewWebsocketClient(clientCfg)

	balance, err = devnetClient.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		balance = "0"
	}

	fmt.Println("Balance", wallet.ClassicAddress, balance)

	fmt.Println("Funding wallet", wallet.ClassicAddress)
	err = devnetClient.FundWallet(&wallet)
	if err != nil {
		fmt.Println(err)
		return
	}

	balance, err = devnetClient.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Balance", wallet.ClassicAddress, balance)
}
