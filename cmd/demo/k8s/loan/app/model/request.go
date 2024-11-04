package model

import "time"

type CreateBasicInfoReq struct {
	LoanId       int64     `form:"loan_id" binding:"required"`
	Name         string    `form:"name" binding:"required"`
	Principal    float64   `form:"principal" binding:"required"`
	InterestRate float64   `form:"interest_rate" binding:"required"`
	Periods      int       `form:"periods" binding:"required"`
	LoanType     int       `form:"loan_type" binding:"required"`
	StartDate    time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
}

type ModifyBasicInfoReq struct {
	Id           int64     `form:"id" binding:"required"`
	Name         string    `form:"name"`
	LoanId       int64     `form:"loan_id"`
	Principal    float64   `form:"principal"`
	InterestRate float64   `form:"interest_rate"`
	Periods      int       `form:"periods"`
	LoanType     int       `form:"loan_type"`
	StartDate    time.Time `form:"start_date" time_format:"2006-01-02"`
}

type CreateLoanReq struct {
	Name string `form:"name" binding:"required"`
}

type ModifyLoanReq struct {
	Id   int64  `form:"id" binding:"required"`
	Name string `form:"name" binding:"required"`
}

type CutInterestRateReq struct {
	Id            int64     `form:"id" binding:"required"`
	EffectiveDate time.Time `form:"effective_date" time_format:"2006-01-02"  binding:"required"`
	Rate          float64   `form:"rate" binding:"required"`
}
