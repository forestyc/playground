package context

import (
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/model/config"
	"github.com/forestyc/playground/pkg/db"
	"github.com/forestyc/playground/pkg/http"
	"github.com/forestyc/playground/pkg/log/zap"
	"github.com/forestyc/playground/pkg/redis"
)

// Context 全局context
type Context struct {
	C          config.Config
	Db         *db.Mysql
	Cache      *redis.Redis
	Logger     *zap.Zap
	HttpServer *http.Server
}

func NewContext(c config.Config) (Context, error) {
	ctx := Context{
		C:          c,
		Db:         db.NewMysql(c.Database),
		Logger:     zap.NewZap(c.Log),
		HttpServer: http.NewServer(c.Server.Addr),
	}
	// redis
	r, err := redis.NewRedis(c.Redis)
	if err != nil {
		return ctx, err
	}
	ctx.Cache = r
	return ctx, err
}
