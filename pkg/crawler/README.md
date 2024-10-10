# crawler

## example

```
// init
crawler = crawler.NewColly(
           task,
           url,
           // db pipeline回调函数(可选)
           crawler.WithPipeline(pipeline),
           // 爬虫回调函数(可选)
           crawler.WithCrawlCallback(callback),
           // ...
        )

// run
crawler.Run(
    // db pipeline回调函数(可选)
    crawler.WithPipeline(pipeline),
    // 爬虫回调函数(可选)
    crawler.WithCrawlCallback(callback),
)
```



