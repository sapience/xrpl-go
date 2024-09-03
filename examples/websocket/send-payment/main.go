package main

import (
	"fmt"
	"strconv"

	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/utils"
)

func main() {
	wallet, err := xrpl.NewWalletFromSeed("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	receiverWallet, err := xrpl.NewWalletFromSeed("sEd7d8Ci9nevdLCeUMctF3uGXp9WQqJ", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := websocket.NewWebsocketClient(
		websocket.NewWebsocketClientConfig().
			WithHost("wss://s.altnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)

	balance, err := client.GetXrpBalance(wallet.GetAddress())

	if err != nil || balance == "0" {
		fmt.Println("Balance: 0")
		fmt.Println("Funding wallet")
		err = client.FundWallet(&wallet)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	balance, _ = client.GetXrpBalance(wallet.GetAddress())

	fmt.Println("Balance: ", balance)

	amount, err := utils.XrpToDrops("1")
	if err != nil {
		fmt.Println(err)
		return
	}

	amountUint, err := strconv.ParseUint(amount, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	payment := transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: types.Address(wallet.GetAddress()),
		},
		Destination: types.Address(receiverWallet.GetAddress()),
		Amount:      types.XRPCurrencyAmount(amountUint),
	}

	flatTx := payment.Flatten()

	err = client.Autofill(&flatTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transaction autofilled")

	txBlob, _, err := wallet.Sign(flatTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transaction signed")
	fmt.Println("Transaction submitted\n")

	response, err := client.SubmitTransactionBlob(txBlob, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transaction engine result:", response.EngineResult)
	fmt.Println("Transaction accepted:", response.Accepted)
	fmt.Println("Transaction hash:", response.Tx["hash"])
}
