package model

import "time"

type CreateInfoReq struct {
	Name         string    `form:"name" binding:"required"`
	Principal    float64   `form:"principal" binding:"required"`
	InterestRate float64   `form:"interest_rate" binding:"required"`
	Periods      int       `form:"periods" binding:"required"`
	LoanType     int       `form:"loan_type" binding:"required"`
	StartDate    time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
}
