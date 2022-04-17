package pfcp

import (
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
	"time"
)

func Bool2Uint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
func Uint82Bool(u uint8) bool {
	if u == 1 {
		return true
	}
	return false
}
func Uint82Bools(u uint8) (bit8, bit7, bit6, bit5, bit4, bit3, bit2, bit1 bool) {
	return u>>7 == 1, u>>6&1 == 1, u>>5&1 == 1, u>>4&1 == 1, u>>3&1 == 1, u>>2&1 == 1, u>>1&1 == 1, u&1 == 1
}
func Bools2Uint8(bit8, bit7, bit6, bit5, bit4, bit3, bit2, bit1 bool) uint8 {
	return Bool2Uint8(bit8)<<7 | Bool2Uint8(bit7)<<6 | Bool2Uint8(bit6)<<5 | Bool2Uint8(bit5)<<4 | Bool2Uint8(bit4)<<3 | Bool2Uint8(bit3)<<2 | Bool2Uint8(bit2)<<1 | Bool2Uint8(bit1)
}

func Bytes2Uint64(b []byte) uint64 {
	if len(b) > 8 {
		b = b[:8]
	}
	if len(b) < 8 {
		b = append(make([]byte, 8-len(b)), b...)
	}
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
}
func Uint642Bytes(u uint64) []byte {
	return []byte{byte(u >> 56), byte(u >> 48), byte(u >> 40), byte(u >> 32), byte(u >> 24), byte(u >> 16), byte(u >> 8), byte(u)}
}
func DigitSting2Uint8(s string) uint8 {
	switch s {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	default:
		return 0
	}
}
func DigitsSting2Bytes(s string) []byte {
	ret := make([]byte, 0, len(s))
	for _, v := range s {
		switch v {
		case '1':
			ret = append(ret, 1)
		case '2':
			ret = append(ret, 2)
		case '3':
			ret = append(ret, 3)
		case '4':
			ret = append(ret, 4)
		case '5':
			ret = append(ret, 5)
		case '6':
			ret = append(ret, 6)
		case '7':
			ret = append(ret, 7)
		case '8':
			ret = append(ret, 8)
		case '9':
			ret = append(ret, 9)
		case '0':
			ret = append(ret, 0)
		case 'F':
			ret = append(ret, 15)
		default:
		}
	}
	return ret
}
func Bytes2DigitSting(b []byte) string {
	var ret string
	for _, v := range b {
		switch v {
		case 1:
			ret += "1"
		case 2:
			ret += "2"
		case 3:
			ret += "3"
		case 4:
			ret += "4"
		case 5:
			ret += "5"
		case 6:
			ret += "6"
		case 7:
			ret += "7"
		case 8:
			ret += "8"
		case 9:
			ret += "9"
		case 0:
			ret += "0"
		case 15:
			ret += "F"
		default:
		}
	}
	return ret
}
func DigitsBytes2EncodeBytes(digits []byte) []byte {
	ret := make([]byte, (len(digits))/2)
	for i := 0; i < len(digits); i += 2 {
		ret[i/2] = digits[i] + digits[i+1]<<4
	}
	return ret
}
func DecodeBytes2DigitsBytes(b []byte) []byte {
	ret := make([]byte, len(b)*2)
	for i := 0; i < len(b); i++ {
		ret[i*2] = b[i] & 0xf
		ret[i*2+1] = b[i] >> 4
	}
	return ret
}
func StringReverse(string2 string) string {
	ret := ""
	for i := len(string2) - 1; i >= 0; i-- {
		ret += string2[i : i+1]
	}
	return ret
}
func Convert2Uint8(n interface{}) uint8 {
	switch reflect.ValueOf(n).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(n).Convert(reflect.TypeOf(uint8(0))).Interface().(uint8)
	}
	return uint8(0)
}

type PcapGlobalHear struct {
	Magic    uint32 //：4Byte：标记文件开始，并用来识别文件自己和字节顺序。0xa1b2c3d4用来表示按照原来的顺序读取，0xd4c3b2a1表示下面的字节都要交换顺序读取。考虑到计算机内存的存储结构，一般会采用0xd4c3b2a1，即所有字节都需要交换顺序读取。
	Major    uint16 //：2Byte： 当前文件主要的版本号，一般为 0x0200【实际上因为需要交换读取顺序，所以计算机看到的应该是 0x0002】
	Minor    uint16 //：2Byte： 当前文件次要的版本号，一般为 0x0400【计算机看到的应该是 0x0004】
	ThisZone uint32 //：4Byte：当地的标准时间，如果用的是GMT则全零，一般都直接写
	SigFigs  uint32 //：4Byte：时间戳的精度，设置为 全零 即可
	SnapLen  uint32 //：4Byte：最大的存储长度，如果想把整个包抓下来，设置为 ，但一般来说 ff7f 0000就足够了【计算机看到的应该是 0000 ff7f 】
	LinkType uint32 //：4Byte：链路类型  以太网或者环路类型为
}

func (h *PcapGlobalHear) SetDefault() *PcapGlobalHear {
	h.Magic = 0xa1b2c3d4
	h.Major = 0x0002
	h.Minor = 0x0004
	h.SnapLen = 0x0000ff7f
	h.LinkType = 1
	return h
}
func (h *PcapGlobalHear) Marshal() []byte {
	buf := make([]byte, 24)
	binary.LittleEndian.PutUint32(buf[0:4], h.Magic)
	binary.LittleEndian.PutUint16(buf[4:6], h.Major)
	binary.LittleEndian.PutUint16(buf[6:8], h.Minor)
	binary.LittleEndian.PutUint32(buf[8:12], h.ThisZone)
	binary.LittleEndian.PutUint32(buf[12:16], h.SigFigs)
	binary.LittleEndian.PutUint32(buf[16:20], h.SnapLen)
	binary.LittleEndian.PutUint32(buf[20:24], h.LinkType)
	return buf
}
func (h *PcapGlobalHear) UnMarshal(buf []byte) error {
	if len(buf) != 24 {
		return fmt.Errorf("buffer len error,unmarshal failed")
	}
	h.Magic = binary.LittleEndian.Uint32(buf[0:4])
	h.Major = binary.LittleEndian.Uint16(buf[4:6])
	h.Minor = binary.LittleEndian.Uint16(buf[6:8])
	h.ThisZone = binary.LittleEndian.Uint32(buf[8:12])
	h.SigFigs = binary.LittleEndian.Uint32(buf[12:16])
	h.SnapLen = binary.LittleEndian.Uint32(buf[16:20])
	h.LinkType = binary.LittleEndian.Uint32(buf[20:24])
	return nil
}

type PacketHeader struct {
	TimestampH uint32 //：被捕获时间的高位，单位是seconds
	TimestampL uint32 //：被捕获时间的低位，单位是microseconds
	CapLen     uint32 //：当前数据区的长度，即抓取到的数据帧长度，不包括Packet Header本身的长度，单位是 Byte ，由此可以得到下一个数据帧的位置。
	Len        uint32 //：离线数据长度：网络中实际数据帧的长度，一般不大于caplen，多数情况下和Caplen数值相等。
}

func (p *PacketHeader) Marshal() []byte {
	buf := make([]byte, 16)
	binary.LittleEndian.PutUint32(buf[0:4], p.TimestampH)
	binary.LittleEndian.PutUint32(buf[4:8], p.TimestampL)
	binary.LittleEndian.PutUint32(buf[8:12], p.CapLen)
	binary.LittleEndian.PutUint32(buf[12:16], p.Len)
	return buf
}
func (p *PacketHeader) UnMarshal(b []byte) error {
	if len(b) != 16 {
		return fmt.Errorf("length bytes should be 16,rather than %d", len(b))
	}
	p.TimestampH = binary.LittleEndian.Uint32(b[0:4])
	p.TimestampL = binary.LittleEndian.Uint32(b[4:8])
	p.CapLen = binary.LittleEndian.Uint32(b[8:12])
	p.Len = binary.LittleEndian.Uint32(b[12:16])
	return nil
}

func (p *PacketHeader) Packing(packets []byte, t *time.Time) []byte {
	p.Len = uint32(len(packets))
	if len(packets) > 0xffffffff {
		return nil
	}
	p.CapLen = p.Len
	if t != nil {
		t1 := t.UnixNano()
		p.TimestampH = uint32(t1 / 1e9)       //单位是s
		p.TimestampL = uint32(t1 % 1e9 / 1e3) //单位是ms
	} else {
		t2 := time.Now().UnixNano()
		p.TimestampH = uint32(t2 / 1e9)       //单位是s
		p.TimestampL = uint32(t2 % 1e9 / 1e3) //单位是ms
	}
	pac := make([]byte, 0, 0)
	pac = append(pac, p.Marshal()...)
	pac = append(pac, packets...)
	return pac
}

func PacketGenerator(Data []byte, Port uint16) []byte {
	var port = make([]byte, 2, 2)
	binary.BigEndian.PutUint16(port, Port)
	//ip                                          ps            pd
	head := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xa4, 0x16, 0xe7, 0x7c, 0xf5, 0x1e, 0x08, 0x00, 0x45, 0xc0, 0x01, 0x71, 0x44, 0xfd, 0x00, 0x00, 0xff, 0x11, 0x74, 0xbf, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, port[0], port[1], port[0], port[1], 0x01, 0x5d, 0x00, 0x00}
	ipLen := 16
	udpLen := 38
	data := Data
	dataLen := uint16(len(data) + 8)
	buf := make([]byte, 2, 2)
	binary.BigEndian.PutUint16(buf[0:2], dataLen)
	head[udpLen] = buf[0]
	head[udpLen+1] = buf[1]
	binary.BigEndian.PutUint16(buf[0:2], dataLen+20)
	head[ipLen] = buf[0]
	head[ipLen+1] = buf[1]
	return append(head, data...)
}
func isDirExist(dir string) bool {
	if _, err := os.Stat("./pcap"); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func WritePacket(name string, b []byte) {
	g := &PcapGlobalHear{}
	p := &PacketHeader{}
	dir := "./pcap/"
	if !isDirExist(dir) {
		if err := os.Mkdir(dir, 0666); err != nil {
			fmt.Println(err)
			return
		}
	}
	err := os.WriteFile(
		dir+name,
		append(g.SetDefault().Marshal(), p.Packing(PacketGenerator(b, 8805), nil)...),
		0666,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}
