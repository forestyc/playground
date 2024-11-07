package db

import "time"

type LoanBasicInfo struct {
	Id           int64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name         string    `json:"name"`
	LoanId       int64     `json:"loan_id"`
	LoanType     int       `json:"loan_type"`
	Principal    float64   `json:"principal"`
	InterestRate float64   `json:"interest_rate"`
	Periods      int       `json:"periods"`
	StartDate    time.Time `json:"start_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (t *LoanBasicInfo) TableName() string {
	return "loan_basic_info"
}
