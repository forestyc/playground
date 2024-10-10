package main

import (
	"flag"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/entity/config"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/handler"
	"github.com/forestyc/playground/pkg/core/component"
)

func main() {
	var confFile string
	flag.StringVar(&confFile, "conf", "./etc/loan.yaml", "configuration")
	var c config.Config
	err := config.Load(confFile, &c)
	if err != nil {
		panic(err)
	}

	ctx := context.NewContext(c)

	component.Register(
		ctx.HttpServer.WithHandler(
			handler.NewLoanBasicInfo(ctx),
			handler.NewLoan(ctx),
		),
	)
	component.Serve()
	component.Close()
}
