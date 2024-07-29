package db

type LoanType struct {
	Id    int64 `gorm:"primary_key;AUTO_INCREMENT"`
	Name  string
	Value string
}

func (t *LoanType) TableName() string {
	return "loan_type"
}
