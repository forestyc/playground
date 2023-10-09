package logic

import (
	"path"
	"time"

	"github.com/forestyc/playground/cmd/crawler/app/common"
	"github.com/forestyc/playground/cmd/crawler/app/logic/gfex"

	"github.com/forestyc/playground/cmd/crawler/app/context"
)

var (
	crawler map[string]Crawler
	ctx     context.Context
)

// Register 注册
func Register(c context.Context) {
	ctx = c
	crawler = map[string]Crawler{
		common.GfexNews:         &gfex.News{},         // 广期所-本所要闻
		common.GfexAnnouncement: &gfex.Announcement{}, // 广期所-通知公告
		common.GfexFocus:        &gfex.Focus{},        // 广期所-媒体聚焦
	}
}

// Run 执行爬虫任务
func Run(labels []string) {
	if ctx.C.Prometheus.HasPrometheus() {
		ctx.C.Prometheus.Start()
	}
	for _, label := range labels {
		job, ok := crawler[label]
		if ok {
			log := ctx.C.Log
			log.Prefix = label
			log.Director = path.Join(log.Director, log.Prefix)
			log.LinkName = path.Join(log.Director, "last_log")
			job.Init(ctx, log, label)
			go job.Run()
		}
	}
	time.Sleep(time.Minute) // 防止prometheus遗漏数据
}
