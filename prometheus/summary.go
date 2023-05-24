package prometheus

import "github.com/prometheus/client_golang/prometheus"

type Summary struct {
	vec *prometheus.SummaryVec
}

func NewSummary(name, help string, label ...string) Summary {
	s := Summary{
		vec: prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name: name,
				Help: help,
			}, label),
	}
	prometheus.MustRegister(s.vec)
	return s
}

func (s Summary) Observe(val float64, labelValue ...string) {
	if !Enable() {
		return
	}
	s.vec.WithLabelValues(labelValue...).Observe(val)
}
