package epidemic_mobile

import (
	"context"
	"github.com/forestyc/playground/go-zero/epidemic/rpc/user/user"
	"time"

	"github.com/forestyc/playground/go-zero/epidemic/api/internal/svc"
	"github.com/forestyc/playground/go-zero/epidemic/api/internal/types"

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

func (l *MobileUserFindLogic) MobileUserFind(token string) (resp *types.MobileUserFindResp, err error) {
	rpcReq := user.MobileFindUserReq{
		Token: token,
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	rpcResp := &user.MobileFindUserResp{}
	if rpcResp, err = l.svcCtx.UserRpc.MobileFindUser(ctx, &rpcReq); err != nil {
		return
	}
	resp = &types.MobileUserFindResp{}
	resp.ID = int(rpcResp.Id)
	resp.Name = rpcResp.Name
	resp.Sex = rpcResp.Sex
	resp.Mobile = rpcResp.Mobile
	resp.BirthDate = rpcResp.BirthDate
	resp.Company = rpcResp.Company
	resp.Department = rpcResp.Department
	resp.Organization = rpcResp.Organization
	resp.Location = rpcResp.Location
	resp.ResidentialAddressCode = rpcResp.ResidentialAddressCode
	resp.ResidentialAddressName = rpcResp.ResidentialAddressName
	resp.ResidentialAddressDetail = rpcResp.ResidentialAddressDetail
	resp.VaccinationRecord = rpcResp.VaccinationRecord
	resp.VaccinationTime = rpcResp.VaccinationTime
	resp.ReasonForNonVaccination = rpcResp.ReasonForNonVaccination
	resp.Role = rpcResp.Role
	resp.Status = rpcResp.Status
	resp.ApprovedBy = rpcResp.ApprovedBy
	resp.ApprovalComments = rpcResp.ApprovalComments
	resp.ApprovalDate = rpcResp.ApprovalDate
	resp.Label = rpcResp.Label
	return
}
