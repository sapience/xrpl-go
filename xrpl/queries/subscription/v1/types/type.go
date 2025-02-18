package types

type Type string

const (
	LedgerStreamType      Type = "ledgerClosed"
	ValidationStreamType  Type = "validationReceived"
	TransactionStreamType Type = "transaction"
	PeerStatusStreamType  Type = "peerStatusChange"
	// TODO example lists OrderBookStreamType as "transaction"
	OrderBookStreamType Type = TransactionStreamType
	ConsensusStreamType Type = "consensusPhase"
)
