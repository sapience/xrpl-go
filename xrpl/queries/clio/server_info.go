package clio

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type ServerInfoRequest struct {
}

func (*ServerInfoRequest) Method() string {
	return "server_info"
}

type ServerInfoResponse struct {
	Info      ServerInfo `json:"info"`
	Validated bool       `json:"validated"`
	Status    string     `json:"status,omitempty"`
}

type ServerInfo struct {
	CompleteLedgers      string      `json:"complete_ledgers"`
	Counters             *Counters   `json:"counters,omitempty"`
	LoadFactor           int         `json:"load_factor"`
	ClioVersion          string      `json:"clio_version"`
	ValidationQuorum     int         `json:"validation_quorum,omitempty"`
	RippledVersion       string      `json:"rippled_version,omitempty"`
	ValidatedLedger      *LedgerInfo `json:"validated_ledger,omitempty"`
	ValidatorListExpires string      `json:"validator_list_expires,omitempty"`
	Cache                Cache       `json:"cache"`
	ETL                  *ETL        `json:"etl,omitempty"`
}

type Counters struct {
	RPC           map[string]RPC `json:"rpc"`
	Subscriptions Subscriptions  `json:"subscriptions"`
}

type RPC struct {
	Started    string `json:"started,omitempty"`
	Finished   string `json:"finished,omitempty"`
	Errored    string `json:"errored,omitempty"`
	Forwarded  string `json:"forwarded,omitempty"`
	DurationUS string `json:"duration_us,omitempty"`
}

type Subscriptions struct {
	Ledger               int `json:"ledger"`
	Transactions         int `json:"transactions"`
	TransactionsProposed int `json:"transactions_proposed"`
	Manifests            int `json:"manifests"`
	Validations          int `json:"validations"`
	Account              int `json:"account"`
	AccountsProposed     int `json:"accounts_proposed"`
	Books                int `json:"books"`
}

type LedgerInfo struct {
	Age            uint          `json:"age"`
	BaseFeeXRP     float32       `json:"base_fee_xrp"`
	Hash           types.Hash256 `json:"hash"`
	ReserveBaseXRP float32       `json:"reserve_base_xrp"`
	ReserveIncXRP  float32       `json:"reserve_inc_xrp"`
	Seq            uint          `json:"seq"`
}

type Cache struct {
	Size            int  `json:"size"`
	IsFull          bool `json:"is_full"`
	LatestLedgerSeq int  `json:"latest_ledger_seq"`
}

type ETL struct {
	ETLSources            []ETLSource `json:"etl_sources"`
	IsWriter              bool        `json:"is_writer"`
	ReadOnly              bool        `json:"read_only"`
	LastPublishAgeSeconds string      `json:"last_publish_age_seconds"`
}

type ETLSource struct {
	ValidatedRange    string `json:"validated_range"`
	IsConnected       string `json:"is_connected"`
	IP                string `json:"ip"`
	WSPort            string `json:"ws_port"`
	GRPCPort          string `json:"grpc_port"`
	LastMsgAgeSeconds string `json:"last_msg_age_seconds"`
}
