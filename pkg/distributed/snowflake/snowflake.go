package snowflake

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Snowflake struct {
	NodeId     int
	SequenceNo atomic.Int64
}

func New(nodeId int) *Snowflake {
	return &Snowflake{
		NodeId: nodeId,
	}
}

func (s *Snowflake) Gen() int64 {
	timestampField := (time.Now().Unix() & 0x1FFFFFFFFFF) << 22
	nodeIdField := int64(s.NodeId&0x3FF) << 12
	sequenceNoField := s.SequenceNo.Add(1) & 0xFFF
	return timestampField | nodeIdField | sequenceNoField
}

func main() {
	sf := New(1)
	idMap := make(map[int64]bool)
	for i := 0; i < 4096; i++ {
		id := sf.Gen()
		if _, ok := idMap[id]; !ok {
			idMap[id] = true
		} else {
			fmt.Println("id:", id)
			panic("duplicate id")
		}
	}
}
