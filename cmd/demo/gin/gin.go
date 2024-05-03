package main

import (
	"github.com/forestyc/playground/pkg/core/http"
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
