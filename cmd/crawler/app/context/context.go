package context

import (
	"sync"

	"github.com/forestyc/playground/pkg/db"
	"github.com/forestyc/playground/pkg/redis"

	"github.com/forestyc/playground/cmd/crawler/app/config"
	"github.com/forestyc/playground/pkg/log/zap"
)

// Context 全局context
type Context struct {
	C      config.Config
	Db     db.Mysql
	Cache  redis.Redis
	Logger zap.Zap
	Wg     *sync.WaitGroup
}

func NewContext(c config.Config) (Context, error) {
	// 注意: logger不在此处初始化，各爬虫内初始化
	ctx := Context{
		C:  c,
		Db: db.NewMysql(c.Database),
		Wg: &sync.WaitGroup{},
	}
	r, err := redis.NewRedis(c.Redis)
	if err != nil {
		return ctx, err
	}
	ctx.Cache = r
	return ctx, err
}
