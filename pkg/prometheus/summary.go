package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Summary struct {
	vec  *prometheus.SummaryVec
	once sync.Once
}

var s Summary

func NewSummary(name, help string, label ...string) *Summary {
	s.once.Do(func() {
		s.vec = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name: name,
				Help: help,
			}, label,
		)
		prometheus.MustRegister(s.vec)
	})
	return &s
}

func (s *Summary) Observe(val float64, labelValue ...string) {
	if !Enable() {
		return
	}
	s.vec.WithLabelValues(labelValue...).Observe(val)
}

func (s *Summary) Unregister() {
	if !Enable() {
		return
	}
	prometheus.Unregister(s.vec)
}
