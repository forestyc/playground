# http

http server and client

## 1. Client
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

## 2. Server
```
import (
	"github.com/forestyc/playground/pkg/http"
	"github.com/gin-gonic/gin"
	netHttp "net/http"
)

func main() {
	// ...
	server := http.NewServer(":8080")
	server.Serve()
	server.Router.GET("/", func(c *gin.Context) {
		c.String(netHttp.StatusOK, "Hello World")
	})
	// ...
}
```



