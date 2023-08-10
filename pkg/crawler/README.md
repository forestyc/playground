# crawler

## example

```
// 初始化爬虫
crawler = crawler.NewColly(
           task,
           url,
           // db pipeline回调函数
           crawler.WithPipeline(pipeline),
           // 爬虫回调函数
           crawler.WithCrawlCallback(callback),
           // ...
        )

// 执行爬虫
crawler.Run()
```



