package pfcp

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"reflect"
)

var fct factory = make(map[uint32]reflect.Type)

const ieStart uint32 = 256

type factory map[uint32]reflect.Type

func (f factory) new(t interface{}) (interface{}, error) {
	var typ reflect.Type
	var ok bool
	switch TypeOf(t) {
	case TypeMessage:
		typ, ok = f[uint32(t.(MessageType))]
	case TypeIE:
		typ, ok = f[ieStart+uint32(t.(IEType))]
	default:
		return nil, errors.Errorf("unknown type of %s", spew.Sdump(t))
	}
	if !ok {
		return nil, errors.New("unknown message type")
	}
	return reflect.New(typ.Elem()).Interface(), nil
}

func (f factory) store(t interface{}, v interface{}) {
	switch TypeOf(t) {
	case TypeMessage:
		f[uint32(t.(MessageType))] = reflect.TypeOf(v)
	case TypeIE:
		f[uint32(t.(IEType))+ieStart] = reflect.TypeOf(v)
	}
}
