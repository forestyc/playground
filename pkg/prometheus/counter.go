package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Counter struct {
	vec  *prometheus.CounterVec
	once sync.Once
}

var c Counter

func NewCounter(name, help string, label ...string) *Counter {
	c.once.Do(func() {
		c.vec = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: name,
				Help: help,
			}, label,
		)
		prometheus.MustRegister(c.vec)
	})
	return &c
}

func (c *Counter) Inc(labelVal ...string) {
	if !Enable() {
		return
	}
	c.vec.WithLabelValues(labelVal...).Inc()
}

func (c *Counter) Add(val float64, labelVal ...string) {
	if !Enable() {
		return
	}
	c.vec.WithLabelValues(labelVal...).Add(val)
}

func (c *Counter) Unregister() {
	if !Enable() {
		return
	}
	prometheus.Unregister(c.vec)
}
