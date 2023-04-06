package gfex

import (
	"fmt"
	"github.com/Baal19905/playground/colly/pkg/crawler"
	"github.com/Baal19905/playground/colly/pkg/util"
	"github.com/gocolly/colly/v2"
	"strings"
	"time"
)

func NewNews(url string) *Article {
	a := &Article{}
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
		fmt.Println(err)
	}
	if len(a.Title) == 0 || len(a.PublishDate) == 0 || len(a.Body) == 0 {
		fmt.Println("nil article")
	}
}

// 获取总页数
func (a *Article) getPublishDate() crawler.Callback {
	return func() {
		selector := "body > div.mainBox.clearfix > div > div > div.InfoTitle > div > span:nth-child(1) > font > publishtime"
		a.crawler.Crawler.OnHTML(selector, func(e *colly.HTMLElement) {
			a.PublishDate = e.Text
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
		// 排序日期
		a.SortDate = a.PublishDate
		if len(a.SortDate) == 0 {
			a.SortDate = time.Now().Format("2006-01-02 15:04:05") // sort_date必填
		}
		// 来源
		a.Origin = origin
		// 初始栏目
		a.ColumnLevel = column
		a.RefColumns = column
	}
}
