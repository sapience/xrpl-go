package wallet

import (
	"errors"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/keypairs"
	"github.com/Peersyst/xrpl-go/xrpl/hash"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
)

var (
	ErrTransactionMustBeBatch = errors.New("transaction must be a batch transaction")
	
)

type BatchTransaction struct {
	transaction.BaseTx
	RawTransactions []transaction.FlatTransaction
	Flags uint32
}

type BatchSignable struct {
	Flags uint32
	TxIDs []string
}

func (b *BatchSignable) Flatten() map[string]interface{} {
	return map[string]interface{}{
		"flags": b.Flags,
		"txIDs": b.TxIDs,
	}
}

type BatchAccount struct {
	value string
}

func (b *BatchAccount) String() string {
	return b.value
}

type SignMultiBatchOptions struct {
	BatchAccount *BatchAccount
	Multisign bool
	MultisignAccount string
}

func SignMultiBatch(wallet Wallet, tx BatchTransaction, opts SignMultiBatchOptions) error {
	var batchAccount string
	var multisignAddress string

	if opts.BatchAccount != nil {
		batchAccount = opts.BatchAccount.String()
	} else {
		batchAccount = wallet.ClassicAddress.String()
	}

	if opts.MultisignAccount != "" {
		multisignAddress = opts.MultisignAccount
	} else if opts.Multisign {
		multisignAddress = wallet.ClassicAddress.String()
	}

	if tx.TxType() != "BatchTransaction" {
		return ErrTransactionMustBeBatch
	}

	// Check batch account exists in RawTransactions.Account
	txIDs := make([]string, len(tx.RawTransactions))
	for i, rawTx := range tx.RawTransactions {
		signedTx, err := hash.SignTx(rawTx)
		if err != nil {
			return err
		}
		txIDs[i] = signedTx
	}

	payload := BatchSignable{
		Flags: tx.Flags,
		TxIDs: txIDs,
	}

	encodedBatch, err := binarycodec.EncodeForSigningBatch(payload.Flatten())
	if err != nil {
		return err
	}

	signature, err := keypairs.Sign(encodedBatch, wallet.PrivateKey)
	if err != nil {
		return err
	}

	if multisignAddress != "" {
		// Build BatchSigner with multisignAddress
	} else {
		// Build BatchSigner
	}

	return nil
}