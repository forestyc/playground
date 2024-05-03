package crawler

import (
	"github.com/forestyc/playground/pkg/core/prometheus"
	"github.com/gocolly/colly/v2"
)

const (
	Success     = "success"
	crawFail    = "crawl fail"
	piplineFail = "pipline fail"
)

type ReqType uint32

const (
	Get = iota
	Post
	PostRaw
)

type Callback func()
type Pipeline func() error
type Option func(c *Colly)

type Colly struct {
	Task     string
	Url      string
	Callback []Callback
	Pip      []Pipeline
	Counter  *prometheus.Counter
	reqType  ReqType
	postForm map[string]string
	postRaw  []byte
	Crawler  *colly.Collector
}

func NewColly(task, url string, options ...Option) *Colly {
	c := Colly{
		Crawler: colly.NewCollector(),
		Url:     url,
		Task:    task,
		Counter: prometheus.NewCounter(
			"crawler_status",
			"记录爬虫执行情况",
			"task", "url", "status"),
	}
	c.Crawler.IgnoreRobotsTxt = false
	for _, option := range options {
		if option != nil {
			option(&c)
		}
	}
	return &c
}

func WithPipeline(pipeline Pipeline) Option {
	return func(c *Colly) {
		if pipeline != nil {
			c.Pip = append(c.Pip, pipeline)
		}
	}
}

func WithCrawlCallback(cb Callback) Option {
	return func(c *Colly) {
		if cb != nil {
			c.Callback = append(c.Callback, cb)
		}
	}
}

func WithReqType(t ReqType) Option {
	return func(c *Colly) {
		c.reqType = t
	}
}

func WithPostForm(form map[string]string) Option {
	return func(c *Colly) {
		c.postForm = form
	}
}

func WithPostRaw(raw []byte) Option {
	return func(c *Colly) {
		c.postRaw = raw
	}
}

func (c *Colly) Run(options ...Option) error {
	// add option
	for _, option := range options {
		if option != nil {
			option(c)
		}
	}
	// crawl
	for _, e := range c.Callback {
		e()
	}
	var err error
	switch c.reqType {
	case Get:
		err = c.Crawler.Visit(c.Url)
	case Post:
		err = c.Crawler.Post(c.Url, c.postForm)
	case PostRaw:
		err = c.Crawler.PostRaw(c.Url, c.postRaw)
	}
	if err != nil {
		c.Counter.Inc(c.Task, c.Url, crawFail)
		return err
	}
	if c.Pip != nil {
		for _, pip := range c.Pip {
			if err = pip(); err != nil {
				c.Counter.Inc(c.Task, c.Url, piplineFail)
				return err
			}
		}
	}
	c.Counter.Inc(c.Task, c.Url, Success)
	return nil
}
