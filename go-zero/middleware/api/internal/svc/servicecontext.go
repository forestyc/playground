package svc

import (
	"github.com/forestyc/playground/go-zero/middleware/api/internal/config"
	"github.com/forestyc/playground/go-zero/middleware/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
