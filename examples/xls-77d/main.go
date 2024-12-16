package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

func main() {
	fmt.Println("Connecting to devnet...")
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.devnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewDevnetFaucetProvider()),
	)
	defer client.Disconnect()

	if err := client.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	if !client.IsConnected() {
		fmt.Println("Failed to connect to testnet")
		return
	}

	fmt.Println("Connected to testnet")
	fmt.Println()

	coldWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Printf("❌ Error creating cold wallet: %s\n", err)
		return
	}

	w1, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Cold wallet:", coldWallet.GetAddress())
	fmt.Println("Wallet 1:", w1.GetAddress())
	fmt.Println()
	fmt.Println("Requesting XRP from faucet for cold wallet...")
	if err := client.FundWallet(&coldWallet); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Cold wallet funded")
	fmt.Println("Requesting XRP from faucet for wallet 1...")
	if err := client.FundWallet(&w1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Wallet 1 funded")
	fmt.Println()

	coldWalletAccountSet := &transaction.AccountSet{
		BaseTx: transaction.BaseTx{
			Account: types.Address(coldWallet.ClassicAddress),
		},
		TickSize:     5,
		TransferRate: 0,
		Domain:       "6578616D706C652E636F6D", // example.com
	}

	coldWalletAccountSet.SetAsfDefaultRipple()
	coldWalletAccountSet.SetDisallowXRP()

	coldWalletAccountSet.SetRequireDestTag()

	flattenedTx := coldWalletAccountSet.Flatten()

	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err := coldWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error signing transaction: %s\n", err)
		return
	}

	response, err := client.Submit(txBlob, false)
	if err != nil {
		fmt.Printf("❌ Error submitting transaction: %s\n", err)
		return
	}

	if response.EngineResult != "tesSUCCESS" {
		fmt.Println("❌ Cold address settings configuration failed!", response.EngineResult)
		fmt.Println("Try again!")
		fmt.Println()
		return
	}

	fmt.Println("✅ Cold address settings configured!")
	fmt.Printf("🌐 Hash: %s\n", response.Tx["hash"])
	fmt.Println()


	ts := &transaction.TrustSet{
		BaseTx: transaction.BaseTx{
			Account: types.Address(w1.ClassicAddress),
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   types.Address(coldWallet.ClassicAddress),
			Value:    "100000000",
		},
	}

	ts.SetSetFreezeFlag()
	ts.SetSetDeepFreezeFlag()

	flatTs := ts.Flatten()

	err = client.Autofill(&flatTs)
	if err != nil {
		fmt.Printf("❌ Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err = w1.Sign(flatTs)
	if err != nil {
		fmt.Printf("❌ Error signing transaction: %s\n", err)
		return
	}

	response, err = client.Submit(txBlob, false)
	if err != nil {
		fmt.Printf("❌ Error submitting transaction: %s\n", err)
		return
	}

	fmt.Println("✅ Trust set submitted!")
	fmt.Printf("🌐 Hash: %s\n", response.Tx["hash"])
	fmt.Println("Transaction result:", response.EngineResult)
	fmt.Println()
}
