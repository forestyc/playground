package crawler

import (
	"errors"
	"github.com/forestyc/playground/pkg/crawler/robots"
	"github.com/forestyc/playground/pkg/prometheus"
	"github.com/gocolly/colly/v2"
)

const (
	Success     = "success"
	crawFail    = "crawl fail"
	piplineFail = "pipline fail"
)

type Callback func()
type Pipeline func() error

type Colly struct {
	Task     string
	Url      string
	Callback []Callback
	Pip      Pipeline
	Counter  *prometheus.Counter
	Crawler  *colly.Collector
}

func NewColly(task, url string, pip Pipeline, cb ...Callback) Colly {
	c := Colly{
		Crawler:  colly.NewCollector(),
		Url:      url,
		Task:     task,
		Callback: cb,
		Pip:      pip,
		Counter: prometheus.NewCounter(
			"crawler_status",
			"记录爬虫执行情况",
			"task", "url", "status"),
	}
	return c
}

func (c Colly) Run() error {
	// check disallow
	robot := robots.NewRobots(c.Url, c.Crawler.UserAgent)
	robot.Run()
	if robot.Disallow(c.Url) {
		return errors.New("DISALLOW!!!")
	}
	// crawl
	for _, e := range c.Callback {
		e()
	}
	var err error
	if err = c.Crawler.Visit(c.Url); err != nil {
		c.Counter.Inc(c.Task, c.Url, crawFail)
		return err
	}
	if c.Pip != nil {
		if err = c.Pip(); err != nil {
			c.Counter.Inc(c.Task, c.Url, piplineFail)
			return err
		}
	}
	c.Counter.Inc(c.Task, c.Url, Success)
	return nil
}