package main

import (
	"context"
	"fmt"
	"github.com/forestyc/playground/pkg/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	var err error
	var cli *etcd.Client
	if cli, err = etcd.NewClient(clientv3.Config{
		Endpoints: []string{
			"81.70.188.168:12379",
			"140.143.163.171:12379",
			"101.42.23.168:12379"},
		DialTimeout: 5 * time.Second,
	}); err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		fmt.Println(cli.Close())
	}()
	ctxTimeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	leaseId, _ := cli.NewLease(ctxTimeout, 2)
	ctxLease, cancel := context.WithCancel(context.Background())
	defer cancel()
	cli.KeepAlive(ctxLease, leaseId)
	if err = cli.Put(ctxTimeout, "api", "http://localhost:8080", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(time.Minute)
	//var result []string
	//if result, err = cli.Get(ctxTimeout, "key1"); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(result)
	//cli.Watch(context.Background(), WatchKey1, "key1")
	//exit := make(chan bool)
	//for {
	//	select {
	//	case <-exit:
	//		break
	//	}
	//}
}

func WatchKey1(ctx context.Context, chnl clientv3.WatchChan) {
	for {
		select {
		case events := <-chnl:
			for _, e := range events.Events {
				fmt.Println("watch key:", string(e.Kv.Key), ",value:", string(e.Kv.Value))
			}
		case <-ctx.Done():
			fmt.Println("WatchKey1 exit")
			return
		}
	}
}
