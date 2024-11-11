package stream

type ConsensusStream struct {
	Type      Type   `json:"type"`
	Consensus string `json:"consensus"`
}
