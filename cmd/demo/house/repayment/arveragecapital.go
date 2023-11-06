package repayment

import (
	"github.com/forestyc/playground/cmd/demo/house/global"
	"github.com/forestyc/playground/pkg/utils"
	"github.com/pkg/errors"
)

type AverageCapital struct {
	principal           float64
	interestRate        float64
	periods             int
	principalByMonth    float64
	interestRateByMonth float64
}

func NewAverageCapital(principal, interestRate float64, periods int) (AverageCapital, error) {
	var ac AverageCapital
	if periods == 0 ||
		utils.EqualFloat(principal, 0, global.EPSILON) ||
		utils.EqualFloat(interestRate, 0, global.EPSILON) {
		return ac, errors.New("invalid param")
	}
	ac.principal = principal
	ac.interestRate = interestRate
	ac.periods = periods
	ac.principalByMonth = ac.principal / float64(ac.periods)
	ac.interestRateByMonth = ac.interestRate / global.MONTH
	return ac, nil
}

func (ac AverageCapital) Repayment(period int) float64 {
	return ac.principalByMonth + (ac.principal-ac.principalByMonth*float64(period-1))*ac.interestRateByMonth
}

func (ac AverageCapital) TotalInterest() float64 {
	return float64(ac.periods+1) * ac.principal * ac.interestRateByMonth / 2
}

func (ac AverageCapital) Prepayment(period int, amount float64) {
	ac.principalByMonth = (ac.principal - amount - ac.principalByMonth*float64(period)) / global.MONTH
	ac.periods -= period
}
