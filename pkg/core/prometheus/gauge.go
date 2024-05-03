package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Gauge struct {
	vec *prometheus.GaugeVec
}

func NewGauge(name, help string, label ...string) *Gauge {
	var g Gauge
	g.vec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name,
			Help: help,
		}, label,
	)
	prometheus.MustRegister(g.vec)
	return &g
}

func (g *Gauge) Inc(labelValue ...string) {
	g.vec.WithLabelValues(labelValue...).Inc()
}

func (g *Gauge) Dec(labelValue ...string) {
	g.vec.WithLabelValues(labelValue...).Dec()
}

func (g *Gauge) Add(val float64, labelValue ...string) {
	g.vec.WithLabelValues(labelValue...).Add(val)
}

func (g *Gauge) Sub(val float64, labelValue ...string) {
	g.vec.WithLabelValues(labelValue...).Sub(val)
}

func (g *Gauge) Unregister() {
	prometheus.Unregister(g.vec)
}
