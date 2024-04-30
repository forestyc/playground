package main

import (
	"flag"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/handler"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/model/config"
	"github.com/forestyc/playground/pkg/component"
)

func main() {
	var confFile string
	flag.StringVar(&confFile, "conf", "./etc/houseloan.yaml", "configuration")
	var c config.Config
	err := config.Load(confFile, &c)
	if err != nil {
		panic(err)
	}
	ctx, err := context.NewContext(c)
	if err != nil {
		panic(err)
	}

	handler.RegisterPrincipalInterestRouters(ctx)

	component.Register(ctx.HttpServer)
	component.Serve()
	component.Close()
}
