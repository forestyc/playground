package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker working")
		}
	}
}
