package yage

import (
	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/forestyc/playground/pkg/crawler"
	"github.com/forestyc/playground/pkg/log/zap"
)

type BalanceOfLawAndGospel struct {
	task string
	ctx  context.Context
}

func (blg *BalanceOfLawAndGospel) Init(ctx context.Context, config zap.Config, task string) {
	blg.ctx = ctx
	blg.ctx.C.Log = config
	// 初始化日志
	blg.ctx.Logger = zap.NewZap(blg.ctx.C.Log)
	blg.task = task
}

func (blg *BalanceOfLawAndGospel) Run() {
	c := crawler.NewColly(
		blg.task,
		"http://api.yageapp.com/api/web/share/postor.php?aid=19143&sid=702424&bundleid=&base_uid=1532937",
	)
	c.Run(crawler.WithCrawlCallback(blg.getList()))
}

func (blg *BalanceOfLawAndGospel) getList() crawler.Callback {
	return func() {

	}
}
