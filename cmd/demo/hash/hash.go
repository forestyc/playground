package main

import (
	"fmt"
	"hash/maphash"
)

func main() {
	seed := maphash.MakeSeed()
	for i := 0; i < 100; i++ {
		fmt.Println(String(seed, "ZC401C1010#kjajdsfadjf#aaaaa"))
	}
}

func String(seed maphash.Seed, s string) uint64 {
	var h maphash.Hash
	h.SetSeed(seed)
	h.WriteString(s)
	return h.Sum64()
}
