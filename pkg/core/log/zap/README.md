# zap

## 1. 配置
```
// 添加配置
log:
  level: info
  format: console
  director: ./log
  show-line: true
  encode-level: LowercaseColorLevelEncoder
  stacktrace-key: stacktrace
  log-in-console: true
```

## 2. 代码
### 1. 读取配置
```
import "github.com/Baal19905/playground/pkg/log/zap"

// 在配置中添加zap对象
type Config struct {
    // ...
    Log        zap.Config            `mapstructure:"log"`
    // ...
}

// load config
```

### 2. 使用db
```
import "github.com/Baal19905/playground/pkg/log/zap"

// 初始化指标
func (j *job) Init() {
    // ...
    j.logger = zap.NewZap(config)
    // ...
}

// 记录指标
func (j *job) Run() {
    // ...
    j.logger.Info("log")
    // ...
}
```



