package epidemic_mobile

import (
	"context"

	"github.com/Baal19905/playground/go-zero/epidemic/api/internal/svc"
	"github.com/Baal19905/playground/go-zero/epidemic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MobileUserFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMobileUserFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MobileUserFindLogic {
	return &MobileUserFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MobileUserFindLogic) MobileUserFind() (resp *types.MobileUserFindResp, err error) {
	// todo: add your logic here and delete this line

	return
}
