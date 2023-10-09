package logic

import (
	"context"
	"github.com/forestyc/playground/go-zero/bookstore/rpc/check/checker"

	"github.com/forestyc/playground/go-zero/bookstore/api/internal/svc"
	"github.com/forestyc/playground/go-zero/bookstore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req *types.CheckReq) (resp *types.CheckResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckReq{
		Book: req.Book,
	})
	if err != nil {
		logx.Error(err)
		return &types.CheckResp{}, err
	}

	return &types.CheckResp{
		Found: r.Found,
		Price: r.Price,
	}, nil
}
