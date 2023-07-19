package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	filePtr, err := os.OpenFile("mylogs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = filePtr
	e := gin.Default()
	e.GET("/", Root)
	e.Run()
}

func Root(c *gin.Context) {
	c.JSON(http.StatusOK, "root")
}
