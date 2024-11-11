package ledger

type FeeSettings struct {
	BaseFee           string
	Flags             uint
	LedgerEntryType   EntryType
	ReferenceFeeUnits uint
	ReserveBase       uint
	ReserveIncrement  uint
}

func (*FeeSettings) EntryType() EntryType {
	return FeeSettingsEntry
}
