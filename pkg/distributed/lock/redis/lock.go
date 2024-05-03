package redis

import (
	"context"
	"github.com/forestyc/playground/pkg/core/redis"
	"time"
)

type Lock struct {
	key string
	ttl time.Duration
	cli redis.Redis
}

// NewLock 初始化
func NewLock(cli redis.Redis, key string, ttl int) *Lock {
	l := Lock{
		key: key,
		cli: cli,
		ttl: time.Second * time.Duration(ttl),
	}
	if ttl == -1 {
		l.ttl = -1
	}
	return &l
}

// Lock 加锁
func (l *Lock) Lock() error {
	for {
		if ok, err := l.lock(); err != nil {
			return err
		} else if ok {
			return nil
		}
	}
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
