# 一、项目结构

```
├─algorithm
├─cmd
│  └─crawler
├─demo
│  ├─channel
│  ├─chromedp
│  ├─context
│  ├─docker
│  ├─etcd
│  ├─gin
│  ├─go-zero
│  ├─gomod
│  ├─micro
│  ├─reflect
│  ├─selenium
│  ├─tag
│  ├─ticker
│  ├─timer
│  └─url
└─pkg
    ├─crawler
    ├─crypto
    ├─db
    ├─distributed
    ├─http
    ├─log
    ├─prometheus
    ├─redis
    ├─security
    ├─util
    └─version
```

# 1、应用

在`cmd`目录自行创建应用目录，每个应用应有自己的`README.md`和`CHANGELOG.md`。

# 2、中间件

## 1、目录结构

| 文件夹         |    说明    | 描述                     |
| :------------- | :--------: | :----------------------- |
| `pkg`          |            |                          |
| `─crawler`     |    缓存    | go-redis                 |
| `─crypto`      |    爬虫    | colly                    |
| `─db`          |   数据库   | gorm                     |
| `─distributed` |    log     | zap                      |
| `─http`        | prometheus | prometheus metrics       |
| `─log`         |  version   | 记录版本信息，由CICD传入 |
| `─prometheus`  |  version   | 记录版本信息，由CICD传入 |
| `─redis`       |  version   | 记录版本信息，由CICD传入 |
| `─security`    |  version   | 记录版本信息，由CICD传入 |
| `─util`        |  version   | 记录版本信息，由CICD传入 |
| `─version`     |  version   | 记录版本信息，由CICD传入 |



