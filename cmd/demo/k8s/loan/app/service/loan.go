package service

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/algorithm/loan"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/entity/db"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/model"
	"github.com/forestyc/playground/pkg/distributed/snowflake"
	"gorm.io/gorm"
)

type Loan struct {
	ctx       *context.Context
	snowflake *snowflake.Snowflake
}

func NewLoan(ctx *context.Context) *Loan {
	return &Loan{
		ctx:       ctx,
		snowflake: snowflake.New(ctx.C.Server.Id),
	}
}

func (l *Loan) Create(req model.CreateInfoReq) error {
	loanBasicInfo := db.LoanBasicInfo{
		Id:           l.snowflake.Gen(),
		Principal:    req.Principal,
		LoanType:     req.LoanType,
		InterestRate: req.InterestRate,
		Periods:      req.Periods,
		StartDate:    req.StartDate,
	}
	repayments, err := l.createRepaymentList(loanBasicInfo)
	if err != nil {
		return err
	}
	session := l.ctx.Db.Session()

	return session.Transaction(func(tx *gorm.DB) error {
		// save loan basic info
		if err = tx.Create(&loanBasicInfo).Error; err != nil {
			return err
		}
		// save repayment
		if err = tx.Create(&repayments).Error; err != nil {
			return err
		}
		return nil
	})
}

func (l *Loan) createRepaymentList(loanBasicInfo db.LoanBasicInfo) ([]db.Repayment, error) {
	loanAlgorithm, err := loan.NewLoan(loanBasicInfo.LoanType, loan.BasicInfo{
		Principal:    loanBasicInfo.Principal,
		InterestRate: loanBasicInfo.InterestRate,
		Periods:      loanBasicInfo.Periods,
		StartDate:    loanBasicInfo.StartDate,
	})
	if err != nil {
		return nil, err
	}
	var repayments []db.Repayment
	for i := 0; i < loanBasicInfo.Periods; i++ {
		amount, date := loanAlgorithm.Repayment(i)
		period := i + 1
		repayments = append(repayments, db.Repayment{
			Id:            l.snowflake.Gen(),
			LoanId:        loanBasicInfo.Id,
			Period:        period,
			Amount:        amount,
			RepaymentDate: date,
		})
	}
	return repayments, nil
}
