package main

import (
	"fmt"
	"github.com/forestyc/playground/pkg/distributed/lock"
	redis2 "github.com/forestyc/playground/pkg/distributed/lock/redis"
	"github.com/forestyc/playground/pkg/redis"
)

func main() {
	cli, err := redis.NewRedis(redis.Config{
		Address:          "140.143.163.171:6379",
		Password:         "k8s-node1#12345",
		MaxOpen:          10,
		IdleConns:        300,
		IdleTimout:       5,
		OperationTimeout: 10,
	})
	if err != nil {
		return
	}
	var l lock.Lock
	l = redis2.NewLock(cli, "lock", 30)

	fmt.Println(l.Lock())
	l.Unlock()
}
