package status

type FetchInfoRequest struct {
	Clear bool `json:"clear"`
}

func (*FetchInfoRequest) Method() string {
	return "fetch_info"
}

type FetchInfoResponse struct {
	Info map[string]FetchInfo `json:"info"`
}

type FetchInfo struct {
	Hash              string   `json:"hash"`
	HaveHeader        bool     `json:"have_header"`
	HaveTransactions  bool     `json:"have_transactions"`
	NeededStateHashes []string `json:"needed_state_hashes"`
	Peers             int      `json:"peers"`
	Timeouts          int      `json:"timeouts"`
}
