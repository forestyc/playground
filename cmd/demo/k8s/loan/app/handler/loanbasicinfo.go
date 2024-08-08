package handler

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/model"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/service"
	"github.com/forestyc/playground/pkg/core/message"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LoanBasicInfo struct {
	loanService *service.LoanBasicInfo
}

func NewLoanBasicInfo(ctx *context.Context) *LoanBasicInfo {
	return &LoanBasicInfo{
		loanService: service.NewLoan(ctx),
	}
}

func (lbi *LoanBasicInfo) Register(engine *gin.Engine) {
	uriBasicInfo := "/basic-info"
	engine.GET(uriBasicInfo, lbi.getBasicInfo())
	engine.POST(uriBasicInfo, lbi.createBasicInfo())
	engine.PUT(uriBasicInfo, lbi.modifyBasicInfo())
	engine.DELETE(uriBasicInfo, lbi.deleteBasicInfo())
}

func (lbi *LoanBasicInfo) getBasicInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.DefaultQuery("id", "0")
		id, _ := strconv.ParseInt(query, 10, 64)
		if basicInfo, err := lbi.loanService.Take(id); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, message.SuccessWithObject(basicInfo))
		}
	}
}

func (lbi *LoanBasicInfo) deleteBasicInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.DefaultQuery("id", "0")
		id, _ := strconv.ParseInt(query, 10, 64)
		if err := lbi.loanService.Delete(id); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, message.Success())
		}
	}
}

func (lbi *LoanBasicInfo) createBasicInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateBasicInfoReq
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

func (lbi *LoanBasicInfo) modifyBasicInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ModifyBasicInfoReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		if err := lbi.loanService.Modify(req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		c.JSON(http.StatusOK, message.Success())
	}
}
