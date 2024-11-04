package service

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/algorithm/loan"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/entity/db"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/model"
	"github.com/forestyc/playground/pkg/distributed/snowflake"
	"gorm.io/gorm"
)

const (
	condLoanBasicInfoId = "loan_basic_info_id=?"
)

type LoanBasicInfo struct {
	ctx       *context.Context
	snowflake *snowflake.Snowflake
}

func NewLoanBasicInfo(ctx *context.Context) *LoanBasicInfo {
	return &LoanBasicInfo{
		ctx:       ctx,
		snowflake: snowflake.New(ctx.C.Server.Id),
	}
}

func (l *LoanBasicInfo) GetByLoanId(loanId int64) ([]db.LoanBasicInfo, error) {
	var loanBasicInfo []db.LoanBasicInfo
	session := l.ctx.Db.Session()
	if err := session.Where("loan_id=?", loanId).Find(&loanBasicInfo).Error; err != nil {
		return nil, err
	}
	return loanBasicInfo, nil
}

func (l *LoanBasicInfo) GetById(id int64) (model.GetBasicInfoResp, error) {
	var resp model.GetBasicInfoResp
	session := l.ctx.Db.Session()
	if err := session.Where("id=?", id).Take(&resp.BasicInfo).Error; err != nil {
		return resp, err
	}
	if err := session.Where(condLoanBasicInfoId, id).Find(&resp.RepaymentList).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func (l *LoanBasicInfo) Delete(id int64) error {
	session := l.ctx.Db.Session()
	return session.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = session.Where("id=?", id).Delete(&db.LoanBasicInfo{}).Error; err != nil {
			return err
		}
		if err = session.Where(condLoanBasicInfoId, id).Delete(&db.Repayment{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (l *LoanBasicInfo) Create(req model.CreateBasicInfoReq) error {
	loanBasicInfo := db.LoanBasicInfo{
		Id:           l.snowflake.Gen(),
		Name:         req.Name,
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

func (l *LoanBasicInfo) Modify(req model.ModifyBasicInfoReq) error {
	loanBasicInfo := db.LoanBasicInfo{
		Id:           req.Id,
		Name:         req.Name,
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
		if err = tx.Where(condLoanBasicInfoId, loanBasicInfo.Id).Delete(&db.Repayment{}).Error; err != nil {
			return err
		}
		// save repayment
		if err = tx.Create(&repayments).Error; err != nil {
			return err
		}
		return nil
	})
}

func (l *LoanBasicInfo) createRepaymentList(loanBasicInfo db.LoanBasicInfo) ([]db.Repayment, error) {
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
