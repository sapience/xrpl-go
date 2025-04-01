package websocket

import "errors"

// Static errors
var (
	ErrMissingTxSignatureOrSigningPubKey      = errors.New("transaction must have a TxSignature or SigningPubKey set")
	ErrMissingLastLedgerSequenceInTransaction = errors.New("missing LastLedgerSequence in transaction")
	ErrMissingWallet                          = errors.New("wallet must be provided when submitting an unsigned transaction")
)

// Dynamic errors

type ClientError struct {
	ErrorString string
}

func (e *ClientError) Error() string {
	return e.ErrorString
}
