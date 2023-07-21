package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Summary struct {
	vec *prometheus.SummaryVec
}

func NewSummary(name, help string, label ...string) *Summary {
	return poolSummary.Get(name, help, label...)
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

var (
	poolSummary = &summaryPool{}
)

type summaryPool struct {
	pool map[string]*Summary
	lock sync.Mutex
}

func (sp *summaryPool) Get(name, help string, label ...string) *Summary {
	sp.lock.Lock()
	defer sp.lock.Unlock()
	if sp.pool == nil {
		sp.pool = make(map[string]*Summary)
	}
	c, has := sp.pool[name]
	if has {
		return c
	} else {
		c = &Summary{
			vec: prometheus.NewSummaryVec(
				prometheus.SummaryOpts{
					Name: name,
					Help: help,
				}, label),
		}
		prometheus.MustRegister(c.vec)
		sp.pool[name] = c
		return c
	}
}
