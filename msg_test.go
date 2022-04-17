package pfcp

import (
	"fmt"
	dump "github.com/davecgh/go-spew/spew"
	"testing"
)

func EncodeAndDecodeForMsg(msg Message) error {
	dump.Dump("before to encode===========================", msg)
	data, err := Encode(msg)
	if err != nil {
		return err
	}
	dump.Dump("encoded bytes", data)
	WritePacket(fmt.Sprintf("%T.pcap", msg)[6:], data)
	msg1, err := NewMessage(data)
	if err != nil {
		return err
	}
	dump.Dump("decoded msg", msg1)
	return nil
}

func ErrLog() {
	if err := recover(); err != nil {
		fmt.Printf("err:%+v\n", err)
	}
}
func TestHeartbeatRequest(t *testing.T) {
	//defer ErrLog()
	msg := &HeartbeatRequest{
		RecoveryTimeStamp: &RecoveryTimeStamp{
			Value: 0x12345678,
		},
		SourceIPAddress: &SourceIPAddress{
			V4:          true,
			V6:          false,
			IPv4Address: []byte{0x01, 0x02, 0x03, 0x04},
			MaskLen:     16,
		},
	}
	fmt.Printf("%+v\n", EncodeAndDecodeForMsg(msg))
}
