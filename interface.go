package pfcp

import (
	"io"
)

type Writer interface {
	io.Writer
	MWrite(...interface{})
	TestWrite(test bool, v ...interface{})
	Error(info interface{}) error
}

type Reader interface {
	io.Reader
	TestRead(test bool, l int, v interface{})
	ReadUint8() uint8
	ReadBool() bool
	ReadMBool() (bool, bool, bool, bool, bool, bool, bool, bool)
	ReadUint16() uint16
	ReadUint24() uint32
	ReadUint32() uint32
	ReadUint40() uint64
	ReadUint48() uint64
	ReadUint56() uint64
	ReadUint64() uint64
	ReadBytes(int) []byte
	Error(info interface{}) error
}

type encoder interface {
	encode(w Writer) error
}

type decoder interface {
	decode(r Reader) error
}

type length interface {
	length() uint16
}

type coder interface {
	encoder
	decoder
	length
}

type header interface {
	coder
	setType(uint16)   header
	setLength(uint16) header
}

type GroupIE interface {
	IE
}

type NonGroupIE interface {
	IE
	coder
}

type IE interface {
	Type() IEType
}

type Message interface {
	Type() MessageType
}
