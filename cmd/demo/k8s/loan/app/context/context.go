package context

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/entity/config"
	"github.com/forestyc/playground/pkg/core/db"
	"github.com/forestyc/playground/pkg/core/http"
)

// Context 全局context
type Context struct {
	C          config.Config
	HttpServer *http.Server
	Db         *db.Mysql
}

func NewContext(c config.Config) *Context {
	ctx := &Context{
		C:          c,
		HttpServer: http.NewServer(c.Server.Addr, http.WithProbeLiveness()),
		Db:         db.NewMysql(c.Database),
	}
	return ctx
}
