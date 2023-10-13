package main

import (
	"context"
	"fmt"
	"time"
)

type Chnl chan int

func main() {
	chnl1 := make(Chnl, 1)
	chnl2 := make(Chnl, 1)
	ctx, cancel := context.WithCancel(context.Background())
	go goroutine1(ctx, chnl1)
	go goroutine2(ctx, chnl2)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)
}

func goroutine1(ctx context.Context, chnl Chnl) {
	fmt.Println("load goroutine1")
	for {
		select {
		case <-chnl:
			fmt.Println("goroutine1")
		case <-ctx.Done():
			fmt.Println("goroutine1 done", ctx.Err())
			close(chnl)
			return
		}
	}
}

func goroutine2(ctx context.Context, chnl Chnl) {
	fmt.Println("load goroutine2")
	for {
		select {
		case <-chnl:
			fmt.Println("goroutine2")
		case <-ctx.Done():
			fmt.Println("goroutine2 done", ctx.Err())
			close(chnl)
			return
		}
	}
}
