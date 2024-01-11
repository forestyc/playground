package dce

import (
	"fmt"
	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/forestyc/playground/cmd/crawler/app/model/fip"
	"github.com/forestyc/playground/cmd/crawler/app/util"
	"github.com/forestyc/playground/pkg/crawler"
	"github.com/forestyc/playground/pkg/log/zap"
	"github.com/gocolly/colly/v2"
	"github.com/xuri/excelize/v2"
	rawZap "go.uber.org/zap"
	"strconv"
)

var (
	announcementUrlPrefix = []string{
		"http://www.dce.com.cn/dalianshangpin/ywfw/jystz/ywtz/13305-",
		"http://www.dce.com.cn/dalianshangpin/ywfw/jystz/hdtz/13303-",
	}
	announcementPageMax = 5
)

type Announcement struct {
	task     string
	ctx      context.Context
	articles []fip.Article
	crawlers []*crawler.Colly
	id       int
}

func (an *Announcement) Init(ctx context.Context, config zap.Config, task string) {
	an.ctx = ctx
	an.ctx.C.Log = config
	// 初始化日志
	an.ctx.Logger = zap.NewZap(an.ctx.C.Log)
	an.task = task
	for _, prefix := range announcementUrlPrefix {
		for i := 1; i <= announcementPageMax; i++ {
			url := prefix + strconv.Itoa(i) + ".html"
			c := crawler.NewColly(
				an.task,
				url,
			)
			an.crawlers = append(an.crawlers, c)
		}
	}
}

func (an *Announcement) Run() {
	for _, c := range an.crawlers {
		if err := c.Run(crawler.WithCrawlCallback(an.callback(c))); err != nil {
			an.ctx.Logger.Error("爬取文章失败", rawZap.Error(err), rawZap.String("url", c.Url))
		}
	}
	an.toXslx()
	an.ctx.Wg.Done()
	an.ctx.Logger.Info("Run success", rawZap.String("task", an.task))
}

func (an *Announcement) callback(c *crawler.Colly) crawler.Callback {
	return func() {
		c.Crawler.OnHTML(`body > div.container_w > div > div:nth-child(2)`, func(ul *colly.HTMLElement) {
			ul.ForEach(`li > a`, func(i int, a *colly.HTMLElement) {
				uri := a.Attr("href")
				url := util.Host(ul) + uri
				cArticle := crawler.NewColly(
					an.task,
					url,
				)
				if err := cArticle.Run(crawler.WithCrawlCallback(an.getArticle(cArticle))); err != nil {
					an.ctx.Logger.Error("爬取文章失败", rawZap.Error(err), rawZap.String("url", c.Url))
				}
			})

		})

	}
}

func (an *Announcement) getArticle(c *crawler.Colly) crawler.Callback {
	return func() {
		selector := `body > div.container_w > div > div > div:nth-child(2) > div.detail_inner`
		c.Crawler.OnHTML(selector, func(e *colly.HTMLElement) {
			an.id += 1
			article := fip.Article{}
			article.Id = an.id
			selectorTitle := `div.tit_header > h2`
			article.Title = e.DOM.Find(selectorTitle).Text()
			selectorDate := `#zoom > span > p.noice_date`
			article.Date = e.DOM.Find(selectorDate).Contents().Not(`span`).Text()
			selectorSource := `div.tit_header > p > span > span`
			article.Source = e.DOM.Find(selectorSource).Text()
			selectorContent := `#zoom`
			article.Content, _ = e.DOM.Find(selectorContent).Html()
			article.Content = util.FormatArticle(article.Content, e)
			an.articles = append(an.articles, article)
			an.ctx.Logger.Info("爬取成功", rawZap.Int("Id", article.Id), rawZap.String("title", article.Title), rawZap.String("url", c.Url))
		})
	}
}

func (an *Announcement) toXslx() {
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
	for line, article := range an.articles {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(line+2), article.Id)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(line+2), article.Title)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(line+2), article.Content)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(line+2), article.Source)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(line+2), article.Summary)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(line+2), article.Date)
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("announcement.xlsx"); err != nil {
		fmt.Println(err)
	}
}
