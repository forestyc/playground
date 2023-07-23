package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	go c()
	time.Sleep(time.Second)
	fmt.Println("main exit")
}

func c() (i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in c", r)
		}
	}()
	p()
	return 1
}

func p() {
	panic("p")
}
