package model

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/entity/db"
)

type GetBasicInfoResp struct {
	BasicInfo     db.LoanBasicInfo `json:"basic_info"`
	RepaymentList []db.Repayment   `json:"repayment_list"`
}

type repayment struct {
	Amount        float64 `json:"amount"`
	RepaymentDate string  `json:"repayment_date"`
	Name          string  `json:"name"`
}

type GetLoanInfoResp struct {
	LoanInfo      db.Loan     `json:"loan_info"`
	RepaymentList []repayment `json:"repayment_list"`
}
