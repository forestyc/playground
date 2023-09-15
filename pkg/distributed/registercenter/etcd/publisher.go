package etcd

import (
	"context"
	"github.com/forestyc/playground/pkg/distributed/registercenter"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Publisher struct {
	cli           *clientv3.Client
	stopKeepAlive context.CancelFunc
}

func (p *Publisher) Register(endpoints []string, service registercenter.Service, ttl int64) error {
	var err error
	// connect etcd
	if p.cli, err = clientv3.New(clientv3.Config{Endpoints: endpoints}); err != nil {
		return err
	}
	// keepalive
	var ctx context.Context
	ctx, p.stopKeepAlive = context.WithCancel(context.Background())
	leaseId, err := p.keepAlive(ctx, ttl)
	if err != nil {
		return err
	}
	// register
	_, err = p.cli.Put(ctx, service.Name, service.Url, clientv3.WithLease(leaseId))
	return err
}

func (p *Publisher) Close() error {
	if p.stopKeepAlive != nil {
		p.stopKeepAlive()
	}
	return p.cli.Close()
}

func (p *Publisher) keepAlive(ctx context.Context, ttl int64) (clientv3.LeaseID, error) {
	lease, err := p.cli.Grant(context.Background(), ttl)
	if err != nil {
		return 0, nil
	}
	errChan := make(chan error)
	go func() {
		rspChan, err := p.cli.KeepAlive(ctx, lease.ID)
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
	return lease.ID, <-errChan
}
