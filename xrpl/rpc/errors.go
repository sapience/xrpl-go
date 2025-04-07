package rpc

import "errors"

// Static errors
var (
	ErrIncorrectID                            = errors.New("incorrect id")
	ErrMissingTxSignatureOrSigningPubKey      = errors.New("transaction must have a TxSignature or SigningPubKey set")
	ErrSignerDataIsEmpty                      = errors.New("signer data is empty")
	ErrCannotFundWalletWithoutClassicAddress  = errors.New("cannot fund wallet without classic address")
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
