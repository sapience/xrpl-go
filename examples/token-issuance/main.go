package main

import (
	"fmt"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

const (
	currencyCode = "FOO"
)

func main() {
	//
	// Configure client
	//
	client := websocket.NewWebsocketClient(
		websocket.NewWebsocketClientConfig().
		WithHost("wss://s.altnet.rippletest.net").
		WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)

	//
	// Configure wallets
	//
	coldWallet, err := xrpl.NewWallet(addresscodec.ED25519)
	if err != nil {
		fmt.Printf("Error creating cold wallet: %s\n", err)
		return
	}
	hotWallet, err := xrpl.NewWallet(addresscodec.ED25519)
	if err != nil {
		fmt.Printf("Error creating hot wallet: %s\n", err)
		return
	}

	customerOneWallet, err := xrpl.NewWallet(addresscodec.ED25519)
	if err != nil {
		fmt.Printf("Error creating token wallet: %s\n", err)
		return
	}

	customerTwoWallet, err := xrpl.NewWallet(addresscodec.ED25519)
	if err != nil {
		fmt.Printf("Error creating token wallet: %s\n", err)
		return
	}

	//
	// Configure cold address settings
	//
	coldWalletAccountSet := &transactions.AccountSet{
		BaseTx: transactions.BaseTx{
			Account: types.Address(coldWallet.ClassicAddress),
		},
		TickSize: 5,
		TransferRate: 0,
		Domain: "6578616D706C652E636F6D", // example.com
		SetFlag: 0,
	}

	coldWalletAccountSet.SetDisallowXRP()
	coldWalletAccountSet.SetRequireDestTag()

	flattenedTx := coldWalletAccountSet.Flatten()

	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("Error autfilling transaction: %s\n", err)
		return
	}

	txBlob, _, err := coldWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("Error signing transaction: %s\n", err)
		return
	}

	response, err := client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("Error submitting transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction submitted: %s\n", response.Tx["hash"])

	//
	// Configure hot address settings
	//
	hotWalletAccountSet := &transactions.AccountSet{
		BaseTx: transactions.BaseTx{
			Account: types.Address(hotWallet.ClassicAddress),
		},
		Domain: "6578616D706C652E636F6D", // example.com
	}

	hotWalletAccountSet.SetDisallowXRP()
	hotWalletAccountSet.SetRequireDestTag()

	flattenedTx = hotWalletAccountSet.Flatten()
	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err = hotWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("Error signing transaction: %s\n", err)
		return
	}

	response, err = client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("Error submitting transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction submitted: %s\n", response.Tx["hash"])

	//
	// Create trust line from hot to cold address
	//
	hotColdTrustSet := &transactions.TrustSet{
		BaseTx: transactions.BaseTx{
			Account: types.Address(hotWallet.ClassicAddress),
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Currency: currencyCode,
			Issuer: types.Address(coldWallet.ClassicAddress),
			Value: "100000000000000",
		},
	}

	flattenedTx = hotColdTrustSet.Flatten()
	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err = hotWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("Error signing transaction: %s\n", err)
		return
	}

	response, err = client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("Error submitting transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction submitted: %s\n", response.Tx["hash"])

	//
	// Create trust line from costumer one to cold address
	//
	customerOneColdTrustSet := &transactions.TrustSet{
		BaseTx: transactions.BaseTx{
			Account: types.Address(customerOneWallet.ClassicAddress),
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Currency: currencyCode,
			Issuer: types.Address(coldWallet.ClassicAddress),
			Value: "100000000000000",
		},
	}

	flattenedTx = customerOneColdTrustSet.Flatten()
	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err = customerOneWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("Error signing transaction: %s\n", err)
		return
	}

	response, err = client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("Error submitting transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction submitted: %s\n", response.Tx["hash"])

	//
	// Create trust line from costumer two to cold address
	//
	customerTwoColdTrustSet := &transactions.TrustSet{
		BaseTx: transactions.BaseTx{
			Account: types.Address(customerTwoWallet.ClassicAddress),
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Currency: currencyCode,
			Issuer: types.Address(coldWallet.ClassicAddress),
			Value: "100000000000000",
		},
	}

	flattenedTx = customerTwoColdTrustSet.Flatten()
	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err = customerTwoWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("Error signing transaction: %s\n", err)
		return
	}

	response, err = client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("Error submitting transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction submitted: %s\n", response.Tx["hash"])

	//
	// Send tokens from cold wallet to hot wallet
	//
	coldToHotPayment := &transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: types.Address(coldWallet.ClassicAddress),
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: currencyCode,
			Issuer: types.Address(coldWallet.ClassicAddress),
			Value: "3800",
		},
		Destination: types.Address(hotWallet.ClassicAddress),
		DestinationTag: 1,
	}

	flattenedTx = coldToHotPayment.Flatten()
	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err = coldWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("Error signing transaction: %s\n", err)
		return
	}

	response, err = client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("Error submitting transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction submitted: %s\n", response.Tx["hash"])

	//
	// Send tokens from hot wallet to customer one
	//
	hotToCustomerOnePayment := &transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: types.Address(hotWallet.ClassicAddress),
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: currencyCode,
			Issuer: types.Address(coldWallet.ClassicAddress),
			Value: "100",
		},
		Destination: types.Address(customerOneWallet.ClassicAddress),
		DestinationTag: 1,
	}


	flattenedTx = hotToCustomerOnePayment.Flatten()
	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("Error autofilling transaction: %s\n", err)
		return
	}
	
	txBlob, _, err = hotWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("Error signing transaction: %s\n", err)
		return
	}

	response, err = client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("Error submitting transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction submitted: %s\n", response.Tx["hash"])
}