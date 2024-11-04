package service

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/entity/db"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/model"
	"github.com/forestyc/playground/pkg/distributed/snowflake"
	"time"
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

func (l *Loan) GetList() ([]db.Loan, error) {
	var list []db.Loan
	session := l.ctx.Db.Session()
	if result := session.Find(&list); result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}

func (l *Loan) GetById(id int64) (model.GetLoanInfoResp, error) {
	var resp model.GetLoanInfoResp
	session := l.ctx.Db.Session()
	if result := session.Where("id=?", id).Take(&resp.LoanInfo); result.Error != nil {
		return resp, result.Error
	}
	now := time.Now()
	firstDateOfMonth := now.AddDate(0, 0, -now.Day()+1)
	if result := session.Table("repayment r").
		Joins("inner join loan_basic_info l on l.id=r.loan_basic_info_id").
		Select("name", "amount", "repayment_date").
		Where("repayment_date>? and repayment_date<?",
			firstDateOfMonth.AddDate(0, 0, -1), firstDateOfMonth.AddDate(0, 1, 0)).
		Find(&resp.RepaymentList); result.Error != nil {
		return resp, result.Error
	}
	return resp, nil
}

func (l *Loan) Create(req model.CreateLoanReq) error {
	loan := db.Loan{
		Id:   l.snowflake.Gen(),
		Name: req.Name,
	}
	session := l.ctx.Db.Session()
	if result := session.Save(&loan); result.Error != nil {
		return result.Error
	}
	return nil
}

func (l *Loan) Delete(id int64) error {
	session := l.ctx.Db.Session()
	if result := session.Delete(&db.Loan{Id: id}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (l *Loan) Modify(req model.ModifyLoanReq) error {
	loan := db.Loan{
		Id:   req.Id,
		Name: req.Name,
	}
	session := l.ctx.Db.Session()
	if result := session.Updates(&loan); result.Error != nil {
		return result.Error
	}
	return nil
}
