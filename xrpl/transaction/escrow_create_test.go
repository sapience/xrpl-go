package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscrowCreate_TxType(t *testing.T) {
	entry := &EscrowCreate{}
	assert.Equal(t, EscrowCreateTx, entry.TxType())
}
