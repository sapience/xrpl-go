package data

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
)

type CanDeleteRequest struct {
	CanDelete common.LedgerSpecifier `json:"can_delete,omitempty"`
}

func (*CanDeleteRequest) Method() string {
	return "can_delete"
}

func (r *CanDeleteRequest) UnmarshalJSON(data []byte) error {
	type cdHelper struct {
		CanDelete json.RawMessage `json:"can_delete"`
	}
	var h cdHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	r.CanDelete, err = common.UnmarshalLedgerSpecifier(h.CanDelete)

	return err
}
