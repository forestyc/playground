package gfex

import (
	"github.com/forestyc/playground/cmd/crawler/app/common"
	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/forestyc/playground/cmd/crawler/app/handler/gfex/announcement"
	"github.com/forestyc/playground/pkg/log/zap"
	rawZap "go.uber.org/zap"
)

type Announcement struct {
	ctx     context.Context
	species string
	task    string
}

// Init 初始化
func (gn *Announcement) Init(ctx context.Context, config zap.Config, task string) {
	gn.ctx = ctx
	gn.ctx.C.Log = config
	// 初始化日志
	gn.ctx.Logger = zap.NewZap(gn.ctx.C.Log)
	gn.task = task
}

// Run 广期所-通知公告
func (gn *Announcement) Run() {
	// 爬取列表页
	href := gn.CrawlPage()
	// 爬取文章
	for _, e := range href {
		if !common.Skip(gn.ctx, e) {
			gn.CrawlArticle(e)
			common.Record(gn.ctx, e)
		} else {
			gn.ctx.Logger.Info("[通知公告]跳过", rawZap.String("已爬取", e))
		}
	}
	// 标记完成
	gn.ctx.Wg.Done()
}

// CrawlPage 获取所有文章
func (gn *Announcement) CrawlPage() []string {
	gn.ctx.Logger.Info("[通知公告]列表页爬取开始")
	defer gn.ctx.Logger.Info("[通知公告]列表页爬取结束")
	listPage := announcement.NewListPage(gn.ctx, gn.task)
	listPage.Run()
	return listPage.GetArticleHref()
}

// CrawlArticle 爬取文章
func (gn *Announcement) CrawlArticle(url string) *announcement.Article {
	gn.ctx.Logger.Info("[通知公告]文章爬取开始")
	defer gn.ctx.Logger.Info("[通知公告]文章爬取结束")
	news := announcement.NewNews(gn.ctx, gn.task, url)
	news.Run()
	return news
}
