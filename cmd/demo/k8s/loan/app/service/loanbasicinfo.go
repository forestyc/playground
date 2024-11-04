package service

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/algorithm/loan"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/entity/db"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/model"
	"github.com/forestyc/playground/pkg/distributed/snowflake"
	"github.com/forestyc/playground/pkg/utils"
	"gorm.io/gorm"
	"time"
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
	repayments, err := l.createRepaymentList(loanBasicInfo, 0)
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
	var loanBasicInfo db.LoanBasicInfo
	session := l.ctx.Db.Session()
	if err := session.Where("id=?", req.Id).Take(&loanBasicInfo).Error; err != nil {
		return err
	}

	if req.Name != "" {
		loanBasicInfo.Name = req.Name
	}
	if req.LoanId != 0 {
		loanBasicInfo.LoanId = req.LoanId
	}
	if req.Periods != 0 {
		loanBasicInfo.Periods = req.Periods
	}
	if !utils.EqualFloat(req.Principal, 0, loan.EPSILON) {
		loanBasicInfo.Principal = req.Principal
	}
	if !utils.EqualFloat(req.InterestRate, 0, loan.EPSILON) {
		loanBasicInfo.InterestRate = req.InterestRate
	}
	if req.LoanType != 0 {
		loanBasicInfo.LoanType = req.LoanType
	}
	if !req.StartDate.IsZero() {
		loanBasicInfo.StartDate = req.StartDate
	}
	repayments, err := l.createRepaymentList(loanBasicInfo, 0)
	if err != nil {
		return err
	}

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

func (l *LoanBasicInfo) CutInterestRate(req model.CutInterestRateReq) error {
	var loanBasicInfo db.LoanBasicInfo
	session := l.ctx.Db.Session()
	if err := session.Where("id=?", req.Id).Take(&loanBasicInfo).Error; err != nil {
		return err
	}

	var oldPeriod int64
	if err := session.Table("repayment").
		Where("repayment_date<? and loan_basic_info_id=?", req.EffectiveDate, req.Id).
		Count(&oldPeriod).Error; err != nil {
		return err
	}

	loanBasicInfo.InterestRate = req.Rate

	repayments, err := l.createRepaymentList(loanBasicInfo, int(oldPeriod))
	if err != nil {
		return err
	}

	// amend first repayment
	repayments[0].Amount = l.amendRepaymentAmount(repayments[0], req.EffectiveDate)

	return session.Transaction(func(tx *gorm.DB) error {
		// save loan basic info
		if err = tx.Updates(&loanBasicInfo).Error; err != nil {
			return err
		}
		// remove repayment
		if err = tx.Where("repayment_date>=? and loan_basic_info_id=?", repayments[0].RepaymentDate, req.Id).
			Delete(&db.Repayment{}).Error; err != nil {
			return err
		}
		// save repayment
		if err = tx.Create(&repayments).Error; err != nil {
			return err
		}
		return nil
	})
}

func (l *LoanBasicInfo) createRepaymentList(loanBasicInfo db.LoanBasicInfo, offset int) ([]db.Repayment, error) {
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
		if i > offset-1 {
			repayments = append(repayments, db.Repayment{
				Id:              l.snowflake.Gen(),
				LoanBasicInfoId: loanBasicInfo.Id,
				Period:          i + 1,
				Amount:          amount,
				RepaymentDate:   date,
			})
		}
	}
	return repayments, nil
}

func (l *LoanBasicInfo) amendRepaymentAmount(repayment db.Repayment, effectiveDate time.Time) float64 {
	var startDate, endDate time.Time
	endDate = repayment.RepaymentDate
	startDate = repayment.RepaymentDate.AddDate(0, -1, 1)

	dayDiff := float64(endDate.Sub(startDate).Hours() / 24)
	beforeCutPercent := float64(effectiveDate.Sub(startDate).Hours()/24+1) / dayDiff
	afterCutPercent := float64(endDate.Sub(effectiveDate).Hours()/24-1) / dayDiff
	var oldRepayment db.Repayment
	l.ctx.Db.Session().
		Where("loan_basic_info_id=? and period=?", repayment.LoanBasicInfoId, repayment.Period).
		Take(&oldRepayment)

	repayment.Amount = oldRepayment.Amount*beforeCutPercent + afterCutPercent*repayment.Amount
	return repayment.Amount
}
