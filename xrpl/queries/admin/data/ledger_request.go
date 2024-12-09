package data

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/ledger"
)

type LedgerRequestRequest struct {
	LedgerIndex common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash  `json:"ledger_hash,omitempty"`
}

func (*LedgerRequest) Method() string {
	return "ledger_request"
}

type LedgerRequestResponse struct {
	LedgerHeader  *ledger.Header `json:"ledger,omitempty"`
	LedgerRequest *LedgerRequest `json:"-"`
	Acquiring     *LedgerRequest `json:"acquiring,omitempty"`
}

type LedgerRequest struct {
	Hash                    common.LedgerHash `json:"hash,omitempty"`
	HaveHeader              bool              `json:"have_header"`
	HaveState               bool              `json:"have_state,omitempty"`
	HaveTransactions        bool              `json:"have_transactions,omitempty"`
	NeededStateHashes       []string          `json:"needed_state_hashes,omitempty"`
	NeededTransactionHashes []string          `json:"needed_transaction_hashes,omitempty"`
	Peers                   int               `json:"peers"`
	Timeouts                int               `json:"timeouts"`
}

func (r LedgerRequestResponse) MarshalJSON() ([]byte, error) {
	type rHelper LedgerRequestResponse
	data := make(map[string]interface{})
	h := rHelper(r)

	first, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(first, &data)
	if err != nil {
		return nil, err
	}

	if r.LedgerRequest != nil {
		second, err := json.Marshal(r.LedgerRequest)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(second, &data)
		if err != nil {
			return nil, err
		}
	}

	return json.Marshal(data)
}

func (r *LedgerRequestResponse) UnmarshalJSON(data []byte) error {
	type rHelper LedgerRequestResponse
	h := rHelper(*r)
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	r.Acquiring = h.Acquiring
	r.LedgerHeader = h.LedgerHeader
	if r.Acquiring == nil && r.LedgerHeader == nil {
		if err := json.Unmarshal(data, &(r.LedgerRequest)); err != nil {
			return err
		}
	}

	return nil
}
