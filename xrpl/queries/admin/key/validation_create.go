package key

type ValidationCreateRequest struct {
	Secret string `json:"secret,omitempty"`
}

func (*ValidationCreateRequest) Method() string {
	return "validation_create"
}

type ValidationCreateResponse struct {
	ValidationKey       string `json:"validation_key"`
	ValidationPublicKey string `json:"validation_public_key"`
	ValidationSeed      string `json:"validation_seed"`
}
