package websocket

import (
	"fmt"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	transaction "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

// getSignedTx returns a signed transaction blob.
// It accepts either a pre-encoded transaction (string)
// or a SubmittableTransaction that will be autofilled and signed.
// It ensures the returned blob contains a signature.
func getSignedTx(client *Client, tx transaction.FlatTransaction, autofill bool, wallet *wallet.Wallet) (string, error) {
		// Check if the transaction is already signed.
		_, hasSig := tx["TxSignature"].(string)
		// _, hasTxnSig := tx["TxnSignature"].(string)
		_, hasPubKey := tx["SigningPubKey"].(string)
		if hasSig ||  hasPubKey {
			// Encode and return.
			blob, err := binarycodec.Encode(tx)
			if err != nil {
				return "", err
			}
			return blob, nil
		}

		// Autofill if required.
		if autofill {
			if err := client.Autofill(&tx); err != nil {
				return "", err
			}
		}

		// Ensure a wallet is provided for signing.
		if wallet == nil {
			return "", fmt.Errorf("wallet must be provided when submitting an unsigned transaction")
		}

		// Sign the transaction.
		txBlob, _, err := wallet.Sign(tx)
		if err != nil {
			return "", err
		}

		// Validate that the signed blob contains a signature.
		decoded, err := binarycodec.Decode(txBlob)
		if err != nil {
			return "", err
		}
		_, hasSig = decoded["TxSignature"].(string)
		_, hasPubKey = decoded["SigningPubKey"].(string)
		if !hasSig  && !hasPubKey {
			return "", ErrMissingTxSignatureOrSigningPubKey
		}
		return txBlob, nil
}