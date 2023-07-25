package logic

import (
	"context"
	"github.com/forestyc/playground/go-zero/epidemic/rpc/user/user"
	"time"

	"github.com/forestyc/playground/go-zero/epidemic/api/internal/svc"
	"github.com/forestyc/playground/go-zero/epidemic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgCodeLogic {
	return &MsgCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgCodeLogic) MsgCode(req *types.CommonMsgCodeReq) (resp *types.CommonMsgCodeResp, err error) {
	resp = &types.CommonMsgCodeResp{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	rpcReq := user.MsgCodeReq{Mobile: req.Mobile}
	rpcResp := &user.MsgCodeResp{}
	rpcResp, err = l.svcCtx.UserRpc.MsgCode(ctx, &rpcReq)
	if err != nil {
		logx.Error("msg-code failed", err)
		resp.Message = err.Error()
		return
	}
	resp.Message = rpcResp.Message
	return
}
