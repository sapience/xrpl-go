package wallet

import (
	"cmp"
	"errors"
	"slices"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/keypairs"
	"github.com/Peersyst/xrpl-go/xrpl/hash"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	wallettypes "github.com/Peersyst/xrpl-go/xrpl/wallet/types"
)

var (
	// ErrBatchAccountNotFound is returned when the batch account is not found in the transaction.
	ErrBatchAccountNotFound = errors.New("batch account not found in transaction")
	// ErrTransactionMustBeBatch is returned when the transaction is not a batch transaction.
	ErrTransactionMustBeBatch = errors.New("transaction must be a batch transaction")
	// ErrNoTransactionsProvided is returned when no transactions are provided.
	ErrNoTransactionsProvided = errors.New("no transactions provided")
	// ErrTxMustIncludeBatchSigner is returned when the transaction does not include a batch signer.
	ErrTxMustIncludeBatchSigner = errors.New("transaction must include a batch signer")
	// ErrTransactionAlreadySigned is returned when the transaction has already been signed.
	ErrTransactionAlreadySigned = errors.New("transaction has already been signed")
	// ErrBatchSignableNotEqual is returned when the batch signable is not equal.
	ErrBatchSignableNotEqual = errors.New("batch signable is not equal")
)

// SignMultiBatchOptions is a set of options for signing a multi-account Batch transaction.
// BatchAccount is the account that will be used to sign the batch transaction.
// Multisign is a boolean that indicates if the wallet should be used as a multisign account.
// MultisignAccount is the account that will be used to sign the batch transaction.
type SignMultiBatchOptions struct {
	BatchAccount     *wallettypes.BatchAccount
	Multisign        bool
	MultisignAccount string
}

// Sign a multi-account Batch transaction.
// It takes a wallet, a batch transaction, and a set of options.
// It returns an error if the transaction is invalid.
func SignMultiBatch(wallet Wallet, tx *transaction.Batch, opts *SignMultiBatchOptions) error {
	var batchAccount string
	var multisignAddress string

	if opts != nil {
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
	}

	// Check batch account exists in RawTransactions.Account
	batchAccountExists := false
	for _, rawTx := range tx.RawTransactions {
		if acc, ok := rawTx.RawTransaction["Account"]; ok && acc == batchAccount {
			batchAccountExists = true
			break
		}
	}

	if !batchAccountExists {
		return ErrBatchAccountNotFound
	}

	txIDs := make([]string, len(tx.RawTransactions))
	for i, rawTx := range tx.RawTransactions {
		signedTx, err := hash.SignTx(rawTx.RawTransaction)
		if err != nil {
			return err
		}
		txIDs[i] = signedTx
	}

	payload := wallettypes.BatchSignable{
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

	var batchSigner types.BatchSigner

	if multisignAddress != "" {
		batchSigner = types.BatchSigner{
			BatchSigner: types.BatchSignerData{
				Account: types.Address(batchAccount),
				Signers: []types.Signer{
					{
						SignerData: types.SignerData{
							Account:       types.Address(multisignAddress),
							SigningPubKey: wallet.PublicKey,
							TxnSignature:  signature,
						},
					},
				},
			},
		}
	} else {
		batchSigner = types.BatchSigner{
			BatchSigner: types.BatchSignerData{
				Account:       types.Address(batchAccount),
				SigningPubKey: wallet.PublicKey,
				TxnSignature:  signature,
			},
		}
	}

	tx.BatchSigners = []types.BatchSigner{batchSigner}

	return nil
}

// CombineBatchSigners combines the batch signers of a set of transactions into a single transaction.
// It takes a slice of transactions and returns a single transaction with the combined batch signers.
// It returns an error if the transactions are invalid.
func CombineBatchSigners(transactions []transaction.Batch) (string, error) {
	if len(transactions) == 0 {
		return "", ErrNoTransactionsProvided
	}

	var prevBatchSignable *wallettypes.BatchSignable

	signers := []types.BatchSigner{}

	for index, tx := range transactions {
		if len(tx.BatchSigners) == 0 {
			return "", ErrTxMustIncludeBatchSigner
		}

		if tx.TxnSignature != "" || len(tx.Signers) > 0 {
			return "", ErrTransactionAlreadySigned
		}

		batchSignable, err := wallettypes.FromBatchTransaction(&tx)
		if err != nil {
			return "", err
		}

		if index == 0 {
			prevBatchSignable = batchSignable
		} else if !prevBatchSignable.Equals(batchSignable) {
			return "", ErrBatchSignableNotEqual
		}

		// Add signers from this transaction, excluding the batch submitter
		for _, signer := range tx.BatchSigners {
			if signer.BatchSigner.Account != transactions[0].Account {
				signers = append(signers, signer)
			}
		}
	}

	slices.SortFunc(signers, func(a, b types.BatchSigner) int {
		return cmp.Compare(a.BatchSigner.Account, b.BatchSigner.Account)
	})

	tx := transactions[0]
	tx.BatchSigners = signers

	return binarycodec.Encode(tx.Flatten())
}
