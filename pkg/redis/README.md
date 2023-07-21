# redis

## 1. 配置
```
// 添加配置
redis:
  address: localhost:6379
  password: 1234
  max-open: 10
  idle-timeout: 300     # 5分钟
  idle-conns: 5
  operation-timeout: 10 # 10秒
```

## 2. 代码
### 1. 读取配置
```
import "github.com/Baal19905/playground/pkg/redis"

// 在配置中添加redis对象
type Config struct {
    // ...
    Redis      redis.Config          `mapstructure:"redis"`
    // ...
}

// load config
```

### 2. 使用redis
```
import "github.com/Baal19905/playground/pkg/redis"

// 初始化指标
func (j *job) Init() {
    // ...
    j.r, err := redis.NewRedis(config)
    if err != nil {
        return ctx, err
    }
    // ...
}

// 记录指标
func (j *job) Run() {
    // ...
    j.r.Get(ctx, "key")
    // ...
}
```



