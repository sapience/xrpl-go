package path

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type RipplePathFindResponse struct {
	Alternatives          []Alternative `json:"alternatives"`
	DestinationAccount    types.Address `json:"destination_account"`
	DestinationCurrencies []string      `json:"destination_currencies"`
}
