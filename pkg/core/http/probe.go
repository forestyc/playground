package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func probeLiveness() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Custom-Header", "Awesome")
		context.JSON(http.StatusOK, "health")
	}
}
