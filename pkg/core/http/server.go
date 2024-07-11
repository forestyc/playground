package http

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

type Handler interface {
	Register(*gin.Engine)
}

type Server struct {
	server   *http.Server
	router   *gin.Engine
	handlers []Handler
}

type Option func(*Server)

func NewServer(addr string, options ...Option) *Server {
	router := gin.Default()
	server := &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		router: router,
	}
	for _, option := range options {
		option(server)
	}
	return server
}

func (s *Server) Serve() {
	go func() {
		for _, handler := range s.handlers {
			handler.Register(s.router)
		}
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		panic(err)
	}
}

func (s *Server) WithHandler(handlers ...Handler) *Server {
	s.handlers = append(s.handlers, handlers...)
	return s
}

func WithPrometheus(uri string) Option {
	return func(server *Server) {
		server.router.GET(uri, gin.WrapH(promhttp.Handler()))
	}
}

func WithProbeLiveness() Option {
	return func(server *Server) {
		server.router.GET("/probe-liveness", probeLiveness())
	}
}
