package db

import "time"

type Info struct {
	Id                        *int
	Name                      *string
	BusinessPrincipal         *float64
	BusinessInterestRate      *float64
	BusinessPeriods           *int
	ProvidentFundPrincipal    *float64
	ProvidentFundInterestRate *float64
	ProvidentPeriods          *int
	StartDate                 *time.Time
	CreatedAt                 *time.Time
	UpdatedAt                 *time.Time
}

func (t *Info) TableName() string {
	return "info"
}
