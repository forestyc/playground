package announcement

import (
	"github.com/Baal19905/playground/colly/cmd/crawler/app/context"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/crawler/gfex/const"
	"github.com/Baal19905/playground/colly/pkg/crawler"
	"github.com/Baal19905/playground/colly/pkg/util"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"strings"
	"time"
)

// Article 文章
type Article struct {
	PublishDate string                `gorm:"publish_date"`
	SortDate    string                `gorm:"sort_date"`
	Title       string                `gorm:"title"`
	Origin      string                `gorm:"origin"`
	Body        string                `gorm:"body"`
	ColumnLevel string                `gorm:"column_level"`
	crawler     crawler.Colly         `gorm:"-"`
	ctx         context.GlobalContext `gorm:"-"`
	url         string                `gorm:"-"`
}

func NewNews(ctx context.GlobalContext, url string) *Article {
	a := &Article{
		ctx: ctx,
		url: url,
	}
	a.crawler = crawler.NewColly(
		url,
		a.getPublishDate(),
		a.getTittle(),
		a.getContent(),
		a.setConst(),
	)
	return a
}

func (a *Article) Run() {
	if err := a.crawler.Run(); err != nil {
		a.ctx.Logger.Error("[通知公告]爬取文章失败", zap.Error(err), zap.Any("url", a.url))
		return
	}
	if len(a.Title) == 0 || len(a.PublishDate) == 0 || len(a.Body) == 0 {
		a.ctx.Logger.Error("非法的文章", zap.Any("article", a), zap.Any("url", a.url))
	}
	a.ctx.Logger.Info("[通知公告]爬取文章成功", zap.Any("url", a.url))
}

// 获取发布日期
func (a *Article) getPublishDate() crawler.Callback {
	return func() {
		selector := "body > div.mainBox.clearfix > div > div > div.InfoTitle > div > span:nth-child(1) > font > publishtime"
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
		})
	}
}

// 常量赋值
func (a *Article) setConst() crawler.Callback {
	return func() {
		// 来源
		a.Origin = _const.Origin
		// 初始栏目
		a.ColumnLevel = _const.Column
	}
}

// 错误
func (a *Article) err() crawler.Callback {
	return func() {
		a.crawler.Crawler.OnError(func(r *colly.Response, e error) {
			a.ctx.Logger.Error("[通知公告]爬取文章失败", zap.Error(e), zap.Any("rsp", r))
		})
	}
}
