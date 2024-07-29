package snowflake

import (
	"testing"
)

var idMap map[int64]bool

func TestSameSecondAndSameNode(t *testing.T) {
	idMap = make(map[int64]bool)
	sf := New(1)
	if ret := genSameSecondAndSameNode(sf); ret == -1 {
		panic("duplicate id")
	}
}

func BenchmarkGen(b *testing.B) {
	idMap = make(map[int64]bool)
	sf := New(1)
	for n := 0; n < b.N; n++ {
		sf.Gen()
	}
}

func genSameSecondAndSameNode(sf *Snowflake) int {
	for i := 0; i < 4096; i++ {
		id := sf.Gen()
		if _, ok := idMap[id]; !ok {
			idMap[id] = true
		} else {
			return -1
		}
	}
	return 0
}
