package main

import (
	"fmt"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

const (
	AccountSeed = "sEd7MgLAff94dLx91rVRByUbLrdSrdj"
	DestinationAddress = "rDwvihpE48E48F8rvNrqTb2UGWv62xqYTg"
	Currency = "USD"
	Value = "100"
	Issuer = "rDwvihpE48E48F8rvNrqTb2UGWv62xqYTg"
)

func main() {
	wallet, err := xrpl.NewWallet(addresscodec.ED25519)
	if err != nil {
		panic(err)
	}
	fmt.Println("Wallet generated from random seed")

	fmt.Printf("Private key: %s\n", wallet.PrivateKey)
	fmt.Printf("Public 	key: %s\n", wallet.PublicKey)
	fmt.Printf("Classic address: %s\n", wallet.ClassicAddress)
	fmt.Printf("Seed: %s\n", wallet.Seed)

	walletFromSeed, _ := xrpl.NewWalletFromSeed(wallet.Seed, "")

	fmt.Println("\nWallet generated from previous seed")

	fmt.Printf("Private key: %s\n", walletFromSeed.PrivateKey)
	fmt.Printf("Public 	key: %s\n", walletFromSeed.PublicKey)
	fmt.Printf("Classic address: %s\n", walletFromSeed.ClassicAddress)
	fmt.Printf("Seed: %s\n", walletFromSeed.Seed)

	walletFromSecret, _ := xrpl.NewWalletFromSecret(wallet.Seed)

	fmt.Println("\nWallet generated from previous seed")

	fmt.Printf("Private key: %s\n", walletFromSecret.PrivateKey)
	fmt.Printf("Public 	key: %s\n", walletFromSecret.PublicKey)
	fmt.Printf("Classic address: %s\n", walletFromSecret.ClassicAddress)
	fmt.Printf("Seed: %s\n", walletFromSecret.Seed)

	fmt.Println("\nSigning a transaction")

	tx := transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: types.Address(wallet.ClassicAddress),
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   Issuer,
			Currency: Currency,
			Value:    Value,
		},
		Destination: types.Address(DestinationAddress),
	}

	fmt.Println(tx.Flatten())

	txBlob, hash, err := wallet.Sign(tx.Flatten())
	if err != nil {
		panic(err)
	}

	fmt.Printf("txBlob: %s\n", txBlob)
	fmt.Printf("hash: %s\n", hash)
}
