package hash

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"strings"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
)

var (
	ErrMissingSignature = errors.New("transaction must have at least one of TxnSignature, Signers, or SigningPubKey")
)

// SignTxBlob hashes a signed transaction blob
// It takes a transaction blob and returns the hash of the signed transaction.
// It returns an error if the transaction blob is invalid.
func SignTxBlob(txBlob string) (string, error) {
	tx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return "", err
	}

	// Check if this is an inner batch transaction (has TfInnerBatchTxn flag)
	isInnerBatchTxn := false
	if flags, ok := tx["Flags"].(uint32); ok {
		isInnerBatchTxn = (flags & transaction.TfInnerBatchTxn) != 0
	}

	// Check if the transaction has at least one of the required signature fields
	hasTxnSignature := tx["TxnSignature"] != nil
	hasSigners := tx["Signers"] != nil
	hasSigningPubKey := tx["SigningPubKey"] != nil

	// For non-inner batch transactions, require at least one signature field
	if !hasTxnSignature && !hasSigners && !hasSigningPubKey && !isInnerBatchTxn {
		return "", ErrMissingSignature
	}

	// Create a byte slice with the correct capacity
	payload := make([]byte, 4+len(txBlob)/2)

	// Convert TRANSACTION_PREFIX to big-endian bytes
	binary.BigEndian.PutUint32(payload[:4], TransactionPrefix)

	// Decode the txBlob into the rest of the payload
	_, err = hex.Decode(payload[4:], []byte(txBlob))
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(crypto.Sha512Half(payload))), nil
}
