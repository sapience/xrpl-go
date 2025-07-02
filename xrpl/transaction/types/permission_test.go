package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermission_Flatten(t *testing.T) {
	tests := []struct {
		name       string
		permission Permission
		expected   map[string]interface{}
	}{
		{
			name: "pass - valid permission",
			permission: Permission{
				Permission: PermissionValue{
					PermissionValue: "Payment",
				},
			},
			expected: map[string]interface{}{
				"Permission": map[string]interface{}{
					"PermissionValue": "Payment",
				},
			},
		},
		{
			name: "pass - empty permission value",
			permission: Permission{
				Permission: PermissionValue{
					PermissionValue: "",
				},
			},
			expected: map[string]interface{}{
				"Permission": map[string]interface{}{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.permission.Flatten()
			require.Equal(t, test.expected, result)
		})
	}
}

func TestPermissionValue_Flatten(t *testing.T) {
	tests := []struct {
		name            string
		permissionValue PermissionValue
		expected        map[string]interface{}
	}{
		{
			name: "pass - valid permission value",
			permissionValue: PermissionValue{
				PermissionValue: "Payment",
			},
			expected: map[string]interface{}{
				"PermissionValue": "Payment",
			},
		},
		{
			name: "pass - empty permission value",
			permissionValue: PermissionValue{
				PermissionValue: "",
			},
			expected: map[string]interface{}{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.permissionValue.Flatten()
			require.Equal(t, test.expected, result)
		})
	}
}
