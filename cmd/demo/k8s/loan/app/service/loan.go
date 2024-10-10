package service

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/entity/db"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/model"
	"github.com/forestyc/playground/pkg/distributed/snowflake"
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

func (l *Loan) GetById(id int64) (db.Loan, error) {
	var list db.Loan
	session := l.ctx.Db.Session()
	if result := session.Where("id=?", id).Take(&list); result.Error != nil {
		return list, result.Error
	}
	return list, nil
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
