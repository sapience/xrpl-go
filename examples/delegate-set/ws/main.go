package main

import (
	"fmt"
	"time"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

func main() {
	// Connect to testnet
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

	// Create and fund wallets
	delegatorWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	delegateeWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("⏳ Funding wallets...")
	if err := client.FundWallet(&delegatorWallet); err != nil {
		fmt.Println(err)
		return
	}
	if err := client.FundWallet(&delegateeWallet); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("💸 Wallets funded")

	// Wait for accounts to be available
	time.Sleep(5 * time.Second)

	// Check initial balances
	delegatorBalance := getAccountBalance(client, delegatorWallet.ClassicAddress)
	delegateeBalance := getAccountBalance(client, delegateeWallet.ClassicAddress)

	fmt.Printf("💳 Delegator initial balance: %s XRP\n", delegatorBalance)
	fmt.Printf("💳 Delegatee initial balance: %s XRP\n", delegateeBalance)
	fmt.Println()

	// Create DelegateSet transaction
	delegateSetTx := &transactions.DelegateSet{
		BaseTx: transactions.BaseTx{
			Account: types.Address(delegatorWallet.ClassicAddress),
		},
		Authorize: types.Address(delegateeWallet.ClassicAddress),
		Permissions: []types.Permission{
			{
				Permission: types.PermissionValue{
					PermissionValue: "Payment",
				},
			},
		},
	}

	// Submit DelegateSet transaction
	flattenedTx := delegateSetTx.Flatten()
	if err := client.Autofill(&flattenedTx); err != nil {
		fmt.Println(err)
		return
	}

	txBlob, _, err := delegatorWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := client.SubmitTxBlobAndWait(txBlob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ DelegateSet transaction submitted")
	fmt.Printf("🌐 Hash: %s\n", response.Hash)
	fmt.Printf("🌐 Validated: %t\n", response.Validated)
	fmt.Println()

	// Wait for delegation to be processed
	time.Sleep(3 * time.Second)

	// Create delegated payment transaction
	delegatedPaymentTx := &transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account:  types.Address(delegatorWallet.ClassicAddress),
			Delegate: types.Address(delegateeWallet.ClassicAddress),
		},
		Destination: types.Address(delegateeWallet.ClassicAddress),
		Amount:      types.XRPCurrencyAmount(1000000), // 1 XRP
	}

	// Submit delegated payment
	flatDelegatedPayment := delegatedPaymentTx.Flatten()
	if err := client.Autofill(&flatDelegatedPayment); err != nil {
		fmt.Println(err)
		return
	}

	txBlob2, _, err := delegateeWallet.Sign(flatDelegatedPayment)
	if err != nil {
		fmt.Println(err)
		return
	}

	response2, err := client.SubmitTxBlobAndWait(txBlob2, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ Delegated payment submitted")
	fmt.Printf("🌐 Hash: %s\n", response2.Hash)
	fmt.Printf("� Validated: %t\n", response2.Validated)
	fmt.Println()

	// Check final balances
	finalDelegatorBalance := getAccountBalance(client, delegatorWallet.ClassicAddress)
	finalDelegateeBalance := getAccountBalance(client, delegateeWallet.ClassicAddress)

	fmt.Printf("💳 Delegator final balance: %s XRP\n", finalDelegatorBalance)
	fmt.Printf("💳 Delegatee final balance: %s XRP\n", finalDelegateeBalance)
}

func getAccountBalance(client *websocket.Client, address types.Address) string {
	accountInfo, err := client.GetAccountInfo(&account.InfoRequest{
		Account: address,
	})
	if err != nil {
		return "0"
	}

	// Convert drops to XRP (1 XRP = 1,000,000 drops)
	balanceDrops := accountInfo.AccountData.Balance
	balanceXRP := float64(balanceDrops) / 1000000.0

	return fmt.Sprintf("%.6f", balanceXRP)
}
