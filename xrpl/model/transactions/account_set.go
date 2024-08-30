package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

const (
	// The same as SetFlag: asfRequireDest.
	tfRequireDestTag  uint = 65536
	// The same as ClearFlag: asfRequireDestTag.
	tfOptionalDestTag uint = 131072
	// The same as SetFlag: asfRequireAuth.
	tfRequireAuth uint = 262144
	// The same as ClearFlag: asfRequireAuth.
	tfOptionalAuth uint = 524288
	// The same as SetFlag: asfDisallowXRP.
	tfDisallowXRP     uint = 1048576
	// The same as ClearFlag: asfDisallowXRP.
	tfAllowXRP uint = 2097152
)

// An AccountSet transaction modifies the properties of an account in the XRP
// Ledger.
type AccountSet struct {
	BaseTx
	// ClearFlag: asfRequireDestTag, asfOptionalDestTag, asfRequireAuth, asfOptionalAuth, asfDisallowXRP, asfAllowXRP
	ClearFlag     uint          `json:",omitempty"`
	// The domain that owns this account, as a string of hex representing the.
	// ASCII for the domain in lowercase.
	Domain        string        `json:",omitempty"`
	// Hash of an email address to be used for generating an avatar image.
	EmailHash     types.Hash128 `json:",omitempty"`
	//Public key for sending encrypted messages to this account.
	MessageKey    string        `json:",omitempty"`
	// Sets an alternate account that is allowed to mint NFTokens on this
	// account's behalf using NFTokenMint's `Issuer` field.
	NFTokenMinter string        `json:",omitempty"`
	// Integer flag to enable for this account.
	SetFlag       uint          `json:",omitempty"`
	// The fee to charge when users transfer this account's issued currencies,
	// represented as billionths of a unit. Cannot be more than 2000000000 or less
	// than 1000000000, except for the special case 0 meaning no fee.
	TransferRate  uint          `json:",omitempty"`
	// Tick size to use for offers involving a currency issued by this address.
	// The exchange rates of those offers is rounded to this many significant
	// digits. Valid values are 3 to 15 inclusive, or 0 to disable.
	TickSize      uint8         `json:",omitempty"`
	WalletLocator types.Hash256 `json:",omitempty"`
	WalletSize    uint          `json:",omitempty"`
}

// TxType returns the type of the transaction (AccountSet).
func (*AccountSet) TxType() TxType {
	return AccountSetTx
}

// Flatten returns the flattened map of the AccountSet transaction.
func (s *AccountSet) Flatten() map[string]interface{} {
	flattened := s.BaseTx.Flatten()

	if s.ClearFlag != 0 {
		flattened["ClearFlag"] = s.ClearFlag
	}
	if s.Domain != "" {
		flattened["Domain"] = s.Domain
	}
	if s.EmailHash != "" {
		flattened["EmailHash"] = s.EmailHash.String()
	}
	if s.MessageKey != "" {
		flattened["MessageKey"] = s.MessageKey
	}
	if s.NFTokenMinter != "" {
		flattened["NFTokenMinter"] = s.NFTokenMinter
	}
	if s.SetFlag != 0 {
		flattened["SetFlag"] = s.SetFlag
	}
	if s.TransferRate != 0 {
		flattened["TransferRate"] = s.TransferRate
	}
	if s.TickSize != 0 {
		flattened["TickSize"] = s.TickSize
	}
	if s.WalletLocator != "" {
		flattened["WalletLocator"] = s.WalletLocator.String()
	}
	if s.WalletSize != 0 {
		flattened["WalletSize"] = s.WalletSize
	}

	return flattened
}

// SetRequireDestTag sets the require destination tag flag.
// If enable is true, the require destination tag flag is set.
// This flag is disabled by default.
func (s *AccountSet) SetRequireDestTag(enable bool) {
	if enable {
		s.Flags |= tfRequireDestTag
	} else {
		s.Flags |= tfRequireDestTag
	}
}

// SetRequireAuth sets the require auth flag.
// If enable is true, the require auth flag is set.
// This flag is disabled by default.
func (s *AccountSet) SetRequireAuth(enable bool) {
	if enable {
		s.Flags |= tfRequireAuth
	} else {
		s.Flags &= ^tfRequireAuth
	}
}

// SetDisallowXRP sets the disallow XRP flag.
// If enable is true, the disallow XRP flag is set.
// This flag is disabled by default.
func (s *AccountSet) SetDisallowXRP(enable bool) {
	if enable {
		s.Flags |= tfDisallowXRP
	} else {
		s.Flags &= ^tfDisallowXRP
	}
}

// SetOptionalDestTag sets the optional destination tag flag.
// If enable is true, the optional destination tag flag is set.
// This flag is disabled by default.
func (s *AccountSet) SetOptionalDestTag(enable bool) {
	if enable {
		s.Flags |= tfOptionalDestTag
	} else {
		s.Flags &= ^tfOptionalDestTag
	}
}

// SetOptionalAuth sets the optional auth flag.
// If enable is true, the optional auth flag is set.
// This flag is disabled by default.
func (s *AccountSet) SetOptionalAuth(enable bool) {
	if enable {
		s.Flags |= tfOptionalAuth
	} else {
		s.Flags &= ^tfOptionalAuth
	}
}

// SetAllowXRP sets the allow XRP flag.
// If enable is true, the allow XRP flag is set.
// This flag is disabled by default.
func (s *AccountSet) SetAllowXRP(enable bool) {
	if enable {
		s.Flags |= tfAllowXRP
	} else {
		s.Flags &= ^tfAllowXRP
	}
}

