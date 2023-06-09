package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	e := gin.Default()
	e.GET("/", Root)
	e.Run()
}

func Root(c *gin.Context) {
	log.Println("root")
	c.JSON(http.StatusOK, "root")
}
