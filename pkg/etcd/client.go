package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
)

type WathCallback func(context.Context, clientv3.WatchChan)

type Client struct {
	config clientv3.Config
	cli    *clientv3.Client
}

// NewClient 初始化
func NewClient(config clientv3.Config) (*Client, error) {
	var err error
	client := &Client{
		config: config,
	}
	client.cli, err = clientv3.New(client.config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return client, err
}

// Get 获取key
func (c *Client) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (results []string, err error) {
	var resp *clientv3.GetResponse
	resp, err = c.cli.Get(ctx, key, opts...)
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

// Put 写入key的value
func (c *Client) Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) error {
	_, err := c.cli.Put(ctx, key, val, opts...)
	return err
}

// Del 删除key
func (c *Client) Del(ctx context.Context, key string, opts ...clientv3.OpOption) error {
	_, err := c.cli.Delete(ctx, key, opts...)
	return err
}

// Watch 监控key
func (c *Client) Watch(ctx context.Context, callback WathCallback, key string, opts ...clientv3.OpOption) context.CancelFunc {
	watchChnl := c.cli.Watch(ctx, key, opts...)
	ctxChild, cancel := context.WithCancel(ctx)
	go callback(ctxChild, watchChnl)
	return cancel
}

func (c *Client) NewLease(ctx context.Context, ttl int64) (clientv3.LeaseID, error) {
	result, err := c.cli.Grant(ctx, ttl)
	return result.ID, err
}

func (c *Client) KeepAlive(ctx context.Context, leaseId clientv3.LeaseID) error {
	errChan := make(chan error)
	go func() {
		rspChan, err := c.cli.KeepAlive(ctx, leaseId)
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
	return <-errChan
}

// Close 关闭连接
func (c *Client) Close() error {
	return c.cli.Close()
}
