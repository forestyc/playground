package db

import "time"

type Repayment struct {
	Id            int64 `gorm:"primary_key;AUTO_INCREMENT"`
	LoanId        int64
	Period        int
	Amount        float64
	RepaymentDate time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (t *Repayment) TableName() string {
	return "repayment"
}
