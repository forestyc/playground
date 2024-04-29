package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	Router *gin.Engine
}

func NewServer(addr string) *Server {
	router := gin.Default()
	server := &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		Router: router,
	}
	return server
}

func (s *Server) Serve() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
