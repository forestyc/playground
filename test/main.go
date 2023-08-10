package main

import "fmt"

func main() {
	for i := 0; i < 1000; i++ {
		G1(nil)
	}

	//select {}
}

func G1(ch chan bool) {
	var x int
	go func() { x = 1 }()
	fmt.Println(x)
}
