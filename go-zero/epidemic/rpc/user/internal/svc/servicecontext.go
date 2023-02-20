package svc

import (
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/gorm"
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/model"
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/msgcode"
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/sms"
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/token"
	"github.com/Baal19905/playground/go-zero/epidemic/rpc/user/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config  config.Config
	MsgCode msgcode.Generator // 验证码生成器
	Sms     sms.Sms           // 短信
	Mysql   gorm.Mysql        // mysql
	Token   token.Token       // token
}

func NewServiceContext(c config.Config) *ServiceContext {
	model.SetConfig(c.Crypto.SymmetricKey)
	return &ServiceContext{
		Config: c,
		MsgCode: msgcode.NewRedisMsgCode(c.CodeGenerator, redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
		})),
		Sms:   sms.NewHuaweiCloud(c.Sms),
		Mysql: gorm.NewMysql(c.Mysql),
		Token: c.Token,
	}
}
