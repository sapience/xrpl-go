package definitions

type FieldInstance struct {
	FieldName string
	*FieldInfo
	FieldHeader *FieldHeader
	Ordinal     int32
}

type FieldInfo struct {
	Nth            int32
	IsVLEncoded    bool
	IsSerialized   bool
	IsSigningField bool
	Type           string
}

type FieldHeader struct {
	TypeCode  int32
	FieldCode int32
}

func CreateFieldHeader(tc, fc int32) FieldHeader {
	return FieldHeader{
		TypeCode:  tc,
		FieldCode: fc,
	}
}
