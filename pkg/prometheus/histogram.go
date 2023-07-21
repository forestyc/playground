package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Histogram struct {
	vec *prometheus.HistogramVec
}

func NewHistogram(name, help string, label ...string) *Histogram {
	return poolHistogram.Get(name, help, label...)
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

var (
	poolHistogram = &histogramPool{}
)

type histogramPool struct {
	pool map[string]*Histogram
	lock sync.Mutex
}

func (hp *histogramPool) Get(name, help string, label ...string) *Histogram {
	hp.lock.Lock()
	defer hp.lock.Unlock()
	if hp.pool == nil {
		hp.pool = make(map[string]*Histogram)
	}
	c, has := hp.pool[name]
	if has {
		return c
	} else {
		c = &Histogram{
			vec: prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Name: name,
					Help: help,
				}, label),
		}
		prometheus.MustRegister(c.vec)
		hp.pool[name] = c
		return c
	}
}
