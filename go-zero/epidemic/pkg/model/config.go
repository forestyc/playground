package model

type Config struct {
	SymmetricKey []byte // 对称加密秘钥
}

var config Config

func SetConfig(symmetricKey string) {
	config.SymmetricKey = []byte(symmetricKey)
}
