package http

import (
	"context"
	"errors"
	"github.com/forestyc/playground/pkg/prometheus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

type Server struct {
	server        *http.Server
	Router        *gin.Engine
	PromCounter   *prometheus.Counter
	PromGauge     *prometheus.Gauge
	PromHistogram *prometheus.Histogram
	PromSummary   *prometheus.Summary
}

type Option func(*Server)

func NewServer(addr string, options ...Option) *Server {
	router := gin.Default()
	server := &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		Router: router,
	}
	for _, option := range options {
		option(server)
	}
	return server
}

func (s *Server) Serve() {
	go func() {
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

func WithPrometheus(uri string) Option {
	return func(server *Server) {
		server.Router.GET(uri, gin.WrapH(promhttp.Handler()))
	}
}
