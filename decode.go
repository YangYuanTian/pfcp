package pfcp

import (
	"github.com/pkg/errors"
	"io"
	"reflect"
)

//NewMessage build a message from bytes
func NewMessage(buf []byte) (Message, error) {
	r := NewReadWriter(buf)
	h := new(MsgHeader)
	if err := h.decode(r); err != nil {
		return nil, err
	}
	msg, err := fct.new(h.Type)
	if err != nil {
		return nil, err
	}
	setHeader(msg, *h)
	for ie, err := ieDecode(r); errors.Cause(err) != io.EOF; ie, err = ieDecode(r) {
		if err != nil {
			return nil, err
		}
		if err := putIE(msg, ie); err != nil {
			return nil, err
		}
	}
	return msg.(Message), nil
}

//NewIE build a IE from bytes
func NewIE(i []byte) (IE, error) {
	return ieDecode(NewReadWriter(i))
}

func ieDecode(r Reader) (IE, error) {
	h := new(IEHeader)
	if err := h.decode(r); err != nil {
		return nil, err
	}
	rIE, err := fct.new(h.Type)
	if err != nil {
		return nil, err
	}
	r1 := NewReadWriter(r.ReadBytes(int(h.Length)))
	setHeader(rIE, *h)
	if !isGrouped(rIE) {
		return rIE.(IE), rIE.(NonGroupIE).decode(r1)
	}
	for ie, err := ieDecode(r1); errors.Cause(err) != io.EOF; ie, err = ieDecode(r1) {
		if err != nil {
			return nil, err
		}
		if err := putIE(rIE, ie); err != nil {
			return ie, err
		}
	}
	return rIE.(IE), r.Error(rIE)
}

func setHeader(s interface{}, h interface{}) {
	reflect.ValueOf(s).Elem().FieldByName("Header").Set(reflect.ValueOf(h))
}

func put(ie IE) func(group reflect.Value, pos int, slice int, isGroup bool) error {
	return func(group reflect.Value, pos int, slice int, isGroup bool) error {
		t := group.Elem().Field(pos).Type()
		which := group.Elem().Field(pos)
		if slice >= 0 {
			t = t.Elem()
		}
		if t == reflect.TypeOf(ie) {
			if slice >= 0 {
				which.Set(reflect.Append(which, reflect.ValueOf(ie)))
				return nil
			}
			which.Set(reflect.ValueOf(ie))
			return nil
		}
		return nil
	}
}

func putIE(container interface{}, ie IE) error {
	return groupWalk(container, opSet, put(ie))
}
