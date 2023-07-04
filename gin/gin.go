package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/.well-known/pki-validation/3DC8B7244EF16DBD8645BD97DEBB3555.txt", "3DC8B7244EF16DBD8645BD97DEBB3555.txt")
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
