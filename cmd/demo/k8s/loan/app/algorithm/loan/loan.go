package loan

import (
	"github.com/pkg/errors"
	"time"
)

type Loan interface {
	Repayment(period int) (float64, time.Time)
}

type BasicInfo struct {
	Principal    float64
	InterestRate float64
	Periods      int
	StartDate    time.Time
}

const (
	TypePrincipalInterest = iota + 1
	TypeAverageCapital
)

func NewLoan(loanType int, info BasicInfo) (Loan, error) {
	switch loanType {
	case TypeAverageCapital:
		return NewAverageCapital(info)
	case TypePrincipalInterest:
		return NewPrincipalInterest(info)
	default:
		return nil, errors.New("invalid loan type")
	}
}
