package svc

import (
	"github.com/forestyc/playground/go-zero/epidemic/api/internal/config"
	"github.com/forestyc/playground/go-zero/epidemic/api/internal/middleware"
	"github.com/forestyc/playground/go-zero/epidemic/rpc/user/user"
	"github.com/forestyc/playground/go-zero/epidemic/rpc/user/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Auth    rest.Middleware
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Auth:    middleware.NewAuthMiddleware(c.Token).Handle,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
