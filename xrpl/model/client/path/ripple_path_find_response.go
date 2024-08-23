package path

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type RipplePathFindResponse struct {
	Alternatives          []PathAlternative `json:"alternatives"`
	DestinationAccount    types.Address     `json:"destination_account"`
	DestinationCurrencies []string          `json:"destination_currencies"`
}
