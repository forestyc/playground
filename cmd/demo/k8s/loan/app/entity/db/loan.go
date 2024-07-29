package db

import "time"

type Loan struct {
	Id        int64 `gorm:"primary_key;AUTO_INCREMENT"`
	LoanId    int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Loan) TableName() string {
	return "loan"
}
