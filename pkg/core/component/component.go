package component

import (
	"os"
	"os/signal"
	"syscall"
)

type Component interface {
	Serve()
	Close()
}

var (
	components []Component
)

func Register(component ...Component) {
	components = append(components, component...)
}

func Serve() {
	for _, component := range components {
		component.Serve()
	}
}

func Close() {
	quit := make(chan os.Signal)
	// kill (no param) default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	for _, component := range components {
		component.Close()
	}
}
