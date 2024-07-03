package context

import (
	"github.com/forestyc/playground/cmd/demo/k8s/workinghours/app/model/config"
	"github.com/forestyc/playground/pkg/core/db"
	"github.com/forestyc/playground/pkg/core/log/zap"
	"github.com/forestyc/playground/pkg/core/redis"
	"github.com/gin-gonic/gin"
)

// Context 全局context
type Context struct {
	C          config.Config
	Db         *db.Mysql
	Cache      *redis.Redis
	Logger     *zap.Zap
	HttpServer *gin.Engine
}

func NewContext(c config.Config) (Context, error) {
	// 注意: logger不在此处初始化，各爬虫内初始化
	ctx := Context{
		C:          c,
		Db:         db.NewMysql(c.Database),
		Logger:     zap.NewZap(c.Log),
		HttpServer: gin.Default(),
	}
	r, err := redis.NewRedis(c.Redis)
	if err != nil {
		return ctx, err
	}
	ctx.Cache = r
	return ctx, err
}

func (c *Context) Run() {
	err := c.HttpServer.Run(c.C.Server.Addr)
	if err != nil {
		panic(err)
	}
}
