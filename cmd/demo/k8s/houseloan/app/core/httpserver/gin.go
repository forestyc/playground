package httpserver

import (
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/context"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	ctx    context.Context
	engine *gin.Engine
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		engine: gin.Default(),
	}
}

func (s *HttpServer) Serve() {

}
