package wallet

import (
	"errors"

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

	var batchSigner transaction.BatchSigner

	if multisignAddress != "" {
		batchSigner = transaction.BatchSigner{
			BatchSigner: transaction.BatchSignerData{
				Account: types.Address(batchAccount),
				Signers: []transaction.Signer{
					{
						SignerData: transaction.SignerData{
							Account:       types.Address(multisignAddress),
							SigningPubKey: wallet.PublicKey,
							TxnSignature:  signature,
						},
					},
				},
			},
		}
	} else {
		batchSigner = transaction.BatchSigner{
			BatchSigner: transaction.BatchSignerData{
				Account:       types.Address(batchAccount),
				SigningPubKey: wallet.PublicKey,
				TxnSignature:  signature,
			},
		}
	}

	tx.BatchSigners = []transaction.BatchSigner{batchSigner}

	return nil
}
