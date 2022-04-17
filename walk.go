package pfcp

import (
	dump "github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"reflect"
)

type operation uint8

const (
	opGet operation = iota
	opSet
)

// groupWalk provide a uniform server to encode and decode opt which base on  walk over group ie or message
func groupWalk(_group interface{}, opt operation, handle func(group reflect.Value, pos int, slice int, isGroup bool) error) error {
	if !isGrouped(_group) {
		return errors.Errorf("not a group:%v", dump.Sdump(_group))
	}
	group := reflect.ValueOf(_group)
	walk := group.Elem()
	for i := 1; i < walk.NumField(); i++ {
		switch walk.Field(i).Kind() {
		case reflect.Ptr:
			obj := walk.Field(i).Interface()
			isGroup := isGrouped(obj)
			if err := handle(group, i, -1, isGroup); err != nil {
				return err
			}
			if !isGroup || walk.Field(i).IsNil() {
				continue
			}
			if err := groupWalk(obj, opt, handle); err != nil {
				return err
			}
		case reflect.Slice:
			if opt == opSet {
				if err := handle(group, i, 0, false); err != nil {
					return err
				}
				continue
			}
			for j := 0; j < walk.Field(i).Len(); j++ {
				obj := walk.Field(i).Index(j).Interface()
				isGroup := isGrouped(obj)
				if err := handle(group, i, j, isGroup); err != nil {
					return err
				}
				if !isGroup || walk.Field(i).Index(j).IsNil() {
					continue
				}
				if err := groupWalk(obj, opt, handle); err != nil {
					return err
				}
			}
		default:
			return errors.Errorf("unsupported type,%s", dump.Sdump(walk.Field(i).Interface()))
		}
	}
	return nil
}

func isGrouped(t interface{}) bool {
	if _, ok := t.(NonGroupIE); ok {
		return false
	}
	return true
}

func getHeader(h interface{}) header {
	return reflect.ValueOf(h).Elem().FieldByName("Header").Addr().Interface().(header)
}

func fillHeader(h header, t uint16, len uint16) header {
	h.setLength(len)
	return h.setType(t)
}
