package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Counter struct {
	vec *prometheus.CounterVec
}

func NewCounter(name, help string, label ...string) Counter {
	counter := Counter{
		vec: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: name,
				Help: help,
			}, label),
	}
	prometheus.MustRegister(counter.vec)
	return counter
}

func (c Counter) Inc(labelVal ...string) {
	if !Enable() {
		return
	}
	c.vec.WithLabelValues(labelVal...).Inc()
}

func (c Counter) Add(val float64, labelVal ...string) {
	if !Enable() {
		return
	}
	c.vec.WithLabelValues(labelVal...).Add(val)
}