package websocket

import (
	"fmt"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/xrpl/common"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

// getSignedTx returns a signed transaction blob.
// It accepts either a pre-encoded transaction (string)
// or a SubmittableTransaction that will be autofilled and signed.
// It ensures the returned blob contains a signature.
func getSignedTx(client *Client, txInput interface{}, autofill bool, wallet *wallet.Wallet) (string, error) {
	switch tx := txInput.(type) {
	case string:
		// Assume it's already an encoded transaction.
		decoded, err := binarycodec.Decode(tx)
		if err != nil {
			return "", err
		}
		// Check for signature in either field.
		_, hasSig := decoded["TxSignature"].(string)
		// _, hasTxnSig := decoded["TxnSignature"].(string)
		_, hasPubKey := decoded["SigningPubKey"].(string)
		if hasSig ||  hasPubKey {
			return tx, nil
		}
		// Otherwise, we cannot sign a string transaction.
		return "", fmt.Errorf("provided string transaction is not signed")

	case common.SubmittableTransaction:
		// Flatten the transaction.
		flatTx := tx.Flatten()

			// Check if the transaction is already signed.
		_, hasSig := flatTx["TxSignature"].(string)
		// _, hasTxnSig := flatTx["TxnSignature"].(string)
		_, hasPubKey := flatTx["SigningPubKey"].(string)
		if hasSig ||  hasPubKey {
			// Encode and return.
			blob, err := binarycodec.Encode(flatTx)
			if err != nil {
				return "", err
			}
			return blob, nil
		}

		// Autofill if required.
		if autofill {
			if err := client.Autofill(&flatTx); err != nil {
				return "", err
			}
		}

		// Ensure a wallet is provided for signing.
		if wallet == nil {
			return "", fmt.Errorf("wallet must be provided when submitting an unsigned transaction")
		}

		// Sign the transaction.
		txBlob, _, err := wallet.Sign(flatTx)
		if err != nil {
			return "", err
		}

		// Validate that the signed blob contains a signature.
		decoded, err := binarycodec.Decode(txBlob)
		if err != nil {
			return "", err
		}
		_, hasSig = decoded["TxSignature"].(string)
		// _, hasTxnSig = decoded["TxnSignature"].(string)
		_, hasPubKey = decoded["SigningPubKey"].(string)
		if !hasSig  && !hasPubKey {
			return "", ErrMissingTxSignatureOrSigningPubKey
		}
		return txBlob, nil


	default:
		return "", fmt.Errorf("invalid transaction type; expected string or SubmittableTransaction")
	}
}