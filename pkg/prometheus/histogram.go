package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Histogram struct {
	vec  *prometheus.HistogramVec
	once sync.Once
}

var h Histogram

func NewHistogram(name, help string, label ...string) *Histogram {
	h.once.Do(func() {
		h.vec = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: name,
				Help: help,
			}, label,
		)
		prometheus.MustRegister(h.vec)
	})
	return &h
}

func (h *Histogram) Observe(val float64, labelValue ...string) {
	if !Enable() {
		return
	}
	h.vec.WithLabelValues(labelValue...).Observe(val)
}

func (h *Histogram) Unregister() {
	if !Enable() {
		return
	}
	prometheus.Unregister(h.vec)
}
