package model

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/entity/db"
)

type GetBasicInfoResp struct {
	BasicInfo     db.LoanBasicInfo
	RepaymentList []db.Repayment
}

type repayment struct {
	Amount        float64 `json:"amount"`
	RepaymentDate string  `json:"repayment_date"`
	Name          string  `json:"name"`
}

type GetLoanInfoResp struct {
	LoanInfo      db.Loan
	RepaymentList []repayment
}
