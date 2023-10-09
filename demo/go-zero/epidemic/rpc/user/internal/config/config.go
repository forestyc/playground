package config

import (
	"github.com/forestyc/playground/go-zero/epidemic/pkg/gorm"
	"github.com/forestyc/playground/go-zero/epidemic/pkg/msgcode"
	"github.com/forestyc/playground/go-zero/epidemic/pkg/sms"
	"github.com/forestyc/playground/go-zero/epidemic/pkg/token"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CodeGenerator msgcode.Params
	Sms           sms.Params
	Mysql         gorm.MysqlConf
	Token         token.Token
	Crypto        Crypto
}

type Crypto struct {
	SymmetricKey string // 对称加密秘钥
}
