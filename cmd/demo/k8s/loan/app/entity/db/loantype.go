package db

type LoanType struct {
	Id    int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (t *LoanType) TableName() string {
	return "loan_type"
}
