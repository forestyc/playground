package service

import (
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/houseloan/app/entity/db"
)

type Info struct {
	ctx *context.Context
}

func NewInfo(ctx *context.Context) *Info {
	return &Info{
		ctx: ctx,
	}
}

func (i *Info) Get(id int) ([]db.Info, error) {
	var info []db.Info
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
