package logic

import (
	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/forestyc/playground/pkg/log/zap"
)

type Crawler interface {
	Init(ctx context.Context, config zap.Config, task string)
	Run()
}
