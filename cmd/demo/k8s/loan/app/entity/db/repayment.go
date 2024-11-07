package db

import "time"

type Repayment struct {
	Id              int64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	LoanBasicInfoId int64     `json:"loan_basic_info_id"`
	Period          int       `json:"period"`
	Amount          float64   `json:"amount"`
	RepaymentDate   time.Time `json:"repayment_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (t *Repayment) TableName() string {
	return "repayment"
}
