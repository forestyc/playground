package registercenter

// register
// look up

type RegisterCenter interface {
	Register(name, url string) error
	GetService(name string) string
	Close()
}
