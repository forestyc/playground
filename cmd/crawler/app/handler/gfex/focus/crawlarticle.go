package focus

import (
	"strings"
	"time"

	"github.com/forestyc/playground/cmd/crawler/app/common"
	"github.com/forestyc/playground/cmd/crawler/app/dao"
	"github.com/forestyc/playground/cmd/crawler/app/util"

	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/forestyc/playground/pkg/crawler"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

// Article 文章
type Article struct {
	crawler crawler.Colly
	ctx     context.Context
	url     string
	dao.Article
}

func NewNews(ctx context.Context, task, url string) *Article {
	a := &Article{
		ctx: ctx,
		url: url,
	}
	// 品种
	variety, err := common.GetVariety(ctx, []int{common.Exchange})
	if err != nil {
		a.ctx.Logger.Warn("[通知公告]获取广期所品种失败", zap.Error(err))
	}
	a.Article.Species = "[\"" + strings.Join(variety, `","`) + "\"]"
	a.crawler = crawler.NewColly(
		task,
		url,
		crawler.WithPipeline(a.Pipeline()),
		crawler.WithCrawlCallback(a.getOrigin()),
		crawler.WithCrawlCallback(a.getPublishDate()),
		crawler.WithCrawlCallback(a.getTittle()),
		crawler.WithCrawlCallback(a.getContent()),
		crawler.WithCrawlCallback(a.setConst()),
	)
	return a
}

func (a *Article) Run() {
	if err := a.crawler.Run(); err != nil {
		a.ctx.Logger.Error("[媒体聚焦]爬取文章失败", zap.Error(err), zap.Any("url", a.url))
		return
	}
	if len(a.Title) == 0 || len(a.PublishDate) == 0 || len(a.Body) == 0 {
		a.ctx.Logger.Error("非法的文章", zap.Any("article", a), zap.Any("url", a.url))
	}
	a.ctx.Logger.Info("[媒体聚焦]爬取文章成功", zap.Any("url", a.url))
}

func (a *Article) Pipeline() crawler.Pipeline {
	return func() error {
		a.ctx.Logger.Info("[通知公告]文章页保存开始")
		if err := a.Create(a.ctx); err != nil {
			a.ctx.Logger.Error("[通知公告]文章保存失败", zap.Error(err))
		}
		a.ctx.Logger.Info("[通知公告]文章保存结束")
		return nil
	}
}

// 获取来源
func (a *Article) getOrigin() crawler.Callback {
	return func() {
		selector := "body > div.mainBox.clearfix > div > div > div.InfoTitle > div > span:nth-child(1) > font"
		a.crawler.Crawler.OnHTML(selector, func(e *colly.HTMLElement) {
			// 发布日期
			a.Origin = e.Text
			if len(a.Origin) == 0 {
				a.Origin = common.Origin
			}
		})
	}
}

// 获取发布日期
func (a *Article) getPublishDate() crawler.Callback {
	return func() {
		selector := "body > div.mainBox.clearfix > div > div > div.InfoTitle > div > span:nth-child(2) > font > publishtime"
		a.crawler.Crawler.OnHTML(selector, func(e *colly.HTMLElement) {
			// 发布日期
			a.PublishDate = e.Text
			// 排序日期
			a.SortDate = a.PublishDate
			if len(a.SortDate) == 0 {
				a.SortDate = time.Now().Format("2006-01-02 15:04:05") // sort_date必填
			}
		})
	}
}

// 获取标题
func (a *Article) getTittle() crawler.Callback {
	return func() {
		selector := "body > div.mainBox.clearfix > div > div > div.InfoTitle > h1"
		a.crawler.Crawler.OnHTML(selector, func(e *colly.HTMLElement) {
			a.Title = strings.TrimSpace(e.Text)
		})
	}
}

// 获取正文
func (a *Article) getContent() crawler.Callback {
	return func() {
		selector := "body > div.mainBox.clearfix > div > div > div.InfoContent > ucapcontent"
		a.crawler.Crawler.OnHTML(selector, func(e *colly.HTMLElement) {
			content, _ := e.DOM.Html()
			content = strings.TrimSpace(content)
			a.Body = util.ReplaceNBSPinHtml(content)
			host := e.Request.URL.String()
			idx := strings.LastIndex(host, "/")
			host = host[0 : idx+1]
			a.Body = util.AddHost(a.Body, host)
		})
	}
}

// 常量赋值
func (a *Article) setConst() crawler.Callback {
	return func() {
		// 初始栏目
		a.ColumnLevel = common.Column
	}
}

// 错误
func (a *Article) err() crawler.Callback {
	return func() {
		a.crawler.Crawler.OnError(func(r *colly.Response, e error) {
			a.ctx.Logger.Error("[媒体聚焦]爬取文章失败", zap.Error(e), zap.Any("rsp", r))
		})
	}
}
