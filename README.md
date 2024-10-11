# Introduction

A library for golang application servers including mysql, redis, crypto etc. For detail see `pkg` directory.

| DIRECTORY         | DESCRIPTION             |
|:------------------|:------------------------|
| `pkg/concurrency` | workpool                |
| `pkg/crawler`     | colly crawler           |
| `pkg/core`        | redis/db/jwt/log...     |
| `pkg/distributed` | register/lock/snowflake |
| `pkg/security`    | security                |
| `pkg/utils`       | utils                   |
| `pkg/mail`        | mail sender             |
| `pkg/encoding`    | base64/pem              |

# Installation

```shell
go get github.com/forestyc/playground
```

# Contributions

Issue and PR are welcomed

# License

[MIT License](https://github.com/forestyc/playground/blob/master/LICENSE)