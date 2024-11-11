package transaction

import (
	"testing"
)

func TestPaymentChannelClaimFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*PaymentChannelClaim)
		expected uint32
	}{
		{
			name: "SetRenewFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetRenewFlag()
			},
			expected: tfRenew,
		},
		{
			name: "SetCloseFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetCloseFlag()
			},
			expected: tfClose,
		},
		{
			name: "SetRenewFlag and SetCloseFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetRenewFlag()
				p.SetCloseFlag()
			},
			expected: tfRenew | tfClose,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaymentChannelClaim{}
			tt.setter(p)
			if p.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, p.Flags)
			}
		})
	}
}
