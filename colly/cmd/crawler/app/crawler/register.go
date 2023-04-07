package crawler

import "C"
import (
	"github.com/Baal19905/playground/colly/cmd/crawler/app/context"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/crawler/gfex/news"
	"path"
)

var (
	crawler map[string]Crawler
	ctx     context.GlobalContext
)

// Register 注册
func Register(ctx context.GlobalContext) {
	crawler = map[string]Crawler{
		"gfex-news": &news.GfexNews{}, // 广期所-本所要闻
	}
}

// Run 执行爬虫任务
func Run(labels []string) {
	for _, label := range labels {
		job, ok := crawler[label]
		if ok {
			ctx.C.Log.Prefix = label
			ctx.C.Log.Director = path.Join(ctx.C.Log.Director, ctx.C.Log.Prefix)
			ctx.C.Log.LinkName = path.Join(ctx.C.Log.Director, "last_log")
			job.Init(ctx)
			go job.Run()
		}
	}
}
