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

| 文件夹    | 说明 | 描述 |
| :-------- | :--- | :--- |
| `crawler` | 爬虫 | 爬虫 |


# 2、包

## 1、目录结构

| 文件夹        | 说明        | 描述                     |
| :------------ | :---------- | :----------------------- |
| `crawler`     | crawler     | colly封装                |
| `crypto`      | crypto      | 对称、非对称、摘要       |
| `db`          | db          | gorm                     |
| `distributed` | distributed | 注册中心、分布式锁       |
| `http`        | http        | http封装                 |
| `log`         | log         | zap封装                  |
| `prometheus`  | prometheus  | exporter                 |
| `redis`       | redis       | go-redis封装             |
| `security`    | security    | security                 |
| `util`        | util        | util                     |
| `version`     | version     | 记录版本信息，由CICD传入 |
| `jwt`         | jwt         | jwt token                |
| `mail`        | mail        | mail                     |



