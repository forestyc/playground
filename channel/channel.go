package main

import (
	"fmt"
	"time"
)

func main() {
	chnl1 := make(chan int, 10)
	chnl2 := make(chan int, 10)

	go func() {
		select {
		case <-chnl1:
			fmt.Println("chnl1")
		case <-chnl2:
			fmt.Println("chnl2")
			//default:
			//	fmt.Println("default")
		}
		fmt.Println("exit")
	}()
	//chnl1 <- 1
	time.Sleep(3 * time.Second)
}
