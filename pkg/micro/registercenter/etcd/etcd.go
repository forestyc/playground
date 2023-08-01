package etcd

import (
	"context"
	"fmt"
	"github.com/forestyc/playground/pkg/micro/registercenter"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Etcd struct {
	conf          Config
	cli           *clientv3.Client
	cache         registercenter.ServiceCache
	stopKeepAlive context.CancelFunc
	stopWatch     context.CancelFunc
}

type Config struct {
	Ttl    int64
	Config clientv3.Config
}

func NewEtcd(config Config) (*Etcd, error) {
	var err error
	client := &Etcd{}
	client.cli, err = clientv3.New(config.Config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client.cache = registercenter.NewServiceCache()
	return client, err
}

func (e *Etcd) Register(name, url string) error {
	var err error
	var leaseId clientv3.LeaseID
	var ctx context.Context
	// keepalive
	ctx, e.stopKeepAlive = context.WithCancel(context.Background())
	leaseId, err = e.keepAlive(ctx, e.conf.Ttl)
	if err != nil {
		return err
	}
	// register
	_, err = e.cli.Put(ctx, name, url, clientv3.WithLease(leaseId))
	return err
}

func (e *Etcd) GetService(name string) string {
	return ""
}

func (e *Etcd) Close() {
	if e.stopKeepAlive != nil {
		e.stopKeepAlive()
	}
	if e.stopWatch != nil {
		e.stopWatch()
	}
}

func (e *Etcd) keepAlive(ctx context.Context, ttl int64) (clientv3.LeaseID, error) {
	result, err := e.cli.Grant(context.Background(), ttl)
	if err != nil {
		return 0, nil
	}
	errChan := make(chan error)
	go func() {
		rspChan, err := e.cli.KeepAlive(ctx, result.ID)
		if err != nil {
			errChan <- err
		}
		errChan <- nil
		for {
			select {
			case _, ok := <-rspChan:
				if !ok {
					// lease has been expired or revoked
					break
				}
			case <-ctx.Done():
				break
			}
		}
	}()
	return result.ID, <-errChan
}

func (e *Etcd) addWatch(name string) {
	var ctx context.Context
	ctx, e.stopWatch = context.WithCancel(context.Background())
	watchChnl := e.cli.Watch(ctx, name, clientv3.WithPrefix())
	go func(name string) {
		for {
			select {
			case events := <-watchChnl:
				for _, event := range events.Events {
					// todo: getservice不应该有groupname
					// e.cache.SetService(name, )
					fmt.Println("watch key:", string(event.Kv.Key), ",value:", string(event.Kv.Value))
				}
			case <-ctx.Done():
				return
			}
		}
	}(name)
}
