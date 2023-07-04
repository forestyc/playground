package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/", Greetings)
	engine.Run()
}

func Greetings(c *gin.Context) {
	c.JSON(http.StatusOK, "greetings")
}
