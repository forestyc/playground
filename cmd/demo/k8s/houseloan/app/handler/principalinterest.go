package handler

import (
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/model"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/service"
	"github.com/forestyc/playground/pkg/core/message"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PrincipalInterest struct {
	periodsService                        *service.Periods
	businessPrincipalInterestService      *service.PrincipalInterest
	providentFundPrincipalInterestService *service.PrincipalInterest
	ctx                                   context.Context
}

func NewPrincipalInterest(
	startDate string, periods int,
	businessPrincipal, businessInterestRate,
	providentFundPrincipal, providentFundInterestRate float64,
) *PrincipalInterest {
	return &PrincipalInterest{
		periodsService:                        service.NewPeriods(startDate, periods),
		businessPrincipalInterestService:      service.NewPrincipalInterest(businessPrincipal, businessInterestRate, periods),
		providentFundPrincipalInterestService: service.NewPrincipalInterest(providentFundPrincipal, providentFundInterestRate, periods),
	}
}

func (pi *PrincipalInterest) Register(engine *gin.Engine) {
	engine.GET("/repayment", pi.repayment())
	engine.GET("/probe-liveness", pi.probeLiveness())
}

func (pi *PrincipalInterest) repayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		period := pi.periodsService.GetPeriod(
			c.DefaultQuery("date", time.Now().Format("2006-01-02")),
		)
		if period == 0 {
			c.JSON(http.StatusOK, message.Failed())
			return
		}
		response := model.RepaymentResp{
			Business:      pi.businessPrincipalInterestService.Repayment(period),
			ProvidentFund: pi.providentFundPrincipalInterestService.Repayment(period),
		}
		response.Total = response.ProvidentFund + response.Business
		c.JSON(http.StatusOK, message.SuccessWithObject(response))
	}
}

func (pi *PrincipalInterest) probeLiveness() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Custom-Header", "Awesome")
		context.JSON(http.StatusOK, "health")
	}
}
