package main

import (
	"context"
	"fmt"
	"github.com/forestyc/playground/pkg/concurrency/workpool"
	"time"
)

func main() {
	pool := workpool.NewWorkPool()
	for i := 0; i < 10; i++ {
		pool.AddJob(JobBlock(i))
	}
	pool.Start()
	go func() {
		time.Sleep(3 * time.Second)
		pool.Stop()
	}()
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println("processing")
		}
	}
}

func JobBlock(idx int) workpool.Job {
	return func(ctx context.Context) {
		fmt.Printf("JOB %d attached\n", idx+1)
		time.Sleep(time.Second * 2)
		fmt.Printf("JOB %d detached\n", idx+1)
	}
}
