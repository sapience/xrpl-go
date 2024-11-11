package ledger

type FeeSettings struct {
	BaseFee           string
	Flags             uint32
	LedgerEntryType   EntryType
	ReferenceFeeUnits uint32
	ReserveBase       uint32
	ReserveIncrement  uint32
}

func (*FeeSettings) EntryType() EntryType {
	return FeeSettingsEntry
}
