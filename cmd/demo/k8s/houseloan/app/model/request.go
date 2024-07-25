package model

import "time"

type SetInfoReq struct {
	Id                        int       `gin:"id" binding:"required"`
	Name                      string    `gin:"name" binding:"required"`
	BusinessPrincipal         float64   `gin:"business_principal" binding:"required"`
	BusinessInterestRate      float64   `gin:"business_interestRate" binding:"required"`
	BusinessPeriods           int       `gin:"business_periods" binding:"required"`
	ProvidentFundPrincipal    float64   `gin:"provident_fundPrincipal" binding:"required"`
	ProvidentFundInterestRate float64   `gin:"provident_fundInterestRate" binding:"required"`
	ProvidentFundPeriods      int       `gin:"provident_periods" binding:"required"`
	StartDate                 time.Time `gin:"start_date" binding:"required" validate:"format=2006-01-02 15:04:05"`
}
