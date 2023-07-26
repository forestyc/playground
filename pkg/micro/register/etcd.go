package register

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Watch 监控key
func Watch(ctx context.Context, config clientv3.Config, key string, opts ...clientv3.OpOption) error {
	cli, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return err
	}
	watchChnl := cli.Watch(ctx, key, opts...)
	go WatchKey1(ctx, watchChnl)
	return nil
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
