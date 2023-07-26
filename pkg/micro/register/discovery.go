package register

import "sync"

type Discovery interface {
	Name() string
	Roger(host string, status int)
}

type ServiceStatus interface {
	Attach(discovery Discovery)
	Detach(discovery Discovery)
	Notify(status int)
}

type Service struct {
	hosts []string
	name  string
	sync.RWMutex
}

func NewService(name string) *Service {
	return &Service{
		name: name,
	}
}

func (s *Service) Name() string {
	return s.name
}

func (s *Service) Roger(host string, status int) {
	s.Lock()
	defer s.Unlock()
	switch status {
	case Up:
		s.hosts = append(s.hosts, host)
	case Down:
		for i, e := range s.hosts {
			if e == host {
				tail := s.hosts[i+1:]
				s.hosts = s.hosts[0:i]
				s.hosts = append(s.hosts, tail...)
			}
		}
	}
}

type KeyStatus struct {
	Observer map[string]Discovery
}

func (ks *KeyStatus) Attach(discovery Discovery) {
	ks.Observer[discovery.Name()] = discovery
}
func (ks *KeyStatus) Detach(discovery Discovery) {
	delete(ks.Observer, discovery.Name())
}
func (ks *KeyStatus) Notify(key string, status int) {

}
