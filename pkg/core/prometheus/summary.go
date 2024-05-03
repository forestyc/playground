package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Summary struct {
	vec *prometheus.SummaryVec
}

func NewSummary(name, help string, label ...string) *Summary {
	var s Summary
	s.vec = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: name,
			Help: help,
		}, label,
	)
	prometheus.MustRegister(s.vec)
	return &s
}

func (s *Summary) Observe(val float64, labelValue ...string) {
	s.vec.WithLabelValues(labelValue...).Observe(val)
}

func (s *Summary) Unregister() {
	prometheus.Unregister(s.vec)
}
