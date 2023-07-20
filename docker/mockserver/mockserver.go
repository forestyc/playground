package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	filePtr, err := os.OpenFile("log/mylogs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(filePtr)
	gin.DefaultWriter = filePtr
	e := gin.Default()
	e.GET("/", Root)
	e.POST("/create", Create)
	e.Run()
}

func Root(c *gin.Context) {
	c.JSON(http.StatusOK, "root")
}
func Create(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	log.Println("request body .", string(data))
	c.JSON(http.StatusOK, "ok")
}
