package pfcp

import "reflect"

//Len get  the length of  IE or Message(if pass other type,always return 0),which is the total length of the IE or Message including the header
func Len(l interface{}) (len uint16) {
	if TypeOf(l) == Unknown {
		return
	}
	len += getHeader(l).length()
	if !isGrouped(l) {
		return len + l.(NonGroupIE).length()
	}
	groupWalk(l, opGet, func(group reflect.Value, pos int, slice int, isGroup bool) error {
		obj := group.Elem().Field(pos)
		if obj.IsNil() {
			return nil
		}
		len += getHeader(obj.Interface()).length()
		if !isGroup {
			len += obj.Interface().(NonGroupIE).length()
		}
		return nil
	})
	return
}
