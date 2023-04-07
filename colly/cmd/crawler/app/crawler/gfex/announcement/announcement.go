package announcement

import "C"
import (
	"github.com/Baal19905/playground/colly/cmd/crawler/app/context"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/crawler/common"
	"github.com/Baal19905/playground/colly/pkg/log/zap"
	rawZap "go.uber.org/zap"
)

type GfexAnnouncement struct {
	ctx context.GlobalContext
}

// Init 初始化
func (gn *GfexAnnouncement) Init(ctx context.GlobalContext, config zap.Config) {
	gn.ctx = ctx
	gn.ctx.C.Log = config
	// 初始化日志
	gn.ctx.Logger = zap.NewZap(gn.ctx.C.Log)
}

// Run 广期所-通知公告
func (gn *GfexAnnouncement) Run() {
	// 爬取列表页
	href := gn.CrawlPage()
	// 爬取文章
	var articles []*Article
	for _, e := range href {
		var a *Article
		if !common.Skip(gn.ctx, e) {
			a = gn.CrawlArticle(e)
			common.Record(gn.ctx, e)
			articles = append(articles, a)
		}
	}
	// 写入数据库
	gn.SaveArticle(articles)
	// 标记完成
	gn.ctx.Wg.Done()
}

// CrawlPage 获取所有文章
func (gn *GfexAnnouncement) CrawlPage() []string {
	gn.ctx.Logger.Info("[通知公告]列表页爬取开始")
	defer gn.ctx.Logger.Info("[通知公告]列表页爬取结束")
	listPage := NewListPage(gn.ctx)
	listPage.Run()
	return listPage.GetArticleHref()
}

// CrawlArticle 爬取文章
func (gn *GfexAnnouncement) CrawlArticle(url string) *Article {
	gn.ctx.Logger.Info("[通知公告]文章爬取开始")
	defer gn.ctx.Logger.Info("[通知公告]文章爬取结束")
	news := NewNews(gn.ctx, url)
	news.Run()
	return news
}

// SaveArticle 保存文章
func (gn *GfexAnnouncement) SaveArticle(articles []*Article) {
	if articles == nil || len(articles) == 0 {
		return
	}
	gn.ctx.Logger.Info("[通知公告]文章页保存开始")
	session := gn.ctx.Db.Session()
	session = session.Begin()
	session = session.Table(gn.ctx.C.Crawler.Table).Save(articles)
	if int(session.RowsAffected) != len(articles) {
		gn.ctx.Logger.Warn("[通知公告]文章页保存失败")
	}
	session.Commit()
	gn.ctx.Logger.Info("[通知公告]文章保存结束", rawZap.Int64("RowsAffected", session.RowsAffected))
}
