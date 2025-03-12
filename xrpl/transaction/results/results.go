package results

// TxResult represents the result code of a transaction
type TxResult string

const (
	// Success results - https://xrpl.org/docs/references/protocol/transactions/transaction-results/tes-success

	// The transaction was applied and forwarded to other servers.
	// If this appears in a validated ledger, then the transaction's success is final.
	TesSUCCESS TxResult = "tesSUCCESS"
)

// String returns the string representation of the result
func (t TxResult) String() string {
	return string(t)
}
