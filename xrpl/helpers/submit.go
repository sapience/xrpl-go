package helpers

import (
	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/xrpl/client"
	requests "github.com/Peersyst/xrpl-go/xrpl/model/requests/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
)

// SubmitTransaction submits a transaction object to the XRP Ledger.
func SubmitTransaction(tx interface{}, c *client.XRPLClient, failHard bool) (*requests.SubmitRequest, client.XRPLResponse, error) {
	if !isSigned(tx) {
		panic("Transaction must be signed")
	}

	// check if tx is a string
	_, isStr := tx.(string)
	if isStr {
		return submitRequest(tx.(string), c, failHard)
	}

	// else if it is an object, encode it to a string
	txBlob, err := encodeTransaction(tx)
	if err != nil {
		panic("Could not encode transaction")
	}

	return submitRequest(txBlob, c, failHard)
}

func encodeTransaction(tx interface{}) (string, error) {
	obj := StructToMap(tx)
	txBlob, err := binarycodec.Encode(obj)
	if err != nil {
		return "", err
	}
	return txBlob, nil
}

// submitRequest submits a transaction blob to the XRP Ledger.
func submitRequest(txBlob string, c *client.XRPLClient, failHard bool) (*requests.SubmitRequest, client.XRPLResponse, error) {
	request := requests.SubmitRequest{
		TxBlob:   txBlob,
		FailHard: failHard,
	}

	response, err := c.Client().SendRequest(&request)
	if err != nil {
		return nil, nil, err
	}
	return &request, response, nil
}

// isSigned checks whether a transaction is signed.
func isSigned(tx interface{}) bool {
	// switch statement to handle different types of transactions
	switch tx := tx.(type) {
	case transactions.BaseTx:
		// check if Signers field is not nil and if all signers have non-empty SigningPubKey and TxnSignature fields
		if tx.Signers != nil {
			for _, signer := range tx.Signers {
				if signer.SignerData.SigningPubKey == "" || signer.SignerData.TxnSignature == "" {
					return false
				}
			}
			return true
		}
		// check if SigningPubKey and TxnSignature fields are non-empty
		return tx.SigningPubKey != "" && tx.TxnSignature != ""

	case string:
		// attempt to decode the transaction using binarycodec.Decode
		_, err := binarycodec.Decode(tx)
		if err != nil {
			return false
		}
		return true

	default:
		return false
	}
}
