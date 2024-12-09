package status

type FeatureRequest struct {
	Feature string `json:"feature,omitempty"`
	Vetoed  bool   `json:"vetoed,omitempty"`
}

func (*FeatureRequest) Method() string {
	return "feature"
}

// TODO support deserialization of RPC and Websocket
// Currently only supports websocket
type FeatureResponse struct {
	Features map[string]Feature `json:"features"`
}

type Feature struct {
	Enabled   bool   `json:"enabled"`
	Name      string `json:"name,omitempty"`
	Supported bool   `json:"supported"`
	Vetoed    bool   `json:"vetoed"`
}
