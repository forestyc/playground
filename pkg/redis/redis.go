package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// Config 配置信息
type Config struct {
	Address          string `mapstructure:"address"`
	Password         string `mapstructure:"password"`
	MaxOpen          int    `mapstructure:"max-open"`
	IdleConns        int    `mapstructure:"idle-conns"`
	IdleTimout       int    `mapstructure:"idle-timeout"`
	OperationTimeout int    `mapstructure:"operation-timeout"`
}

type Redis struct {
	*redis.Client
}

func NewRedis(config Config) (*Redis, error) {
	r := Redis{}
	r.Client = redis.NewClient(&redis.Options{
		Addr:         config.Address,
		Password:     config.Password,
		PoolSize:     config.MaxOpen,
		MinIdleConns: config.IdleConns,
		PoolTimeout:  time.Duration(config.IdleTimout) * time.Second,
	})
	_, err := r.Client.Ping(context.Background()).Result()
	if err != nil {
		return &r, err
	}
	return &r, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}
