package gfex

import (
	"github.com/forestyc/playground/cmd/crawler/app/common"
	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/forestyc/playground/cmd/crawler/app/handler/gfex/news"
	"github.com/forestyc/playground/pkg/log/zap"
	rawZap "go.uber.org/zap"
)

type News struct {
	ctx     context.Context
	species string
	task    string
}

// Init 初始化
func (gn *News) Init(ctx context.Context, config zap.Config, task string) {
	gn.task = task
	gn.ctx = ctx
	gn.ctx.C.Log = config
	// 初始化日志
	gn.ctx.Logger = zap.NewZap(gn.ctx.C.Log)
}

// Run 广期所-本所要闻
func (gn *News) Run() {
	// 爬取列表页
	href := gn.CrawlPage()
	// 爬取文章
	for _, e := range href {
		if !common.Skip(gn.ctx, e) {
			gn.CrawlArticle(e)
			common.Record(gn.ctx, e)
		} else {
			gn.ctx.Logger.Info("[本所要闻]跳过", rawZap.String("已爬取", e))
		}
	}
	// 标记完成
	gn.ctx.Wg.Done()
}

// CrawlPage 获取所有文章
func (gn *News) CrawlPage() []string {
	gn.ctx.Logger.Info("[本所要闻]列表页爬取开始")
	defer gn.ctx.Logger.Info("[本所要闻]列表页爬取结束")
	listPage := news.NewListPage(gn.ctx, gn.task)
	listPage.Run()
	return listPage.GetArticleHref()
}

// CrawlArticle 爬取文章
func (gn *News) CrawlArticle(url string) *news.Article {
	gn.ctx.Logger.Info("[本所要闻]文章爬取开始")
	defer gn.ctx.Logger.Info("[本所要闻]文章爬取结束")
	n := news.NewNews(gn.ctx, gn.task, url)
	n.Run()
	return n
}
