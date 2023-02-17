package logic

import (
	"context"

	"github.com/Baal19905/playground/go-zero/epidemic/rpc/user/internal/svc"
	"github.com/Baal19905/playground/go-zero/epidemic/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type MobileFindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMobileFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MobileFindUserLogic {
	return &MobileFindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MobileFindUserLogic) MobileFindUser(in *user.MobileFindUserReq) (*user.MobileFindUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.MobileFindUserResp{}, nil
}
