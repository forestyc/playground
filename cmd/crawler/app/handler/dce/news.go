package dce

import (
	"fmt"
	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/forestyc/playground/cmd/crawler/app/model/fip"
	"github.com/forestyc/playground/cmd/crawler/app/util"
	"github.com/forestyc/playground/pkg/crawler"
	"github.com/gocolly/colly/v2"
	"github.com/xuri/excelize/v2"
	rawZap "go.uber.org/zap"
	"strconv"
	"strings"
)

var (
	newsUrlPrefix = []string{
		"http://www.dce.com.cn/dalianshangpin/xwzx93/jysxw/13363-",
		"http://www.dce.com.cn/dalianshangpin/xwzx93/mtkdss/13365-",
		"http://www.dce.com.cn/dalianshangpin/xwzx93/scyjjydf/446f0284-",
	}
	newsPageMax = 15
)

type News struct {
	task     string
	ctx      context.Context
	articles []fip.Article
	crawlers []*crawler.Colly
	id       int
}

func (n *News) Init(ctx context.Context, task string) {
	n.ctx = ctx
	n.task = task
	for _, prefix := range newsUrlPrefix {
		for i := 1; i <= newsPageMax; i++ {
			url := prefix + strconv.Itoa(i) + ".html"
			c := crawler.NewColly(
				n.task,
				url,
			)
			n.crawlers = append(n.crawlers, c)
		}
	}
}

func (n *News) Run() {
	for _, c := range n.crawlers {
		if err := c.Run(crawler.WithCrawlCallback(n.callback(c))); err != nil {
			n.ctx.Logger.Error("爬取文章失败", rawZap.Error(err), rawZap.String("url", c.Url))
		}
	}
	n.toXslx()
	n.ctx.Wg.Done()
	n.ctx.Logger.Info("Run success", rawZap.String("task", n.task))
}

func (n *News) callback(c *crawler.Colly) crawler.Callback {
	return func() {
		c.Crawler.OnHTML(`body > div.container_w > div > div:nth-child(2)`, func(ul *colly.HTMLElement) {
			ul.ForEach(`li > a`, func(i int, a *colly.HTMLElement) {
				uri := a.Attr("href")
				url := util.Host(ul) + uri
				cArticle := crawler.NewColly(
					n.task,
					url,
				)
				if err := cArticle.Run(crawler.WithCrawlCallback(n.getArticle(cArticle))); err != nil {
					n.ctx.Logger.Error("爬取文章失败", rawZap.Error(err), rawZap.String("url", c.Url))
				}
			})

		})

	}
}

func (n *News) getArticle(c *crawler.Colly) crawler.Callback {
	return func() {
		selector := `body > div.container_w > div > div > div:nth-child(2) > div.detail_inner`
		c.Crawler.OnHTML(selector, func(e *colly.HTMLElement) {
			n.id += 1
			article := fip.Article{}
			article.Id = n.id
			selectorTitle := `div.tit_header > h2`
			article.Title = e.DOM.Find(selectorTitle).Text()
			selectorDate := `div.tit_header > p`
			article.Date = strings.TrimLeft(e.DOM.Find(selectorDate).Contents().Not(`span`).Text(), "发布时间：")
			selectorSource := `div.tit_header > p > span > span`
			article.Source = e.DOM.Find(selectorSource).Text()
			selectorContent := `#zoom`
			article.Content, _ = e.DOM.Find(selectorContent).Html()
			article.Content = util.FormatArticle(article.Content, e)
			n.articles = append(n.articles, article)
			n.ctx.Logger.Info("爬取成功", rawZap.Int("Id", article.Id), rawZap.String("title", article.Title), rawZap.String("url", c.Url))
		})
	}
}

func (n *News) toXslx() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// add title
	f.SetCellValue("Sheet1", "A1", "序号")
	f.SetCellValue("Sheet1", "B1", "标题名称")
	f.SetCellValue("Sheet1", "C1", "内容正文")
	f.SetCellValue("Sheet1", "D1", "文件来源")
	f.SetCellValue("Sheet1", "E1", "摘要")
	f.SetCellValue("Sheet1", "F1", "发布日期")

	// add line
	for line, article := range n.articles {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(line+2), article.Id)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(line+2), article.Title)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(line+2), article.Content)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(line+2), article.Source)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(line+2), article.Summary)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(line+2), article.Date)
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("news.xlsx"); err != nil {
		fmt.Println(err)
	}
}
