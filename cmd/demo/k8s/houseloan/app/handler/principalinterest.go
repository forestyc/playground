package handler

import (
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/model/vo"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
}

func (pi *PrincipalInterest) repayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		date := c.Query("date")
		period := pi.periodsService.GetPeriod(date)
		if period == 0 {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		response := vo.Repayment{
			Business:      pi.businessPrincipalInterestService.Repayment(period),
			ProvidentFund: pi.providentFundPrincipalInterestService.Repayment(period),
		}
		c.JSON(http.StatusOK, response)
	}
}
