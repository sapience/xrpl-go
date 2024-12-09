package utility

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type RandomRequest struct{}

func (*RandomRequest) Method() string {
	return "random"
}

type RandomResponse struct {
	Random types.Hash256 `json:"random"`
}
