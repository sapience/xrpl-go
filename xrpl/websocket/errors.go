package websocket

import "errors"

// Static errors
var (
	ErrMissingTxSignatureOrSigningPubKey = errors.New("transaction must have a TxSignature or SigningPubKey set")
)

// Dynamic errors

type ClientError struct {
	ErrorString string
}

func (e *ClientError) Error() string {
	return e.ErrorString
}
