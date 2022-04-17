package pfcp

import "reflect"

type MessageType uint16

type MsgHeader struct {
	Version  uint8
	MP, S    bool
	Type     MessageType
	Length   uint16
	SeqNum   uint32 //3 oct
	SEID     uint64
	Priority uint8
}

func (m *MsgHeader) setType(u uint16) header{
	m.Type = MessageType(u)
	return m
}

func (m *MsgHeader) setLength(u uint16) header{
	m.Length = u
	return m
}

func (m *MsgHeader) encode(w Writer) error {
	w.MWrite(
		((m.Version&7)<<5)|Bool2Uint8(m.MP)<<1|Bool2Uint8(m.S),
		uint8(m.Type),
		m.Length,
		Uint642Bytes(uint64(m.SeqNum))[5:],
	)
	w.TestWrite(m.S, m.SEID)
	w.MWrite(Bool2Uint8(m.MP) * m.Priority)
	return w.Error(m)
}

func (m *MsgHeader) decode(r Reader) error {
	flag := r.ReadUint8()
	m.Version = flag >> 5
	m.MP = (flag>>1)&1 == 1
	m.S = flag&1 == 1
	m.Type = MessageType(r.ReadUint8())
	m.Length = r.ReadUint16()
	m.SeqNum = r.ReadUint24()
	r.TestRead(m.S, 0, &m.SEID)
	m.Priority = r.ReadUint8()
	return r.Error(m)
}

func (m *MsgHeader) length() uint16 {
	if m.S {
		return 16
	}
	return 8
}

type IEType uint16

type IEHeader struct {
	Type   IEType
	Length uint16
}

func (i *IEHeader) setType(u uint16) header{
	i.Type = IEType(u)
	return i
}

func (i *IEHeader) setLength(u uint16) header{
	i.Length = u
	return i
}

func (i *IEHeader) encode(w Writer) error {
	w.MWrite(i.Type, i.Length)
	return w.Error(i)
}

func (i *IEHeader) decode(r Reader) error {
	i.Type = IEType(r.ReadUint16())
	i.Length = r.ReadUint16()
	return r.Error(i)
}

func (i *IEHeader) length() uint16 {
	return 4
}

//GetMsgHeader returns the header of the message
func GetMsgHeader(msg Message) *MsgHeader {
	m := reflect.ValueOf(msg)
	return m.Elem().FieldByName("Header").Addr().Interface().(*MsgHeader)
}

//GetIEHeader returns the header of the IE
func GetIEHeader(ie IE) *IEHeader {
	m := reflect.ValueOf(ie)
	return m.Elem().FieldByName("Header").Addr().Interface().(*IEHeader)
}
