package model

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/entity/db"
)

type GetBasicInfoResp struct {
	BasicInfo     db.LoanBasicInfo
	RepaymentList []db.Repayment
}
