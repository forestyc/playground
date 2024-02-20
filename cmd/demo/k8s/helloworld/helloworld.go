package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello-world", helloWorld())
	engine.GET("/health", health())
	engine.Run()
}

func helloWorld() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello world!")
	}
}

func health() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Custom-Header", "Awesome")
		context.JSON(http.StatusOK, "health")
	}
}
