package server

import (
	servertypes "github.com/Peersyst/xrpl-go/xrpl/queries/server/types"
)

// ############################################################################
// Request
// ############################################################################

type StateRequest struct {
}

func (*StateRequest) Method() string {
	return "server_state"
}

// TODO: Implement V2
func (*StateRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type StateResponse struct {
	State servertypes.State `json:"state"`
}
