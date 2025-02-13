package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferRate(t *testing.T) {
	tests := []struct {
		name  string
		value uint32
		want  *uint32
	}{
		{
			name:  "Valid transfer rate",
			value: 1500000000,
			want:  func() *uint32 { v := uint32(1500000000); return &v }(),
		},
		{
			name:  "Minimum transfer rate",
			value: 1000000000,
			want:  func() *uint32 { v := uint32(1000000000); return &v }(),
		},
		{
			name:  "Maximum transfer rate",
			value: 2000000000,
			want:  func() *uint32 { v := uint32(2000000000); return &v }(),
		},
		{
			name:  "No fee transfer rate",
			value: 0,
			want:  func() *uint32 { v := uint32(0); return &v }(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TransferRate(tt.value)
			require.Equal(t, tt.want, result)
		})
	}
}
