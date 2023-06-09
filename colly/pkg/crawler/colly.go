package crawler

import (
	"errors"
	"github.com/Baal19905/playground/colly/pkg/crawler/robots"
	"github.com/gocolly/colly/v2"
)

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
	robot := robots.NewRobots(c.Url, c.Crawler.UserAgent)
	robot.Run()
	if robot.Disallow(c.Url) {
		return errors.New("DISALLOW!!!")
	}
	for _, e := range c.Callback {
		e()
	}
	return c.Crawler.Visit(c.Url)
}
