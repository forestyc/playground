package context

import (
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/model/config"
	"github.com/forestyc/playground/pkg/core/http"
)

// Context 全局context
type Context struct {
	C          config.Config
	HttpServer *http.Server
}

func NewContext(c config.Config) Context {
	ctx := Context{
		C:          c,
		HttpServer: http.NewServer(c.Server.Addr),
	}
	return ctx
}
