package main

import (
	"fmt"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl"
)

const (
	AccountSeed = "sEd7MgLAff94dLx91rVRByUbLrdSrdj"
	DestinationAddress = "rDwvihpE48E48F8rvNrqTb2UGWv62xqYTg"
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

	walletFromMnemonic, _ := xrpl.NewWalletFromMnemonic("monster march exile fee forget response seven push dragon oil clinic attack black miss craft surface patient stomach tank float cabbage visual image resource")

	fmt.Println("\nWallet generated from mnemonic")

	fmt.Printf("Private key: %s\n", walletFromMnemonic.PrivateKey)
	fmt.Printf("Public 	key: %s\n", walletFromMnemonic.PublicKey)
	fmt.Printf("Classic address: %s\n", walletFromMnemonic.ClassicAddress)
	fmt.Printf("Seed: %s\n", walletFromMnemonic.Seed)

	fmt.Println("\nSigning a transaction")

	tx := map[string]any{
		"Account":         wallet.ClassicAddress,
		"TransactionType": "Payment",
		"Amount":          "15",
		"Destination":     DestinationAddress,
		"Flags":           0,
		"Fee":             "12",
		"Sequence":        1798962,
		"SigningPubKey":   wallet.PublicKey,
	}

	txBlob, hash, err := wallet.Sign(tx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("txBlob: %s\n", txBlob)
	fmt.Printf("hash: %s\n", hash)	
}
