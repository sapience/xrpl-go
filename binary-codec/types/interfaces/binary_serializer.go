package interfaces

import "github.com/Peersyst/xrpl-go/binary-codec/definitions"

type BinarySerializer interface {
	WriteFieldAndValue(fieldInstance definitions.FieldInstance, value []byte) error
	GetSink() []byte
}
