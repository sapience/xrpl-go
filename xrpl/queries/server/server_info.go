package server

import (
	servertypes "github.com/Peersyst/xrpl-go/xrpl/queries/server/types"
)

// ############################################################################
// Request
// ############################################################################

type InfoRequest struct {
}

func (*InfoRequest) Method() string {
	return "server_info"
}

// TODO: Implement V2
func (*InfoRequest) Validate() error {
	return nil
}

// ############################################################################
// Request
// ############################################################################

type InfoResponse struct {
	Info servertypes.Info `json:"info"`
}
