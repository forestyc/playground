package model

import "time"

type CreateBasicInfoReq struct {
	LoanId       int64     `form:"loan_id" binding:"required"`
	Principal    float64   `form:"principal" binding:"required"`
	InterestRate float64   `form:"interest_rate" binding:"required"`
	Periods      int       `form:"periods" binding:"required"`
	LoanType     int       `form:"loan_type" binding:"required"`
	StartDate    time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
}

type ModifyBasicInfoReq struct {
	Id           int64     `form:"id" binding:"required"`
	LoanId       int64     `form:"loan_id"`
	Principal    float64   `form:"principal"`
	InterestRate float64   `form:"interest_rate"`
	Periods      int       `form:"periods"`
	LoanType     int       `form:"loan_type"`
	StartDate    time.Time `form:"start_date" time_format:"2006-01-02"`
}
