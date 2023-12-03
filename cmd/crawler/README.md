# crawler

## 爬虫列表

| 交易所    | 内容   | 标签     |
|--------|------|--------|
| `yage` | `雅歌` | `yage` |

## 执行爬虫

### 1、执行单个爬虫

    ./crawler -task "gfex-news"

### 2、执行多个爬虫

    ./crawler -task "gfex-news gfex-announcement gfex-focus"

**注意：**

+ **-task后的参数必须在""内**

## 添加爬虫任务

### 1、目录结构

| 文件夹             | 说明        | 描述        |
|:----------------|:----------|:----------|
| `crawler`       |           |           |
| `-app`          | 代码目录      | 代码目录      |
| `---common`     | common    | common    |
| `---config`     | config    | config    |
| `---context`    | context   | context   |
| `---dao`        | dao       | dao       |
| `---handler`    | dao       | dao       |
| `---crawler`    | crawler   | crawler   |
| `---util`       | util      | util      |
| `-etc`          | 配置目录      | 配置目录      |
| `-crawler.go`   | main      | main      |
| `-README.md`    | readme    | readme    |
| `-CHANGELOG.md` | changelog | changelog |

### 2、实现接口

```
type Crawler interface {
    Init(ctx context.Context, config zap.Config)
    Run()
}
```

### 3、实现爬取逻辑

+ 在logic创建someexchange目录，并实现logic

```
type Something struct {
    ctx     context.Context   // 必须包含context，ctx包含db及cache等公共中间件
    // todo: 自定义
    ...
}

func (cs *Something) Init(ctx context.Context, config zap.Config) {
    // 初始化ctx、日志（必须包含）
    cs.ctx = ctx
    cs.ctx.C.Log = config
    cs.ctx.Logger = zap.NewZap(cs.ctx.C.Log)
    // todo: 自定义
    ...
}

func (cs *Something) Run() {
	// todo: 自定义
	...
	cs.ctx.Wg.Done()    // 必须调用，通知主协程任务完成
}
```

+ 在handler创建someexchange目录，并实现handler

### 3、注册任务

打开`cmd/crawler/logic/schedule.go`，注册任务

```
// Register 注册
func Register(c context.Context) {
    ctx = c
    crawler = map[string]Crawler{
        "gfex-news":              &gfex.GfexNews{},                 // 广期所-本所要闻
        "gfex-announcement":      &gfex.GfexAnnouncement{},         // 广期所-通知公告
        "gfex-focus":             &gfex.GfexFocus{},                // 广期所-媒体聚焦
        "someexchange-something": &someexchange.Something{},        // 某交易所-某内容
    }
}
```



