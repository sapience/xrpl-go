package types

// Permission represents a transaction permission that can be delegated to another account.
// This matches the xrpl.js Permission interface structure.
type Permission struct {
	Permission PermissionValue `json:"Permission"`
}

// PermissionValue represents the inner permission value structure.
type PermissionValue struct {
	PermissionValue string `json:"PermissionValue"`
}

// Flatten returns the flattened map representation of the Permission.
func (p *Permission) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})
	flattened["Permission"] = p.Permission.Flatten()
	return flattened
}

// Flatten returns the flattened map representation of the PermissionValue.
func (pv *PermissionValue) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})
	if pv.PermissionValue != "" {
		flattened["PermissionValue"] = pv.PermissionValue
	}
	return flattened
}
