package main

import (
	"github.com/forestyc/playground/pkg/core/http"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// ...
	server := http.NewServer(":8080")
	server.Serve()
	gd := &GinDemo{}
	server.WithHandler(gd)

	time.Sleep(time.Hour)
	// ...
}

type GinDemo struct{}

func (gd *GinDemo) Register(g *gin.Engine) {
	g.GET("/repayment", func(context *gin.Context) {
		context.JSON(200, gin.H{"hello": "world"})
	})
}
