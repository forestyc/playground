package redis

import (
	"context"
	"github.com/forestyc/playground/pkg/redis"
	"time"
)

type Lock struct {
	key string
	ttl time.Duration
	cli redis.Redis
}

// 初始化
func NewLock(cli redis.Redis, key string, ttl int) Lock {
	l := Lock{
		key: key,
		cli: cli,
		ttl: time.Second * time.Duration(ttl),
	}
	if ttl == -1 {
		l.ttl = -1
	}
	return l
}

// Lock 加锁
func (l *Lock) Lock() error {
	// todo: 权衡抢占、cpu负载
	ticker := time.NewTicker(l.ttl / 2)
	defer ticker.Stop()
	select {
	case <-ticker.C:
		if ok, err := l.lock(); err != nil {
			return err
		} else if ok {
			return nil
		}
	}
	return nil
}

// Unlock 解锁
func (l *Lock) Unlock() error {
	return l.unlock()
}

func (l *Lock) lock() (bool, error) {
	result := l.cli.SetNX(context.TODO(), l.key, 1, l.ttl)
	return result.Val(), result.Err()
}

func (l *Lock) unlock() error {
	result := l.cli.Del(context.TODO(), l.key)
	return result.Err()
}
