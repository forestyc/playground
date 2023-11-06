package repayment

import (
	"github.com/forestyc/playground/cmd/demo/house/global"
	"github.com/forestyc/playground/pkg/utils"
	"github.com/pkg/errors"
	"math"
)

type PrincipalInterest struct {
	principal           float64
	interestRate        float64
	periods             int
	interestRateByMonth float64
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
	return pi, nil
}

func (pi PrincipalInterest) Repayment() float64 {
	return pi.principal * pi.interestRateByMonth * math.Pow(1+pi.interestRateByMonth, float64(pi.periods)) /
		(math.Pow(1+pi.interestRateByMonth, float64(pi.periods)) - 1)
}

func (pi PrincipalInterest) TotalInterest() float64 {
	return pi.Repayment()*float64(pi.periods) - pi.principal
}

func (pi PrincipalInterest) Prepayment(period int, amount float64) {
	pi.principal -= amount
	pi.periods -= period
}
