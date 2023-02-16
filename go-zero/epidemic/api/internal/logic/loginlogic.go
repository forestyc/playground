package logic

import (
	"context"
	"github.com/Baal19905/playground/go-zero/epidemic/rpc/user/user"
	"time"

	"github.com/Baal19905/playground/go-zero/epidemic/api/internal/svc"
	"github.com/Baal19905/playground/go-zero/epidemic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	resp = &types.LoginResp{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	rpcReq := user.LoginReq{
		Mobile: req.Mobile,
		Device: req.Device,
		Code:   req.Code,
	}
	rpcResp := &user.LoginResp{}
	rpcResp, err = l.svcCtx.UserRpc.Login(ctx, &rpcReq)
	if err != nil {
		logx.Error("msg-code failed", err)
		resp.Message = err.Error()
		return
	}
	resp.Message = rpcResp.Message
	resp.Data.AccessToken = rpcResp.AccessToken
	resp.Data.RefreshToken = rpcResp.RefreshToken
	return
}
