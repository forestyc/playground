// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"github.com/forestyc/playground/go-zero/epidemic/rpc/user/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReq           = user.LoginReq
	LoginResp          = user.LoginResp
	MobileFindUserReq  = user.MobileFindUserReq
	MobileFindUserResp = user.MobileFindUserResp
	MsgCodeReq         = user.MsgCodeReq
	MsgCodeResp        = user.MsgCodeResp

	User interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		MsgCode(ctx context.Context, in *MsgCodeReq, opts ...grpc.CallOption) (*MsgCodeResp, error)
		MobileFindUser(ctx context.Context, in *MobileFindUserReq, opts ...grpc.CallOption) (*MobileFindUserResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) MsgCode(ctx context.Context, in *MsgCodeReq, opts ...grpc.CallOption) (*MsgCodeResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.MsgCode(ctx, in, opts...)
}

func (m *defaultUser) MobileFindUser(ctx context.Context, in *MobileFindUserReq, opts ...grpc.CallOption) (*MobileFindUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.MobileFindUser(ctx, in, opts...)
}