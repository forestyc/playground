package news

import (
	"fmt"
	"strconv"
	"strings"

	_const "github.com/forestyc/playground/cmd/crawler/app/common"

	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/forestyc/playground/pkg/crawler"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

// PageInfo 页面信息
type PageInfo struct {
	crawler     crawler.Colly
	articleHref []string
	ctx         context.Context
}

// ListPage 要闻列表
type ListPage struct {
	currentPage int
	totalPage   int
	page        []*PageInfo
	ctx         context.Context
	task        string
}

func NewListPage(ctx context.Context, task string) *ListPage {
	l := &ListPage{
		currentPage: 1,
		ctx:         ctx,
		task:        task,
	}
	firstPage := &PageInfo{ctx: ctx}
	firstPage.crawler = crawler.NewColly(
		task,
		_const.NewsPageUrlFirst,
		nil,
		crawler.WithCrawlCallback(firstPage.getTotalPage(&l.totalPage)),
		crawler.WithCrawlCallback(firstPage.getArticleHref()),
	)
	l.page = append(l.page, firstPage)
	return l
}

func (l *ListPage) Run() {
	var err error
	// 首页
	if err = l.page[0].crawler.Run(); err != nil {
		l.ctx.Logger.Error("[本所要闻]爬取列表页失败", zap.Error(err), zap.String("url", l.page[0].crawler.Url))
		return
	}
	l.ctx.Logger.Info("[本所要闻]爬取列表页成功", zap.String("url", l.page[0].crawler.Url))
	for i := 1; i <= l.totalPage; i++ {
		page := &PageInfo{ctx: l.ctx}
		url := fmt.Sprintf(_const.NewsPageUrlFormat, i)
		page.crawler = crawler.NewColly(
			l.task,
			url,
			nil,
			crawler.WithCrawlCallback(page.getArticleHref()),
		)
		if err = page.crawler.Run(); err != nil {
			l.ctx.Logger.Error("[本所要闻]爬取列表页失败", zap.Error(err), zap.String("url", page.crawler.Url))
			continue
		}
		l.ctx.Logger.Info("[本所要闻]爬取列表页成功", zap.String("url", page.crawler.Url))
		l.page = append(l.page, page)
	}
}

func (l *ListPage) GetArticleHref() []string {
	var href []string
	for _, e := range l.page {
		href = append(href, e.articleHref...)
	}
	return href
}

// 获取总页数
func (p *PageInfo) getTotalPage(totalPage *int) crawler.Callback {
	return func() {
		selector := "body > div.mainBox.clearfix > div.container.listBox > div.pageList.newsList.news-list-yw > ul>script:nth-of-type(2)"
		p.crawler.Crawler.OnHTML(selector, func(e *colly.HTMLElement) {
			if !strings.HasPrefix(e.Text, _const.PageTotalPrefix) ||
				!strings.Contains(e.Text, _const.PageTotalSuffix) {
				errMsg := `[本所要闻]总页码文本格式错误，应为"` + _const.PageTotalPrefix + `x` + _const.PageTotalSuffix + `", 实际为"` + e.Text + `"`
				p.ctx.Logger.Error("[本所要闻]爬取列表页失败", zap.String("url", errMsg))
				return
			}
			total := strings.TrimPrefix(e.Text, _const.PageTotalPrefix)
			totalSlice := strings.Split(total, ",")
			total = totalSlice[0]
			*totalPage, _ = strconv.Atoi(total)
		})
	}
}

// 获取文章href
func (p *PageInfo) getArticleHref() crawler.Callback {
	return func() {
		selector := `div[class='pageList newsList news-list-yw'] ul li`
		p.crawler.Crawler.OnHTML(selector, func(e *colly.HTMLElement) {
			selector = "div[class='clearfix'] > div[class='item_fm imgScale'] a"
			e.ForEach(selector, func(i int, element *colly.HTMLElement) {
				url := element.Attr("href")
				if strings.HasPrefix(url, "/") {
					url = _const.Host + url
					p.articleHref = append(p.articleHref, url)
				}
			})
		})
	}
}

// 错误
func (p *PageInfo) err() crawler.Callback {
	return func() {
		p.crawler.Crawler.OnError(func(r *colly.Response, e error) {
			p.ctx.Logger.Error("[本所要闻]爬取列表页失败", zap.Error(e), zap.Any("rsp", r))
		})
	}
}
