package registercenter

import (
	"math/rand"
	"sync"
	"time"
)

type Service struct {
	Name string
	Url  string
}

type ServiceGroup struct {
	Name  string
	Group []*Service
}

type ServiceCache struct {
	Cache map[string]*ServiceGroup
	Lock  sync.RWMutex
}

func NewServiceCache() ServiceCache {
	return ServiceCache{
		Cache: make(map[string]*ServiceGroup),
	}
}

func (sc *ServiceCache) SetService(name, url string) {
	sc.Lock.Lock()
	defer sc.Lock.Unlock()
	group, ok := sc.Cache[name]
	if ok {
		group.SetService(name, url)
	} else {
		group = NewServiceGroup(name)
		group.SetService(name, url)
	}
}

func (sc *ServiceCache) GetService(name string) string {
	sc.Lock.RLock()
	defer sc.Lock.RUnlock()
	group, ok := sc.Cache[name]
	if ok {
		return group.GetService()
	}
	return ""
}

func NewServiceGroup(name string) *ServiceGroup {
	return &ServiceGroup{
		Name:  name,
		Group: make([]*Service, 0, 3),
	}
}

func (g *ServiceGroup) GetService() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(g.Group)
	if l > 0 {
		idx := r.Intn(l)
		service := g.Group[idx]
		return service.Url
	} else {
		return ""
	}
}

func (g *ServiceGroup) SetService(name, url string) {
	exist := false
	for _, e := range g.Group {
		if e.Name == name {
			e.Url = url
			exist = true
			break
		}
	}
	if !exist {
		g.Group = append(g.Group, &Service{
			Name: name,
			Url:  url,
		})
	}
}
