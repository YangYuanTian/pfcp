package pfcp

import (
	"errors"
	"reflect"
)

//Encode encode message or ie to bytes,the obj pass to encode must be a interface type Message or IE
func Encode(enc interface{}) ([]byte, error) {
	var t uint16
	switch TypeOf(enc) {
	case TypeIE:
		t = uint16(enc.(IE).Type())
	case TypeMessage:
		t = uint16(enc.(Message).Type())
	default:
		return nil, errors.New("invalid type")
	}
	w := NewReadWriter(nil)
	if err := fillHeader(getHeader(enc), t, Len(enc)-4).encode(w); err != nil {
		return nil, err
	}
	err := encode(enc, w)
	return w.Bytes(), err
}


type Type string

const (
	TypeIE      Type = "IE"
	TypeMessage Type = "Message"
	Unknown     Type = "unknown"
)

func TypeOf(t interface{}) Type {
	if _, ok := t.(IE); ok {
		return TypeIE
	}
	if _, ok := t.(Message); ok {
		return TypeMessage
	}
	if _, ok := t.(IEType); ok {
		return TypeIE
	}
	if _, ok := t.(MessageType); ok {
		return TypeMessage
	}
	return Unknown
}

func getType(t interface{}) uint16 {
	switch TypeOf(t) {
	case TypeIE:
		return uint16(t.(IE).Type())
	case TypeMessage:
		return uint16(t.(Message).Type())
	default:
		return 0
	}
}

func encode(enc interface{}, w Writer) error {
	return groupWalk(enc, opGet, func(group reflect.Value, pos int, slice int, isGroup bool) error {
		var obj = group.Elem().Field(pos)
		if obj.IsNil() {
			return nil
		}
		obj_ := obj.Interface()
		if slice >= 0 {
			obj_ = obj.Index(slice).Interface()
		}
		w.MWrite(getType(obj_), Len(obj_)-4)
		if isGroup {
			return nil
		}
		return obj.Interface().(NonGroupIE).encode(w)
	})
}
