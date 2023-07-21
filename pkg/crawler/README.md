# crawler

## example

```
// 初始化爬虫
crawler = crawler.NewColly(
           task,
           url,
           // db pipeline回调函数
           func pipeline() error {
               // do something
           	return nil
           },
           // 爬虫回调函数
           func callback(){
               return func() {
                   // colly api
                     }
           },
           // ...
        )

// 执行爬虫
crawler.Run()
```



