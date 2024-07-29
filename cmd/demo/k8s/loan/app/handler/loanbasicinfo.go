package handler

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/model"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/service"
	"github.com/forestyc/playground/pkg/core/message"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoanBasicInfo struct {
	loanService *service.Loan
}

func NewLoanBasicInfo(ctx *context.Context) *LoanBasicInfo {
	return &LoanBasicInfo{
		loanService: service.NewLoan(ctx),
	}
}

func (lbi *LoanBasicInfo) Register(engine *gin.Engine) {
	engine.POST("/basic-info", lbi.createBasicInfo())
}

func (lbi *LoanBasicInfo) createBasicInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateInfoReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		if err := lbi.loanService.Create(req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		c.JSON(http.StatusOK, message.Success())
	}
}
