package utility

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

// ############################################################################
// Request
// ############################################################################

type RandomRequest struct{}

func (*RandomRequest) Method() string {
	return "random"
}

func (*RandomRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type RandomResponse struct {
	Random types.Hash256 `json:"random"`
}
