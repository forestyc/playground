// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/Baal19905/playground/go-zero/epidemic/rpc/user/internal/logic"
	"github.com/Baal19905/playground/go-zero/epidemic/rpc/user/internal/svc"
	"github.com/Baal19905/playground/go-zero/epidemic/rpc/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) MsgCode(ctx context.Context, in *user.MsgCodeReq) (*user.MsgCodeResp, error) {
	l := logic.NewMsgCodeLogic(ctx, s.svcCtx)
	return l.MsgCode(in)
}

func (s *UserServer) MobileFindUser(ctx context.Context, in *user.MobileFindUserReq) (*user.MobileFindUserResp, error) {
	l := logic.NewMobileFindUserLogic(ctx, s.svcCtx)
	return l.MobileFindUser(in)
}