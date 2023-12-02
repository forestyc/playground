# crawler

## example

```
// 初始化爬虫
crawler = crawler.NewColly(
           task,
           url,
           // db pipeline回调函数(可选)
           crawler.WithPipeline(pipeline),
           // 爬虫回调函数(可选)
           crawler.WithCrawlCallback(callback),
           // ...
        )

// 执行爬虫
crawler.Run(
    // db pipeline回调函数(可选)
    crawler.WithPipeline(pipeline),
    // 爬虫回调函数(可选)
    crawler.WithCrawlCallback(callback),
)
```



