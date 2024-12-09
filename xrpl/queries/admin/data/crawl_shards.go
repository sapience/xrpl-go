package data

type CrawlShardsRequest struct {
	PublicKey bool `json:"public_key,omitempty"`
	Limit     int  `json:"limit,omitempty"`
}

func (*CrawlShardsRequest) Method() string {
	return "crawl_shards"
}

type CrawlShardsResponse struct {
	CompleteShards string       `json:"complete_shards,omitempty"`
	Peers          []PeerShards `json:"peers,omitempty"`
}

type PeerShards struct {
	CompleteShards   string `json:"complete_shards"`
	IncompleteShards string `json:"incomplete_shards,omitempty"`
	PublicKey        string `json:"public_key,omitempty"`
}
