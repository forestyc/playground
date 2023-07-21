# http

http客户端封装，可选择长连接或短连接客户端

## 1. 使用
```
import (
    "fmt"
    "github.com/Baal19905/playground/pkg/http"
)

func main() {
    client := http.NewClient(false)
    defer client.Close()
    resp, err := client.Do(
        "get",
        "http://localhost:8080",
        nil,
        nil,
    )
    
    if err != nil {
        panic(err)
    }
    fmt.Println(string(resp))
}
```



