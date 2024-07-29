package loan

import (
	"github.com/forestyc/playground/pkg/utils"
	"github.com/pkg/errors"
	"time"
)

type PrincipalInterest struct {
	principal           float64
	interestRate        float64
	periods             int
	interestRateByMonth float64
	principalByMonth    float64
	startDate           time.Time
}

func NewPrincipalInterest(info BasicInfo) (*PrincipalInterest, error) {
	var pi PrincipalInterest
	if info.Periods == 0 ||
		utils.EqualFloat(info.Principal, 0, EPSILON) ||
		utils.EqualFloat(info.InterestRate, 0, EPSILON) {
		return &pi, errors.New("invalid param")
	}
	pi.principal = info.Principal
	pi.interestRate = info.InterestRate
	pi.periods = info.Periods
	pi.interestRateByMonth = pi.interestRate / MONTH
	pi.principalByMonth = utils.Round(pi.principal/float64(info.Periods), 2)
	pi.startDate = info.StartDate
	return &pi, nil
}

func (pi *PrincipalInterest) Repayment(period int) (float64, time.Time) {
	interest := utils.Round((pi.principal-float64(period-1)*pi.principalByMonth)*pi.interestRateByMonth, 2)
	return utils.Round(pi.principalByMonth+interest, 2), pi.startDate.AddDate(0, period, 0)
}
