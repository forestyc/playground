package common

import (
	"encoding/binary"
	"github.com/forestyc/playground/cmd/demo/redis-transfer/proto/redispb"
	"github.com/golang/protobuf/proto"
	"unsafe"
)

const (
	Unknown = uint32(iota)
	String
	Hash
	Set
	ZSet
	List
	Stream
)

const HeadLength = 8

func RedisType(pb proto.Message) uint32 {
	if _, ok := pb.(*redispb.Hash); ok {
		return Hash
	} else if _, ok = pb.(*redispb.String); ok {
		return String
	} else if _, ok = pb.(*redispb.List); ok {
		return List
	} else if _, ok = pb.(*redispb.Set); ok {
		return Set
	} else if _, ok = pb.(*redispb.ZSet); ok {
		return ZSet
	} else if _, ok = pb.(*redispb.Stream); ok {
		return Stream
	} else {
		return Unknown
	}
}

type Header struct {
	Type   uint32
	Length uint32
}

func NewHeader(t, l uint32) Header {
	return Header{
		Type:   t,
		Length: l,
	}
}

func (h *Header) Marshal() []byte {
	var result []byte
	result = binary.LittleEndian.AppendUint32(result, h.Type)
	result = binary.LittleEndian.AppendUint32(result, h.Length)
	return result
}

func (h *Header) Unmarshal(b []byte) {
	h.Type = binary.LittleEndian.Uint32(b)
	h.Length = binary.LittleEndian.Uint32(b[unsafe.Sizeof(h.Type):])
}
