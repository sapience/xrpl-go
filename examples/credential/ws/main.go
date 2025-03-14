package main

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	rippleTime "github.com/Peersyst/xrpl-go/xrpl/time"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

type SubmittableTransaction interface {
	TxType() transaction.TxType
	Flatten() transaction.FlatTransaction // Ensures all transactions can be flattened
}

func main() {

	fmt.Println("⏳ Setting up client...")

	client := getClient()
	fmt.Println("Connecting to server...")
	if err := client.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ Client configured!")
	fmt.Println()

	fmt.Printf("Connection: %t", client.IsConnected())
	fmt.Println()

	// Configure wallets

	// Issuer
	fmt.Println("⏳ Setting up credential issuer wallet...")
	issuer, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Printf("❌ Error creating issuer wallet: %s\n", err)
		return
	}

	err = client.FundWallet(&issuer)
	if err != nil {
		fmt.Printf("❌ Error funding issuer wallet: %s\n", err)
		return
	}
	fmt.Printf("✅ Issuer wallet funded: %s\n", issuer.ClassicAddress)

	// -----------------------------------------------------

	// Subject (destination)
	fmt.Println("⏳ Setting up Subject wallet...")
	subjectWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Printf("❌ Error creating subject wallet: %s\n", err)
		return
	}

	err = client.FundWallet(&subjectWallet)
	if err != nil {
		fmt.Printf("❌ Error funding subject wallet: %s\n", err)
		return
	}
	fmt.Printf("✅ Subject wallet funded: %s\n", subjectWallet.ClassicAddress)

	// -----------------------------------------------------

	// Creating the CredentialCreate transaction
	fmt.Println("⏳ Creating CredentialCreate transaction...")

	expiration, err := rippleTime.IsoTimeToRippleTime(time.Now().Add(time.Hour * 24).Format(time.RFC3339))
	credentialType := types.CredentialType("6D795F63726564656E7469616C")

	if err != nil {
		fmt.Printf("❌ Error converting expiration to ripple time: %s\n", err)
		return
	}

	txn := &transaction.CredentialCreate{
		BaseTx: transaction.BaseTx{
			Account: types.Address(issuer.ClassicAddress),
		},
		CredentialType: credentialType,
		Subject:        types.Address(subjectWallet.ClassicAddress),
		Expiration:     uint32(expiration),
		URI:            hex.EncodeToString([]byte("https://example.com")),
	}

	submitAndWait(client, txn, issuer)

	// -----------------------------------------------------

	// Creating the CredentialAccept transaction
	fmt.Println("⏳ Creating CredentialAccept transaction...")

	acceptTxn := &transaction.CredentialAccept{
		BaseTx: transaction.BaseTx{
			Account: types.Address(subjectWallet.ClassicAddress),
		},
		CredentialType: credentialType,
		Issuer:         types.Address(issuer.ClassicAddress),
	}

	submitAndWait(client, acceptTxn, subjectWallet)

	// -----------------------------------------------------

	// Creating the CredentialDelete transaction
	fmt.Println("⏳ Creating CredentialDelete transaction...")

	deleteTxn := &transaction.CredentialDelete{
		BaseTx: transaction.BaseTx{
			Account: types.Address(issuer.ClassicAddress),
		},
		CredentialType: credentialType,
		Issuer:         types.Address(issuer.ClassicAddress),
		Subject:        types.Address(subjectWallet.ClassicAddress),
	}

	submitAndWait(client, deleteTxn, issuer)
}

// getRpcClient returns a new rpc client
func getClient() *websocket.Client {
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.devnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewDevnetFaucetProvider()),
	)

	return client
}

// submitAndWait submits a transaction and waits for it to be included in a validated ledger
func submitAndWait(client *websocket.Client, txn SubmittableTransaction, wallet wallet.Wallet) {
	fmt.Printf("⏳ Submitting %s transaction...\n", txn.TxType())

	flattenedTx := txn.Flatten()

	err := client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error autofilling %s transaction: %s\n", txn.TxType(), err)
		fmt.Println()
		return
	}

	fmt.Println("flattenedTx", flattenedTx)

	txBlob, _, err := wallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error signing %s transaction: %s\n", txn.TxType(), err)
		fmt.Println()
		return
	}

	response, err := client.SubmitAndWait(txBlob, false)
	if err != nil {
		fmt.Printf("❌ Error submitting %s transaction: %s\n", txn.TxType(), err)
		fmt.Println()
		return
	}

	fmt.Printf("✅ %s transaction submitted\n", txn.TxType())
	fmt.Printf("🌐 Hash: %s\n", response.Hash.String())
	fmt.Println()
}
