// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/forestyc/playground/go-zero/middleware/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Example},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/search/do",
					Handler: searchHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}