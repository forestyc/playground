package context

import (
	"github.com/Baal19905/playground/colly/cmd/crawler/app/config"
	"github.com/Baal19905/playground/colly/pkg/cache/redis"
	"github.com/Baal19905/playground/colly/pkg/db/gorm"
	"github.com/Baal19905/playground/colly/pkg/log/zap"
)

// GlobalContext 配置信息
type GlobalContext struct {
	C      config.Config
	Db     gorm.Mysql
	Cache  redis.Redis
	Logger zap.Zap
}

func NewGlobalContext(c config.Config) (GlobalContext, error) {
	// 注意: logger不在此处初始化，各爬虫内初始化
	ctx := GlobalContext{
		C:  c,
		Db: gorm.NewMysql(c.Database),
	}
	r, err := redis.NewRedis(c.Redis)
	if err != nil {
		return ctx, err
	}
	ctx.Cache = r
	return ctx, err
}
