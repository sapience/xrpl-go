package main

import (
	"fmt"
	"log"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

func main() {
	// Configure the client
	cfg, err := rpc.NewClientConfig(
		"https://s.altnet.rippletest.net:51234/",
		rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}
	client := rpc.NewClient(cfg)
	fmt.Println("Client configured.")

	// Create and fund the cold wallet (issuer)
	issuerWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println("Error creating cold wallet:", err)
		return
	}
	if err := client.FundWallet(&issuerWallet); err != nil {
		fmt.Println("Error funding cold wallet:", err)
		return
	}

	// Create and fund the hot wallet (holder)
	hotWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println("Error creating hot wallet:", err)
		return
	}
	if err := client.FundWallet(&hotWallet); err != nil {
		fmt.Println("Error funding hot wallet:", err)
		return
	}

	// Create and fund a customer wallet
	customerWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println("Error creating customer wallet:", err)
		return
	}
	if err := client.FundWallet(&customerWallet); err != nil {
		fmt.Println("Error funding customer wallet:", err)
		return
	}

	// Output wallet addresses
	fmt.Println("Wallets configured:")
	fmt.Println("  Cold wallet:", issuerWallet.ClassicAddress)
	fmt.Println("  Hot wallet:", hotWallet.ClassicAddress)
	fmt.Println("  Customer wallet:", customerWallet.ClassicAddress)


		// Create the MPTokenIssuanceCreate transaction.
	issuanceTx := &transactions.MPTokenIssuanceCreate{
		BaseTx: transactions.BaseTx{
			Account: types.Address(issuerWallet.ClassicAddress),
		},
		AssetScale:      2,
		TransferFee:     314,
		MaximumAmount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   types.Address(issuerWallet.ClassicAddress),
			Value:    "50000000",
		},
		MPTokenMetadata: "464F4F", // Hex for "FOO"
	}
	// Since TransferFee is provided, enable the tfMPTCanTransfer flag.
	issuanceTx.SetMPTCanTransferFlag()

	// Flatten, autofill, sign, and submit the transaction.
	flattenedTx := issuanceTx.Flatten()
	if err := client.Autofill(&flattenedTx); err != nil {
		log.Fatal("Error autofilling issuance transaction:", err)
	}

	txBlob, _, err := issuerWallet.Sign(flattenedTx)
	if err != nil {
		log.Fatal("Error signing issuance transaction:", err)
	}

	response, err := client.SubmitAndWait(txBlob, false)
	if err != nil {
		log.Fatal("Error submitting issuance transaction:", err)
	}

	if !response.Validated {
		log.Fatal("MPToken issuance transaction failed to validate!")
	}

	fmt.Println("MPToken Issuance Create transaction validated!")
	fmt.Println("Transaction Hash:", response.Hash.String())
}
