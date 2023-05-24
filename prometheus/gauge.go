package prometheus

import "github.com/prometheus/client_golang/prometheus"

type Gauge struct {
	prometheus.GaugeOpts
	vec *prometheus.GaugeVec
}

func NewGauge(name, help string, label ...string) Gauge {
	gauge := Gauge{
		vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: name,
			Help: help,
		}, label),
	}
	prometheus.MustRegister(gauge.vec)
	return gauge
}

func (g Gauge) Inc(labelValue ...string) {
	if !Enable() {
		return
	}
	g.vec.WithLabelValues(labelValue...).Inc()
}

func (g Gauge) Dec(labelValue ...string) {
	if !Enable() {
		return
	}
	g.vec.WithLabelValues(labelValue...).Dec()
}

func (g Gauge) Add(val float64, labelValue ...string) {
	if !Enable() {
		return
	}
	g.vec.WithLabelValues(labelValue...).Add(val)
}

func (g Gauge) Sub(val float64, labelValue ...string) {
	if !Enable() {
		return
	}
	g.vec.WithLabelValues(labelValue...).Sub(val)
}
