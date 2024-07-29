package loan

import (
	"github.com/forestyc/playground/pkg/utils"
	"github.com/pkg/errors"
	"time"
)

type AverageCapital struct {
	principal           float64
	interestRate        float64
	periods             int
	interestRateByMonth float64
	principalByMonth    float64
	startDate           time.Time
}

func NewAverageCapital(info BasicInfo) (*AverageCapital, error) {
	var ac AverageCapital
	if info.Periods == 0 ||
		utils.EqualFloat(info.Principal, 0, EPSILON) ||
		utils.EqualFloat(info.InterestRate, 0, EPSILON) {
		return &ac, errors.New("invalid param")
	}
	ac.principal = info.Principal
	ac.interestRate = info.InterestRate
	ac.periods = info.Periods
	ac.principalByMonth = ac.principal / float64(ac.periods)
	ac.interestRateByMonth = ac.interestRate / MONTH
	ac.startDate = info.StartDate
	return &ac, nil
}

func (ac *AverageCapital) Repayment(period int) (float64, time.Time) {
	return ac.principalByMonth + (ac.principal-ac.principalByMonth*float64(period-1))*ac.interestRateByMonth,
		ac.startDate.AddDate(0, period, 0)
}

func (ac *AverageCapital) TotalInterest() float64 {
	return float64(ac.periods+1) * ac.principal * ac.interestRateByMonth / 2
}

func (ac *AverageCapital) Prepayment(period int, amount float64) {
	ac.principalByMonth = (ac.principal - amount - ac.principalByMonth*float64(period)) / MONTH
	ac.periods -= period
}
