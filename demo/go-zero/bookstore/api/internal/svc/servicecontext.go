package svc

import (
	"github.com/forestyc/playground/go-zero/bookstore/api/internal/config"
	"github.com/forestyc/playground/go-zero/bookstore/rpc/add/adder"
	"github.com/forestyc/playground/go-zero/bookstore/rpc/check/checker"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder     // add rpc
	Checker checker.Checker // check rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),       // add rpc
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)), // check rpc
	}
}
