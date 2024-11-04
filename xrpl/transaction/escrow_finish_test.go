package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscrowFinish_TxType(t *testing.T) {
	entry := &EscrowFinish{}
	assert.Equal(t, EscrowFinishTx, entry.TxType())
}
