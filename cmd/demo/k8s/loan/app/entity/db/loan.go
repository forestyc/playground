package db

import "time"

type Loan struct {
	Id        int64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Loan) TableName() string {
	return "loan"
}
