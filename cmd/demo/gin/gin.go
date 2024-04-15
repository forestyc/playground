package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	go Run(r)
	time.Sleep(time.Second * 3)
	r.GET("/hello", func(c *gin.Context) { c.JSON(http.StatusOK, "world") })
	time.Sleep(time.Hour)
}

func Run(e *gin.Engine) {
	e.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
