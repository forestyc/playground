package config

import (
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/gorm"
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/msgcode"
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/sms"
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/token"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CodeGenerator msgcode.Params
	Sms           sms.Params
	Mysql         gorm.MysqlConf
	Token         token.Token
}
