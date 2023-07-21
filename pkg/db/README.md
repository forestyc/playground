# db

## 1. 配置
```
// 添加配置
database:
  dsn: user:password@tcp(localhost:3306)/db?charset=utf8
  max-open: 10
  idle-conns: 5
  idle-timeout: 300 # 5分钟
  operation-timeout: 10 # 10秒
```

## 2. 代码
### 1. 读取配置
```
import "github.com/Baal19905/playground/pkg/db"

// 在配置中添加redis对象
type Config struct {
    // ...
    Database db.Config  `mapstructure:"database"`
    // ...
}

// load config
```

### 2. 使用db
```
import "github.com/Baal19905/playground/pkg/db"

// 初始化指标
func (j *job) Init() {
    // ...
    j.db = db.NewMysql(config)
    // ...
}

// 记录指标
func (j *job) Run() {
    // ...
    j.db.Create(rows)
    // ...
}
```



