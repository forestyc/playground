package main

import (
	"flag"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/config"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/context"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/crawler"
	"log"
	"strings"
)

func main() {
	var conf string
	var task string
	flag.StringVar(&conf, "conf", "./etc/crawler.yaml", "configuration")
	flag.StringVar(&task, "task", "", "gfex-news/gfex-announcement/gfex-focus")
	flag.Parse()
	// 获取配置
	var c config.Config
	if err := config.Load(conf, &c); err != nil {
		log.Fatalln("读取配置文件失败", err)
	}
	// 初始化公共组件
	ctx, err := context.NewGlobalContext(c)
	if err != nil {
		log.Fatalln("初始化公共组件失败", err)
	}
	taskSlice := strings.Split(task, " ")
	// 初始化爬虫
	crawler.Register(ctx)
	// 执行爬虫任务
	taskCount := len(taskSlice)
	if taskCount == 1 && len(taskSlice[0]) == 0 {
		log.Fatalln("请选择爬虫任务")
	}
	ctx.Wg.Add(taskCount)
	crawler.Run(taskSlice)
	ctx.Wg.Wait()
	log.Println("爬虫任务结束")
}
