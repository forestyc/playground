package registercenter

// Publisher 服务发布
type Publisher interface {
	Register(endpoints []string, service Service, ttl int64) error
	Close() error
}

// Subscriber 服务订阅
type Subscriber interface {
	Listen(endpoints []string, serviceName string) error
	GetService(name string) (Service, bool)
	Close() error
}

type Service struct {
	Name string
	Url  string
}
