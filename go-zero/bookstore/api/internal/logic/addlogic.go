package logic

import (
	"context"
	"github.com/Baal19905/playground/go-zero/bookstore/rpc/add/adder"

	"github.com/Baal19905/playground/go-zero/bookstore/api/internal/svc"
	"github.com/Baal19905/playground/go-zero/bookstore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddResp, err error) {
	r, err := l.svcCtx.Adder.Add(l.ctx, &adder.AddReq{
		Book:  req.Book,
		Price: req.Price,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddResp{
		Ok: r.Ok,
	}, nil
}
