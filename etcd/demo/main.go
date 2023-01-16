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

func main() {
	var err error
	if err = InitEtcd(); err != nil {
		fmt.Println(err)
		return
	}
	defer cli.Close()
	value, err := Get("key1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
	cancel()
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
func Get(key string, opts ...clientv3.OpOption) (results []string, err error) {
	var resp *clientv3.GetResponse
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
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
