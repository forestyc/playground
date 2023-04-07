package crawler

import "C"
import (
	"github.com/Baal19905/playground/colly/cmd/crawler/app/context"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/crawler/gfex/announcement"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/crawler/gfex/focus"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/crawler/gfex/news"
	"path"
)

var (
	crawler map[string]Crawler
	ctx     context.GlobalContext
)

// Register 注册
func Register(c context.GlobalContext) {
	ctx = c
	crawler = map[string]Crawler{
		"gfex-news":         &news.GfexNews{},                 // 广期所-本所要闻
		"gfex-announcement": &announcement.GfexAnnouncement{}, // 广期所-通知公告
		"gfex-focus":        &focus.GfexFocus{},               // 广期所-媒体聚焦
	}
}

// Run 执行爬虫任务
func Run(labels []string) {
	for _, label := range labels {
		job, ok := crawler[label]
		if ok {
			log := ctx.C.Log
			log.Prefix = label
			log.Director = path.Join(log.Director, log.Prefix)
			log.LinkName = path.Join(log.Director, "last_log")
			job.Init(ctx, log)
			go job.Run()
		}
	}
}
