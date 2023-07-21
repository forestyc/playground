# prometheus

## 1. 配置
```
// 添加配置
prometheus:
  addr: :12112      # 监听地址
  path: /metrics    # metrics路径
```

## 2. 代码
### 1. 读取配置
```
import "github.com/Baal19905/playground/pkg/prometheus"

// 在配置中添加Prometheus对象
type Config struct {
    // ...
    Prometheus prometheus.Prometheus `mapstructure:"prometheus"`
    // ...
}

// load config
```
### 2. 启动Prometheus agent
```
import "github.com/Baal19905/playground/pkg/prometheus"

// 加载配置后
func Run() {
    // ...
    if prom.HasPrometheus() {
		ctx.C.Prometheus.Start()
    }
    // ...
}
```
### 3. 记录指标
```
import "github.com/Baal19905/playground/pkg/prometheus"

// 初始化指标
func (j *job) Init() {
    // ...
    j.c = prometheus.NewCounter("name","help", "label1","label2")
    // ...
}

// 记录指标
func (j *job) Run() {
    // ...
    j.c.Inc("value1", "value2")
    // ...
}
```



