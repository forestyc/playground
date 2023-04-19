package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping/:id/:type", func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		fmt.Println(c.Param("type"))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
