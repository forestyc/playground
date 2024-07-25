package service

import (
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/entity/db"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/model"
	"gorm.io/gorm"
)

type Info struct {
	ctx *context.Context
	pi  *PrincipalInterest
}

func NewInfo(ctx *context.Context) *Info {
	return &Info{
		ctx: ctx,
	}
}

func (i *Info) Get(id int) ([]db.ConfigInfo, error) {
	var info []db.ConfigInfo
	var err error
	if id > 0 {
		if err = i.ctx.Db.Session().
			Model(&info).
			Where("id=?", id).
			Take(&info).Error; err != nil {
			return info, err
		}
	} else {
		if err = i.ctx.Db.Session().
			Model(&info).
			Find(&info).Error; err != nil {
			return info, err
		}
	}
	return info, nil
}

func (i *Info) Set(req model.SetInfoReq) error {
	if req.Id == 0 {
		return i.create(req)
	} else {
		return i.modify(req)
	}
}

func (i *Info) create(req model.SetInfoReq) error {
	entity := db.ConfigInfo{
		Name:                      &req.Name,
		BusinessPrincipal:         &req.BusinessPrincipal,
		BusinessInterestRate:      &req.BusinessInterestRate,
		BusinessPeriods:           &req.BusinessPeriods,
		ProvidentFundPrincipal:    &req.ProvidentFundPrincipal,
		ProvidentFundInterestRate: &req.ProvidentFundInterestRate,
		ProvidentFundPeriods:      &req.ProvidentFundPeriods,
		StartDate:                 &req.StartDate,
	}
	providentFundLoan := NewPrincipalInterest(
		req.ProvidentFundPrincipal,
		req.ProvidentFundInterestRate,
		req.ProvidentFundPeriods,
	)

	businessLoan := NewPrincipalInterest(
		req.ProvidentFundPrincipal,
		req.ProvidentFundInterestRate,
		req.ProvidentFundPeriods,
	)

	db := i.ctx.Db.Session()

	return db.Transaction(func(tx *gorm.DB) error {
		// save loan info
		if err := tx.Create(&entity).Error; err != nil {
			return err
		}
		businessLoan.Repayment()
		return nil
	})
}

func (i *Info) modify(req model.SetInfoReq) error {
	entity := db.ConfigInfo{
		Id:                        &req.Id,
		Name:                      &req.Name,
		BusinessPrincipal:         &req.BusinessPrincipal,
		BusinessInterestRate:      &req.BusinessInterestRate,
		BusinessPeriods:           &req.BusinessPeriods,
		ProvidentFundPrincipal:    &req.ProvidentFundPrincipal,
		ProvidentFundInterestRate: &req.ProvidentFundInterestRate,
		ProvidentFundPeriods:      &req.ProvidentFundPeriods,
		StartDate:                 &req.StartDate,
	}
	return i.ctx.Db.Session().Updates(&entity).Error
}
