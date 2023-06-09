package main

import (
	"fmt"
	"github.com/Baal19905/playground/colly/pkg/crawler"
	"github.com/gocolly/colly/v2"
)

var c crawler.Colly

func main() {
	c = crawler.NewColly("http://www.gfex.com.cn/gfex/bsyw/list_yw.shtml") //, test(&body))
	fmt.Println(c.Run())
}

func test(body *string) crawler.Callback {
	return func() {
		c.Crawler.OnResponse(func(r *colly.Response) {
			*body = string(r.Body)
		})
	}
}
