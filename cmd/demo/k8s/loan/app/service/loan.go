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

func (l *Loan) Take(id int64) (db.LoanBasicInfo, error) {
	loanBasicInfo := db.LoanBasicInfo{}
	session := l.ctx.Db.Session()
	if err := session.Where("id=?", id).Take(&loanBasicInfo).Error; err != nil {
		return db.LoanBasicInfo{}, err
	}
	return loanBasicInfo, nil
}

func (l *Loan) Create(req model.CreateBasicInfoReq) error {
	loanBasicInfo := db.LoanBasicInfo{
		Id:           l.snowflake.Gen(),
		LoanId:       req.LoanId,
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

func (l *Loan) Modify(req model.ModifyBasicInfoReq) error {
	loanBasicInfo := db.LoanBasicInfo{
		Id:           req.Id,
		LoanId:       req.LoanId,
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
		if err = tx.Updates(&loanBasicInfo).Error; err != nil {
			return err
		}
		// remove repayment
		if err = tx.Where("loan_basic_info_id=?", loanBasicInfo.Id).Delete(&db.Repayment{}).Error; err != nil {
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
			Id:              l.snowflake.Gen(),
			LoanBasicInfoId: loanBasicInfo.Id,
			Period:          period,
			Amount:          amount,
			RepaymentDate:   date,
		})
	}
	return repayments, nil
}
