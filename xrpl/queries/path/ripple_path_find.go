package path

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	pathtypes "github.com/Peersyst/xrpl-go/xrpl/queries/path/types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type RipplePathFindRequest struct {
	SourceAccount      types.Address                `json:"source_account"`
	DestinationAccount types.Address                `json:"destination_account"`
	DestinationAmount  types.CurrencyAmount         `json:"destination_amount"`
	SendMax            types.CurrencyAmount         `json:"send_max,omitempty"`
	SourceCurrencies   []types.IssuedCurrencyAmount `json:"source_currencies,omitempty"`
	LedgerHash         common.LedgerHash            `json:"ledger_hash,omitempty"`
	LedgerIndex        common.LedgerSpecifier       `json:"ledger_index,omitempty"`
}

func (*RipplePathFindRequest) Method() string {
	return "ripple_path_find"
}

type RipplePathFindResponse struct {
	Alternatives          []pathtypes.Alternative `json:"alternatives"`
	DestinationAccount    types.Address `json:"destination_account"`
	DestinationCurrencies []string      `json:"destination_currencies"`
	FullReply             bool          `json:"full_reply,omitempty"`
	LedgerCurrentIndex    int           `json:"ledger_current_index,omitempty"`
	SourceAccount         types.Address `json:"source_account"`
	Validated             bool          `json:"validated"`
}
