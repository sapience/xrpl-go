package definitions

import "github.com/ugorji/go/codec"

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

func (d *Definitions) CreateFieldHeader(tc, fc int32) FieldHeader {
	return FieldHeader{
		TypeCode:  tc,
		FieldCode: fc,
	}
}

type fieldInstanceMap map[string]*FieldInstance

// CodecEncodeSelf implements the codec.SelfEncoder interface.
func (fi *fieldInstanceMap) CodecEncodeSelf(_ *codec.Encoder) {}

// CodecDecodeSelf implements the codec.SelfDecoder interface.
func (fi *fieldInstanceMap) CodecDecodeSelf(d *codec.Decoder) {
	var x [][]interface{}
	d.MustDecode(&x)
	y := convertToFieldInstanceMap(x)
	*fi = y
}
