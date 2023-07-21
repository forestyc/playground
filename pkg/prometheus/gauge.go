package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Gauge struct {
	prometheus.GaugeOpts
	vec *prometheus.GaugeVec
}

func NewGauge(name, help string, label ...string) *Gauge {
	return poolGauge.Get(name, help, label...)
}

func (g *Gauge) Inc(labelValue ...string) {
	if !Enable() {
		return
	}
	g.vec.WithLabelValues(labelValue...).Inc()
}

func (g *Gauge) Dec(labelValue ...string) {
	if !Enable() {
		return
	}
	g.vec.WithLabelValues(labelValue...).Dec()
}

func (g *Gauge) Add(val float64, labelValue ...string) {
	if !Enable() {
		return
	}
	g.vec.WithLabelValues(labelValue...).Add(val)
}

func (g *Gauge) Sub(val float64, labelValue ...string) {
	if !Enable() {
		return
	}
	g.vec.WithLabelValues(labelValue...).Sub(val)
}

func (g *Gauge) Unregister() {
	if !Enable() {
		return
	}
	prometheus.Unregister(g.vec)
}

var (
	poolGauge = &gaugePool{}
)

type gaugePool struct {
	pool map[string]*Gauge
	lock sync.Mutex
}

func (gp *gaugePool) Get(name, help string, label ...string) *Gauge {
	gp.lock.Lock()
	defer gp.lock.Unlock()
	if gp.pool == nil {
		gp.pool = make(map[string]*Gauge)
	}
	c, has := gp.pool[name]
	if has {
		return c
	} else {
		c = &Gauge{
			vec: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: name,
					Help: help,
				}, label),
		}
		prometheus.MustRegister(c.vec)
		gp.pool[name] = c
		return c
	}
}
