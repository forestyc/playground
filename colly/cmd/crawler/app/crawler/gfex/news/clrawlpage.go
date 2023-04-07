package news

import (
	"fmt"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/context"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/crawler/gfex/common"
	"github.com/Baal19905/playground/colly/pkg/crawler"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

// PageInfo 页面信息
type PageInfo struct {
	crawler     crawler.Colly
	articleHref []string
	ctx         context.GlobalContext
}

// ListPage 要闻列表
type ListPage struct {
	currentPage int
	totalPage   int
	page        []*PageInfo
	ctx         context.GlobalContext
}

func NewListPage(ctx context.GlobalContext) *ListPage {
	l := &ListPage{
		currentPage: 1,
		ctx:         ctx,
	}
	firstPage := &PageInfo{ctx: ctx}
	firstPage.crawler = crawler.NewColly(
		common.Host+"/gfex/bsyw/list_yw.shtml",
		firstPage.getEveryPage(&l.totalPage),
		firstPage.getArticleHref(),
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
	for i := 1; i <= l.totalPage; i++ {
		page := &PageInfo{ctx: l.ctx}
		url := fmt.Sprintf(common.PageUrlFormat, i)
		page.crawler = crawler.NewColly(
			url,
			page.getArticleHref(),
		)
		if err = page.crawler.Run(); err != nil {
			l.ctx.Logger.Error("[本所要闻]爬取列表页失败", zap.Error(err), zap.String("url", page.crawler.Url))
			continue
		}
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
func (p *PageInfo) getEveryPage(totalPage *int) crawler.Callback {
	return func() {
		selector := "body > div.mainBox.clearfix > div.container.listBox > div.pageList.newsList.news-list-yw > ul>script:nth-of-type(2)"
		p.crawler.Crawler.OnHTML(selector, func(e *colly.HTMLElement) {
			if !strings.HasPrefix(e.Text, common.PageTotalPrefix) ||
				!strings.HasSuffix(e.Text, common.PageTotalSuffix) {
				errMsg := `[本所要闻]总页码文本格式错误，应为"` + common.PageTotalPrefix + `x` + common.PageTotalSuffix + `", 实际为"` + e.Text + `"`
				p.ctx.Logger.Error("[本所要闻]爬取列表页失败", zap.String("url", errMsg))
				return
			}
			total := strings.TrimPrefix(e.Text, common.PageTotalPrefix)
			total = strings.TrimSuffix(total, common.PageTotalSuffix)
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
					url = common.Host + url
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
