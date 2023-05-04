package main

import (
	"fmt"
	"github.com/Baal19905/playground/colly/pkg/crawler"
	"github.com/gocolly/colly/v2"
)

var c crawler.Colly

func main() {
	var body string
	c = crawler.NewColly("http://www.gfex.com.cn/u/interfacesWebTtQueryTradPara/loadDayList", test(&body))
	c.Run()
	fmt.Println(body)
}

func test(body *string) crawler.Callback {
	return func() {
		c.Crawler.OnResponse(func(r *colly.Response) {
			*body = string(r.Body)
		})
	}
}
