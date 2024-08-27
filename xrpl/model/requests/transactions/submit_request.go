package transactions

import "errors"

type SubmitRequest struct {
	TxBlob   string `json:"tx_blob"`
	FailHard bool   `json:"fail_hard,omitempty"`
}

func (*SubmitRequest) Method() string {
	return "submit"
}

func (req *SubmitRequest) Validate() error {
	if req.TxBlob == "" {
		return errors.New("no TxBlob defined")
	}
	return nil
}
