package logic

import (
	"context"
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/model"
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
	payload, err := l.svcCtx.Token.ValidateAccessToken(in.Token)
	if err != nil {
		logx.Error("MobileFindUser invalid token", err)
		return nil, err
	}
	session, cancel := l.svcCtx.Mysql.Session()
	defer cancel()
	userInfo := model.EpidemicUser{}
	result := session.Table(userInfo.TableName()).Where("id = ?", payload.ID).Find(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	resp := &user.MobileFindUserResp{
		Id:                       userInfo.ID,
		Name:                     string(userInfo.Name),
		Sex:                      userInfo.Sex,
		Mobile:                   userInfo.Mobile,
		BirthDate:                userInfo.BirthDate.Format("2006-01-02"),
		Company:                  userInfo.Company,
		Department:               userInfo.Department,
		Organization:             userInfo.Organization,
		Location:                 userInfo.Location,
		ResidentialAddressCode:   string(userInfo.ResidentialAddressCode),
		ResidentialAddressName:   string(userInfo.ResidentialAddressName),
		ResidentialAddressDetail: string(userInfo.ResidentialAddressDetail),
		VaccinationRecord:        userInfo.VaccinationRecord,
		VaccinationTime:          userInfo.VaccinationTime,
		ReasonForNonVaccination:  userInfo.ReasonForNonVaccination,
		Role:                     userInfo.Role,
		Status:                   userInfo.Status,
		ApprovedBy:               userInfo.ApprovedBy,
		ApprovalComments:         userInfo.ApprovalComments,
		ApprovalDate:             userInfo.ApprovalDate.Format("2006-01-02"),
		Label:                    userInfo.Label,
	}
	return resp, nil
}
