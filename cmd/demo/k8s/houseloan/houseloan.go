package main

import (
	"flag"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/entity/config"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/handler"
	"github.com/forestyc/playground/pkg/core/component"
)

func main() {
	var confFile string
	flag.StringVar(&confFile, "conf", "./etc/houseloan.yaml", "configuration")
	var c config.Config
	err := config.Load(confFile, &c)
	if err != nil {
		panic(err)
	}

	ctx := context.NewContext(c)

	component.Register(
		ctx.HttpServer.WithHandler(
			handler.NewPrincipalInterest(
				"2024-01-15", 360,
				600000.00, 0.04,
				800000.00, 0.03575,
			),
		),
	)
	component.Serve()
	component.Close()
}
