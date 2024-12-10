package server

type ManifestDetails struct {
	Domain       string `json:"domain"`
	EphemeralKey string `json:"ephemeral_key"`
	MasterKey    string `json:"master_key"`
	Seq          uint   `json:"seq"`
}

// ############################################################################
// Request
// ############################################################################

type ManifestRequest struct {
	PublicKey string `json:"public_key"`
}

func (*ManifestRequest) Method() string {
	return "manifest"
}

// TODO: Implement V2
func (*ManifestRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type ManifestResponse struct {
	Details   ManifestDetails `json:"details,omitempty"`
	Manifest  string          `json:"manifest,omitempty"`
	Requested string          `json:"requested"`
}
