package main

import (
	"fmt"
	"github.com/forestyc/playground/pkg/crawler"
)

func main() {
	colly := crawler.NewColly(
		"demo",
		"https://fipinfo.dfitc.com.cn/frontend/fia/fip_report.html",
		crawler.WithCrawlCallback(callback),
		crawler.WithCrawlCallback(callback),
		crawler.WithCrawlCallback(callback),
	)
	colly.Run()
}

func callback() {
	fmt.Println("hehe")
}
