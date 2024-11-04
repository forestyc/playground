package handler

import (
	"net/http"
	"strconv"

	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/model"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/service"
	"github.com/forestyc/playground/pkg/core/message"
	"github.com/gin-gonic/gin"
)

type LoanBasicInfo struct {
	loanBasicInfoService *service.LoanBasicInfo
}

func NewLoanBasicInfo(ctx *context.Context) *LoanBasicInfo {
	return &LoanBasicInfo{
		loanBasicInfoService: service.NewLoanBasicInfo(ctx),
	}
}

func (lbi *LoanBasicInfo) Register(engine *gin.Engine) {
	uriBasicInfo := "/basic-info"
	engine.GET(uriBasicInfo, lbi.getBasicInfo())
	engine.GET(uriBasicInfo+"/list", lbi.getBasicInfoList())
	engine.POST(uriBasicInfo, lbi.createBasicInfo())
	engine.PUT(uriBasicInfo, lbi.modifyBasicInfo())
	engine.DELETE(uriBasicInfo, lbi.deleteBasicInfo())
	engine.PUT(uriBasicInfo+"/modify-interest-rate", lbi.cutInterestRate())
}

func (lbi *LoanBasicInfo) getBasicInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.DefaultQuery("id", "0")
		id, _ := strconv.ParseInt(query, 10, 64)
		if basicInfo, err := lbi.loanBasicInfoService.GetById(id); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
		} else {
			c.JSON(http.StatusOK, message.SuccessWithObject(basicInfo))
		}
	}
}

func (lbi *LoanBasicInfo) getBasicInfoList() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.DefaultQuery("loan_id", "0")
		id, _ := strconv.ParseInt(query, 10, 64)
		if basicInfo, err := lbi.loanBasicInfoService.GetByLoanId(id); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
		} else {
			c.JSON(http.StatusOK, message.SuccessWithObject(basicInfo))
		}
	}
}

func (lbi *LoanBasicInfo) deleteBasicInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.DefaultQuery("id", "0")
		id, _ := strconv.ParseInt(query, 10, 64)
		if err := lbi.loanBasicInfoService.Delete(id); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
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
		if err := lbi.loanBasicInfoService.Create(req); err != nil {
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
		if err := lbi.loanBasicInfoService.Modify(req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		c.JSON(http.StatusOK, message.Success())
	}
}

func (lbi *LoanBasicInfo) cutInterestRate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CutInterestRateReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		if err := lbi.loanBasicInfoService.CutInterestRate(req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		c.JSON(http.StatusOK, message.Success())
	}
}
