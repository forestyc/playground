package handler

import "github.com/gin-gonic/gin"

func NewPrincipalInterestRouters(prefix string, engine *gin.Engine) {
	group := engine.Group(prefix)
	group.GET("/")
}
