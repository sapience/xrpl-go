package ledger

type NegativeUNL struct {
	DisabledValidators  []DisabledValidatorEntry `json:",omitempty"`
	Flags               uint
	LedgerEntryType     EntryType
	ValidatorToDisable  string `json:",omitempty"`
	ValidatorToReEnable string `json:",omitempty"`
}

func (*NegativeUNL) EntryType() EntryType {
	return NegativeUNLEntry
}

type DisabledValidatorEntry struct {
	DisabledValidator DisabledValidator
}

type DisabledValidator struct {
	FirstLedgerSequence uint
	PublicKey           string
}
