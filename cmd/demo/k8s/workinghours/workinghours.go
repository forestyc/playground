package main

import (
	"flag"
	"github.com/forestyc/playground/cmd/demo/k8s/workinghours/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/workinghours/app/model/config"
)

func main() {
	var confFile string
	flag.StringVar(&confFile, "conf", "./etc/workinghours.yaml", "configuration")
	var c config.Config
	err := config.Load(confFile, &c)
	if err != nil {
		panic(err)
	}
	ctx, err := context.NewContext(c)
	if err != nil {
		panic(err)
	}
	ctx.Run()
}
