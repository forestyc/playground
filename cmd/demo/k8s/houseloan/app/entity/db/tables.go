package db

import "time"

type ConfigInfo struct {
	Id                        *int `gorm:"primary_key;AUTO_INCREMENT"`
	Name                      *string
	BusinessPrincipal         *float64
	BusinessInterestRate      *float64
	BusinessPeriods           *int
	ProvidentFundPrincipal    *float64
	ProvidentFundInterestRate *float64
	ProvidentFundPeriods      *int
	StartDate                 *time.Time
	CreatedAt                 *time.Time
	UpdatedAt                 *time.Time
}

func (t *ConfigInfo) TableName() string {
	return "config_info"
}
