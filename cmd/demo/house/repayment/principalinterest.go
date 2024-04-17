package repayment

import (
	"github.com/forestyc/playground/cmd/demo/house/global"
	"github.com/forestyc/playground/pkg/utils"
	"github.com/pkg/errors"
)

type PrincipalInterest struct {
	principal           float64
	interestRate        float64
	periods             int
	interestRateByMonth float64
	principalByMonth    float64
}

func NewPrincipalInterest(principal, interestRate float64, periods int) (PrincipalInterest, error) {
	var pi PrincipalInterest
	if periods == 0 ||
		utils.EqualFloat(principal, 0, global.EPSILON) ||
		utils.EqualFloat(interestRate, 0, global.EPSILON) {
		return pi, errors.New("invalid param")
	}
	pi.principal = principal
	pi.interestRate = interestRate
	pi.periods = periods
	pi.interestRateByMonth = pi.interestRate / global.MONTH
	pi.principalByMonth = utils.Round(pi.principal/float64(periods), 2)
	return pi, nil
}

func (pi PrincipalInterest) Repayment(period int) float64 {
	interest := utils.Round((pi.principal-float64(period-1)*pi.principalByMonth)*pi.interestRateByMonth, 2)
	return utils.Round(pi.principalByMonth+interest, 2)
}
