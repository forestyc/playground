package etcd

import (
	"context"
	"github.com/forestyc/playground/pkg/micro/registercenter"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"sync/atomic"
)

type Subscriber struct {
	cli        *clientv3.Client
	stopListen context.CancelFunc
	c          cluster
}

func (s *Subscriber) Listen(endpoints []string, serviceName string) error {
	// connect etcd
	var err error
	if s.cli, err = clientv3.New(clientv3.Config{Endpoints: endpoints}); err != nil {
		return err
	}
	// get cluster
	resp, err := s.cli.Get(context.Background(), serviceName, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, kv := range resp.Kvs {
		s.c.Set(registercenter.Service{Name: string(kv.Key), Url: string(kv.Value)})
	}
	// listen and update
	var ctx context.Context
	ctx, s.stopListen = context.WithCancel(context.Background())
	watchChnl := s.cli.Watch(ctx, serviceName, clientv3.WithPrefix())
	go func() {
		for {
			select {
			case events := <-watchChnl:
				for _, event := range events.Events {
					if event.Type == mvccpb.PUT {
						s.c.Set(registercenter.Service{Name: string(event.Kv.Key), Url: string(event.Kv.Value)})
					} else {
						s.c.Delete(registercenter.Service{Name: string(event.Kv.Key), Url: string(event.Kv.Value)})
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (s *Subscriber) GetService(name string) (registercenter.Service, bool) {
	return s.c.loadBalance()
}

func (s *Subscriber) Close() error {
	if s.stopListen != nil {
		s.stopListen()
	}
	return s.cli.Close()
}

type cluster struct {
	cluster []registercenter.Service
	cursor  atomic.Int32
	lock    sync.RWMutex
}

func (c *cluster) loadBalance() (registercenter.Service, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	var service registercenter.Service
	if len(c.cluster) == 0 {
		return service, false
	}
	cursor := c.cursor.Load()
	idx := int(cursor) % len(c.cluster)
	c.cursor.Add(1)
	return c.cluster[idx], true
}

func (c *cluster) Set(service registercenter.Service) {
	c.lock.Lock()
	defer c.lock.Unlock()
	exist := false
	for i, e := range c.cluster {
		if e.Name == service.Name {
			c.cluster[i] = service
			exist = true
			break
		}
	}
	if !exist {
		c.cluster = append(c.cluster, service)
	}
}

func (c *cluster) Delete(service registercenter.Service) {
	c.lock.Lock()
	defer c.lock.Unlock()
	for i, e := range c.cluster {
		if e.Name == service.Name {
			c.cluster = append(c.cluster[0:i], c.cluster[i+1:]...)
			break
		}
	}
}
