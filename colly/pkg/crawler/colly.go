package crawler

import "github.com/gocolly/colly/v2"

type Callback func()

type Colly struct {
	Url      string
	Callback []Callback
	Crawler  *colly.Collector
}

func NewColly(url string, cb ...Callback) Colly {
	c := Colly{
		Crawler:  colly.NewCollector(),
		Url:      url,
		Callback: cb,
	}
	return c
}

func (c Colly) Run() error {
	for _, e := range c.Callback {
		e()
	}
	return c.Crawler.Visit(c.Url)
}
