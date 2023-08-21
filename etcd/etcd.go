package main

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{
			"81.70.188.168:12379",
			"140.143.163.171:12379",
			"101.42.23.168:12379",
		}})
	if err != nil {
		fmt.Println(err)
		return
	}
	session, err := concurrency.NewSession(cli)
	if err != nil {
		fmt.Println(err)
		return
	}
	lock := concurrency.NewLocker(session, "lock")
	lock.Lock()
}
