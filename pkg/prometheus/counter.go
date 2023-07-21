package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Counter struct {
	vec *prometheus.CounterVec
}

func NewCounter(name, help string, label ...string) *Counter {
	return poolCounter.Get(name, help, label...)
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

var (
	poolCounter = &counterPool{}
)

type counterPool struct {
	pool map[string]*Counter
	lock sync.Mutex
}

func (cp *counterPool) Get(name, help string, label ...string) *Counter {
	cp.lock.Lock()
	defer cp.lock.Unlock()
	if cp.pool == nil {
		cp.pool = make(map[string]*Counter)
	}
	c, has := cp.pool[name]
	if has {
		return c
	} else {
		c = &Counter{
			vec: prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: name,
					Help: help,
				}, label),
		}
		prometheus.MustRegister(c.vec)
		cp.pool[name] = c
		return c
	}
}
