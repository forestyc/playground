package gfex

import (
	"fmt"
	"github.com/Baal19905/playground/colly/pkg/crawler"
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
)

// PageInfo 页面信息
type PageInfo struct {
	crawler     crawler.Colly
	articleHref []string
}

// ListPage 要闻列表
type ListPage struct {
	currentPage int
	totalPage   int
	page        []*PageInfo
}

func NewListPage() *ListPage {
	l := &ListPage{
		currentPage: 1,
	}
	firstPage := &PageInfo{}
	firstPage.crawler = crawler.NewColly(
		host+"/gfex/bsyw/list_yw.shtml",
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
		fmt.Println(err)
	}
	for i := 1; i <= l.totalPage; i++ {
		page := &PageInfo{}
		url := fmt.Sprintf(pageUrlFormat, i)
		page.crawler = crawler.NewColly(
			url,
			page.getArticleHref(),
		)
		if err = page.crawler.Run(); err != nil {
			fmt.Println(err)
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
			if !strings.HasPrefix(e.Text, pageTotalPrefix) ||
				!strings.HasSuffix(e.Text, pageTotalSuffix) {
				//errMsg := `[广期所-本所要闻]总页码文本格式错误，应为"` + this.pageTotalPrefix + `x` + this.pageTotalSuffix + `", 实际为"` + pageTotalContent + `"`
				return
			}
			total := strings.TrimPrefix(e.Text, pageTotalPrefix)
			total = strings.TrimSuffix(total, pageTotalSuffix)
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
					url = host + url
					p.articleHref = append(p.articleHref, url)
				}
			})
		})
	}
}
