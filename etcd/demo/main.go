package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	cli    *clientv3.Client
	ctx    context.Context
	cancel context.CancelFunc
)

type Callback func(clientv3.WatchChan)

func main() {
	var err error
	if err = InitEtcd(); err != nil {
		fmt.Println(err)
		return
	}
	defer cli.Close()
	//ctxTimeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	Watch(context.TODO(), WatchKey1, "key1")
}

// InitEtcd 初始化
func InitEtcd() (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"baal1990.asia:23790"},
		DialTimeout: 5 * time.Second,
		Username:    "root",
		Password:    "ETCD@1234567",
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

// Get 读
func Get(ctx context.Context, key string, opts ...clientv3.OpOption) (results []string, err error) {
	var resp *clientv3.GetResponse
	resp, err = cli.Get(ctx, key, opts...)
	if err != nil {
		return results, err
	}
	resultLen := len(resp.Kvs)
	if resultLen == 0 {
		return results, nil
	} else {
		for _, kv := range resp.Kvs {
			results = append(results, string(kv.Value))
		}
	}
	return results, err
}

func Watch(ctx context.Context, callback Callback, key string, opts ...clientv3.OpOption) {
	watchChnl := cli.Watch(ctx, key, opts...)
	callback(watchChnl)
}

func WatchKey1(chnl clientv3.WatchChan) {
	for {
		select {
		case events := <-chnl:
			//fmt.Println(events)
			//return
			for _, e := range events.Events {
				fmt.Println("watch key:", string(e.Kv.Key), ",value:", string(e.Kv.Value))
			}
		}
	}
}
