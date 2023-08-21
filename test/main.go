package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		f(i)
	}

	//select {}
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
