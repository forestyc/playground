package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	ctx, _ := context.WithTimeout(context.Background(), 12*time.Second)
	for {
		timer.Reset(3 * time.Second)
		select {
		case <-timer.C:
			fmt.Println("timer working")
		case <-ctx.Done():
			if t, ok := ctx.Deadline(); ok {
				fmt.Println("deadline", t)
			} else {
				fmt.Println("canceled")
			}

			timer.Stop()
			return
		}
	}
}
