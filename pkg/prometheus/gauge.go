package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Gauge struct {
	vec  *prometheus.GaugeVec
	once sync.Once
}

var g Gauge

func NewGauge(name, help string, label ...string) *Gauge {
	g.once.Do(func() {
		g.vec = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: name,
				Help: help,
			}, label,
		)
		prometheus.MustRegister(g.vec)
	})
	return &g
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
