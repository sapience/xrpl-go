package signing

import "github.com/Peersyst/xrpl-go/xrpl/transaction"

type SignForResponse struct {
	TxBlob string                       `json:"tx_blob"`
	TxJson transaction.FlatTransaction `json:"tx_json"`
}
