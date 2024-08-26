package utility

import "github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"

type RandomResponse struct {
	Random types.Hash256 `json:"random"`
}
