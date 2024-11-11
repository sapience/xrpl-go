package signing

import "github.com/Peersyst/xrpl-go/xrpl/transaction"

type SignResponse struct {
	TxBlob string                      `json:"tx_blob"`
	TxJSON transaction.FlatTransaction `json:"tx_json"`
}
