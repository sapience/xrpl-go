package websocket

import (
	"fmt"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	transaction "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

// getSignedTx returns a fully signed transaction blob. It checks whether the
// transaction already contains a signature or public key and, if not, optionally
// autofills missing fields and signs the transaction using the provided wallet.
// It verifies the resulting blob includes a signature, returning an error if absent.
func getSignedTx(client *Client, tx transaction.FlatTransaction, autofill bool, wallet *wallet.Wallet) (string, error) {
	_, hasSig := tx["TxSignature"].(string)
	_, hasPubKey := tx["SigningPubKey"].(string)
	if hasSig || hasPubKey {

		blob, err := binarycodec.Encode(tx)
		if err != nil {
			return "", err
		}
		return blob, nil
	}
	if autofill {
		if err := client.Autofill(&tx); err != nil {
			return "", err
		}
	}
	if wallet == nil {
		return "", fmt.Errorf("wallet must be provided when submitting an unsigned transaction")
	}

	txBlob, _, err := wallet.Sign(tx)
	if err != nil {
		return "", err
	}

	decoded, err := binarycodec.Decode(txBlob)
	if err != nil {
		return "", err
	}
	_, hasSig = decoded["TxSignature"].(string)
	_, hasPubKey = decoded["SigningPubKey"].(string)
	if !hasSig && !hasPubKey {
		return "", ErrMissingTxSignatureOrSigningPubKey
	}
	return txBlob, nil
}
