package types

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"github.com/Peersyst/xrpl-go/binary-codec/definitions"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

var (
	ErrInvalidJSONNumber         = errors.New("invalid json.Number")
	ErrUnsupportedPermissionType = errors.New("unsupported JSON type for PermissionValue")
)

// PermissionValue represents a 32-bit unsigned integer permission value.
type PermissionValue struct{}

// FromJSON converts a JSON value into a serialized byte slice representing a 32-bit unsigned integer permission value.
// If the input value is a string, it's assumed to be a permission name, and the method will
// attempt to convert it into a corresponding permission value. If the conversion fails, an error is returned.
func (p *PermissionValue) FromJSON(value any) ([]byte, error) {
	if s, ok := value.(string); ok {
		pv, err := definitions.Get().GetDelegatablePermissionValueByName(s)
		if err != nil {
			return nil, err
		}
		value = pv
	}

	var intValue uint32

	switch v := value.(type) {
	case int:
		intValue = uint32(v)
	case int32:
		intValue = uint32(v)
	case int64:
		intValue = uint32(v)
	case uint32:
		intValue = v
	case float64:
		intValue = uint32(v)
	case json.Number:
		num, err := v.Int64()
		if err != nil {
			return nil, ErrInvalidJSONNumber
		}
		intValue = uint32(num)
	default:
		return nil, ErrUnsupportedPermissionType
	}

	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, intValue)
	return buf, nil
}

// ToJSON takes a BinaryParser and optional parameters, and converts the serialized byte data
// back into a JSON value. If a permission name is found for the value, it returns the name;
// otherwise, it returns the numeric value. If the parsing fails, an error is returned.
func (p *PermissionValue) ToJSON(parser interfaces.BinaryParser, _ ...int) (any, error) {
	b, err := parser.ReadBytes(4)
	if err != nil {
		return nil, err
	}

	permissionValue := binary.BigEndian.Uint32(b)

	if name, err := definitions.Get().GetDelegatablePermissionNameByValue(int32(permissionValue)); err == nil {
		return name, nil
	}

	return permissionValue, nil
}
