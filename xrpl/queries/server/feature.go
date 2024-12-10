package server

import "github.com/Peersyst/xrpl-go/xrpl/queries/server/types"

// ############################################################################
// Feature All Request
// ############################################################################

type FeatureAllRequest struct {
}

func (*FeatureAllRequest) Method() string {
	return "feature"
}

// TODO: Implement V2
func (*FeatureAllRequest) Validate() error {
	return nil
}

// ############################################################################
// Feature All Response
// ############################################################################

type FeatureAllResponse struct {
	Features map[string]types.FeatureStatus `json:"features"`
}

// ############################################################################
// Feature One Request
// ############################################################################

type FeatureOneRequest struct {
	Feature string `json:"feature"`
}

func (*FeatureOneRequest) Method() string {
	return "feature"
}

// TODO: Implement V2
func (*FeatureOneRequest) Validate() error {
	return nil
}

// ############################################################################
// Feature One Response
// ############################################################################

type FeatureResponse map[string]types.FeatureStatus
