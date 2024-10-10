package db

import "time"

type LoanBasicInfo struct {
	Id           int64 `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string
	LoanId       int64
	LoanType     int
	Principal    float64
	InterestRate float64
	Periods      int
	StartDate    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (t *LoanBasicInfo) TableName() string {
	return "loan_basic_info"
}
