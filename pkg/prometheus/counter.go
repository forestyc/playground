package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Counter struct {
	vec *prometheus.CounterVec
}

func NewCounter(name, help string, label ...string) *Counter {
	var c Counter
	c.vec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		}, label,
	)
	prometheus.MustRegister(c.vec)
	return &c
}

func (c *Counter) Inc(labelVal ...string) {
	c.vec.WithLabelValues(labelVal...).Inc()
}

func (c *Counter) Add(val float64, labelVal ...string) {
	c.vec.WithLabelValues(labelVal...).Add(val)
}

func (c *Counter) Unregister() {
	prometheus.Unregister(c.vec)
}
