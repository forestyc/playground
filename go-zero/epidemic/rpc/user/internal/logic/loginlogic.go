package logic

import (
	"context"
	"errors"
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/model"
	"github.com/Baal19905/playground/go-zero/epidemic/rpc/user/internal/svc"
	"github.com/Baal19905/playground/go-zero/epidemic/rpc/user/user"
	"gorm.io/gorm"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	var err error
	var resp *user.LoginResp = &user.LoginResp{}
	//if err = l.svcCtx.MsgCode.Check(in.Mobile, in.Code); err != nil {
	//	logx.Error("登录失败，mobile=", in.Mobile, err)
	//	return resp, err
	//}
	user := model.EpidemicUser{}
	session := l.svcCtx.Mysql.Session()
	result := session.Table(user.TableName()).Where("mobile=?", in.Mobile).Find(&user)
	if err = result.Error; err != nil {
		logx.Error("登录失败，mobile=", in.Mobile, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Message = "用户不存在"
		} else {
			resp.Message = "登录失败"
		}
		return resp, err
	}

	// 生成token并返回
	resp.AccessToken, resp.RefreshToken, err = l.svcCtx.Token.GenToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		logx.Error("登录失败，mobile=", in.Mobile, err)
		resp.Message = "登录失败"
		return resp, err
	}
	return resp, nil
}
