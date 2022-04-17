package pfcp

import (
	"encoding/binary"
	dump "github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"io"
	"reflect"
)

func NewReadWriter(buffer []byte) *ReadWriter {
	return &ReadWriter{
		data: buffer,
	}
}

type ReadWriter struct {
	data  []byte
	err   error
	index int
}

func (b *ReadWriter) TestWrite(test bool, v ...interface{}) {
	if test {
		b.MWrite(v...)
	}
}

func (b *ReadWriter) TestRead(test bool, len int, v interface{}) {
	if test {
		if len > 0 {
			reflect.ValueOf(v).Elem().Set(reflect.MakeSlice(reflect.ValueOf(v).Elem().Type(), len, len))
			b.MRead(reflect.ValueOf(v).Elem().Interface())
			return
		}
		b.MRead(v)
	}
}

func (b *ReadWriter) MRead(v ...interface{}) {
	for _, x := range v {
		if err := binary.Read(b, binary.BigEndian, x); err != nil {
			b.err = errors.WithStack(err)
			return
		}
	}
}

func (b *ReadWriter) Write(p []byte) (n int, err error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

func (b *ReadWriter) MWrite(v ...interface{}) {
	for _, x := range v {
		if err := binary.Write(b, binary.BigEndian, x); err != nil {
			b.err = errors.WithStack(err)
			return
		}
	}
}

func (b *ReadWriter) Read(p []byte) (n int, err error) {
	data := b.ReadBytes(len(p))
	if b.err != nil {
		return 0, b.err
	}
	copy(p, data)
	return len(p), nil
}

func (b *ReadWriter) ReadUint8() uint8 {
	return b.ReadBytes(1)[0]
}

func (b *ReadWriter) ReadBool() bool {
	return Uint82Bool(b.ReadUint8())
}

func (b *ReadWriter) ReadMBool() (bool, bool, bool, bool, bool, bool, bool, bool) {
	return Uint82Bools(b.ReadUint8())
}

func (b *ReadWriter) ReadUint32() uint32 {
	return uint32(Bytes2Uint64(b.ReadBytes(4)))
}

func (b *ReadWriter) ReadUint48() uint64 {
	return Bytes2Uint64(b.ReadBytes(6))
}

func (b *ReadWriter) ReadUint56() uint64 {
	return Bytes2Uint64(b.ReadBytes(7))
}

func (b *ReadWriter) ReadUint64() uint64 {
	return Bytes2Uint64(b.ReadBytes(8))
}

func (b *ReadWriter) Error(info interface{}) error {
	return errors.WithMessagef(
		b.err,
		"decode err:length not correct when decode ie:%s",
		dump.Sdump(info),
	)
}

func (b *ReadWriter) Bytes() []byte {
	if b.index >= len(b.data) {
		return nil
	}
	data := b.data[b.index:]
	b.index = len(data)
	return data
}

func (b *ReadWriter) ReadUint40() uint64 {
	return Bytes2Uint64(b.ReadBytes(5))
}

func (b *ReadWriter) ReadUint24() uint32 {
	return uint32(Bytes2Uint64(b.ReadBytes(3)))
}

func (b *ReadWriter) ReadUint16() uint16 {
	return uint16(Bytes2Uint64(b.ReadBytes(2)))
}

func (b *ReadWriter) ReadBytes(num int) []byte {
	if num+b.index > len(b.data) {
		b.err = errors.WithMessagef(io.EOF,"slice:index large chan slice len")
		return make([]byte, num)
	}
	data := b.data[b.index : b.index+num]
	b.index = b.index + num
	return data
}
