package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Histogram struct {
	vec *prometheus.HistogramVec
}

func NewHistogram(name, help string, label ...string) *Histogram {
	var h Histogram
	h.vec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: name,
			Help: help,
		}, label,
	)
	prometheus.MustRegister(h.vec)
	return &h
}

func (h *Histogram) Observe(val float64, labelValue ...string) {
	h.vec.WithLabelValues(labelValue...).Observe(val)
}

func (h *Histogram) Unregister() {
	prometheus.Unregister(h.vec)
}
