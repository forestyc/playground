package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello-world", helloWorld())
	engine.Run()
}

func helloWorld() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello world!")
	}
}
