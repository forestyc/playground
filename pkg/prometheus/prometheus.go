package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync/atomic"
)

var (
	enable atomic.Bool
)

type Prometheus struct {
	Addr string `mapstructure:"addr"`
	Path string `mapstructure:"path"`
}

// Start 开启监听服务
func (p Prometheus) Start() {
	go func() {
		http.Handle(p.Path, promhttp.Handler())
		http.ListenAndServe(p.Addr, nil)
	}()
	enable.Store(true)
}

func (p Prometheus) HasPrometheus() bool {
	if len(p.Path) != 0 && len(p.Addr) != 0 {
		return true
	}
	return false
}

func Enable() bool {
	return enable.Load()
}
