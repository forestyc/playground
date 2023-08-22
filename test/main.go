package main

import (
	"fmt"
	"github.com/forestyc/playground/pkg/util"
)

func main() {
	fmt.Println(util.StandardTime(2024))
}

func G1(ch chan bool) {
	var x int
	go func() { x = 1 }()
	fmt.Println(x)
}

func f(id int) {
	switch id {
	case 1, 2:
		fmt.Println("match")
	default:
		fmt.Println("not match")
	}
}
