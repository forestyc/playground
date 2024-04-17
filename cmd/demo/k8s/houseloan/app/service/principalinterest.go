package service

import (
	"github.com/forestyc/playground/pkg/utils"
)

type PrincipalInterest struct {
	principal           float64
	interestRate        float64
	periods             int
	interestRateByMonth float64
	principalByMonth    float64
}

func NewPrincipalInterest(principal, interestRate float64, periods int) PrincipalInterest {
	var pi PrincipalInterest
	pi.principal = principal
	pi.interestRate = interestRate
	pi.periods = periods
	pi.interestRateByMonth = pi.interestRate / MONTH
	pi.principalByMonth = utils.Round(pi.principal/float64(periods), 2)
	return pi
}

func (pi PrincipalInterest) Repayment(period int) float64 {
	interest := utils.Round((pi.principal-float64(period-1)*pi.principalByMonth)*pi.interestRateByMonth, 2)
	return utils.Round(pi.principalByMonth+interest, 2)
}
