package helpers

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

type SubmittableTransaction interface {
	TxType() transactions.TxType
	Flatten() transactions.FlatTransaction // Ensures all transactions can be flattened
}

// SubmitAndWait submits a transaction and waits for it to be included in a validated ledger
func SubmitAndWait(client *rpc.Client, txn SubmittableTransaction, wallet wallet.Wallet) {
	fmt.Println()
	fmt.Printf("⏳ Submitting %s transaction...\n", txn.TxType())

	flattenedTx := txn.Flatten()

	err := client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error autofilling %s transaction: %s\n", txn.TxType(), err)
		fmt.Println()
		return
	}

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
