package crawler

import (
	"github.com/forestyc/playground/pkg/crawler/robots"
	"github.com/forestyc/playground/pkg/prometheus"
	"github.com/gocolly/colly/v2"
	"github.com/pkg/errors"
)

const (
	Success     = "success"
	crawFail    = "crawl fail"
	piplineFail = "pipline fail"
)

type Callback func()
type Pipeline func() error
type Option func(c *Colly)

type Colly struct {
	Task     string
	Url      string
	Callback []Callback
	Pip      Pipeline
	Counter  *prometheus.Counter
	Crawler  *colly.Collector
}

func NewColly(task, url string, options ...Option) Colly {
	c := Colly{
		Crawler: colly.NewCollector(),
		Url:     url,
		Task:    task,
		Counter: prometheus.NewCounter(
			"crawler_status",
			"记录爬虫执行情况",
			"task", "url", "status"),
	}
	for _, option := range options {
		option(&c)
	}
	return c
}

func WithPipeline(pipeline Pipeline) Option {
	return func(c *Colly) {
		c.Pip = pipeline
	}
}

func WithCrawlCallback(cb Callback) Option {
	return func(c *Colly) {
		c.Callback = append(c.Callback, cb)
	}
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
