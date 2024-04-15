package repayment

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
	pi.principalByMonth = pi.principal / float64(periods)
	return pi
}

func (pi PrincipalInterest) Repayment(period int) float64 {
	interest := (pi.principal - float64(period-1)*pi.principalByMonth) * pi.interestRateByMonth
	return pi.principalByMonth + interest
}
