package serdes

import (
	"github.com/Peersyst/xrpl-go/binary-codec/definitions"
)

// Returns the unique field ID for a given field name.
// This field ID consists of the type code and field code, in 1 to 3 bytes
// depending on whether those values are "common" (<16) or "uncommon" (>16).
func encodeFieldID(fieldName string) ([]byte, error) {
	fh, err := definitions.Get().GetFieldHeaderByFieldName(fieldName)
	if err != nil {
		return nil, err
	}
	var b []byte
	if fh.TypeCode < 16 && fh.FieldCode < 16 {
		return append(b, (byte(fh.TypeCode<<4))|byte(fh.FieldCode)), nil
	}
	if fh.TypeCode >= 16 && fh.FieldCode < 16 {
		return append(b, (byte(fh.FieldCode)), byte(fh.TypeCode)), nil
	}
	if fh.TypeCode < 16 && fh.FieldCode >= 16 {
		return append(b, byte(fh.TypeCode<<4), byte(fh.FieldCode)), nil
	}
	if fh.TypeCode >= 16 && fh.FieldCode >= 16 {
		return append(b, 0, byte(fh.TypeCode), byte(fh.FieldCode)), nil
	}
	return nil, nil
}
