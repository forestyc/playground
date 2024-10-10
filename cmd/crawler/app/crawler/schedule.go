package crawler

import (
	"path"

	"github.com/forestyc/playground/cmd/crawler/app/common"
	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/forestyc/playground/cmd/crawler/app/handler/dce"
	"github.com/forestyc/playground/cmd/crawler/app/handler/yage"
	"github.com/forestyc/playground/pkg/core/log/zap"
)

var (
	crawler map[string]Crawler
	ctx     context.Context
)

// Register 注册
func Register(c context.Context) {
	ctx = c
	crawler = map[string]Crawler{
		common.Yage:            &yage.BalanceOfLawAndGospel{},
		common.DceNews:         &dce.News{},
		common.DceAnnouncement: &dce.Announcement{},
		common.DceVariety:      &dce.Variety{},
	}
}

// Run 执行爬虫任务
func Run(labels []string) {
	for _, label := range labels {
		job, ok := crawler[label]
		if ok {
			// 重新初始化日志
			ctxTask := ctx
			ctxTask.C.Log.Director = path.Join(ctxTask.C.Log.Director, label)
			ctxTask.C.Log.LinkName = path.Join(ctxTask.C.Log.Director, label)
			ctxTask.Logger = zap.NewZap(ctxTask.C.Log)
			job.Init(ctxTask, label)
			ctx.Wg.Add(1)
			go job.Run()
		}
	}
	ctx.Wg.Wait()
}
