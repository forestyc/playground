package logic

import (
	"context"
	"errors"

	"github.com/forestyc/playground/go-zero/epidemic/rpc/user/internal/svc"
	"github.com/forestyc/playground/go-zero/epidemic/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMsgCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgCodeLogic {
	return &MsgCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MsgCodeLogic) MsgCode(in *user.MsgCodeReq) (*user.MsgCodeResp, error) {
	out := &user.MsgCodeResp{}
	if len(in.Mobile) == 0 {
		return out, errors.New("invalid mobile[" + in.Mobile + "]")
	}
	code, err := l.svcCtx.MsgCode.Gen(in.Mobile)
	if err != nil {
		return out, err
	}
	if err = l.svcCtx.Sms.SendMsg([]string{in.Mobile}, code); err != nil {
		return out, err
	}
	return &user.MsgCodeResp{
		Message: "操作成功",
	}, nil
}
