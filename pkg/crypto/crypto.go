package crypto

// Crypto 加密工厂
type Crypto struct {
	Symmetric  SymmetricCrypto  // 对称加密
	Asymmetric AsymmetricCrypto // 非对称加密
	Digest     MessageDigest    // 信息摘要
}

// SymmetricCrypto 对称加密接口
type SymmetricCrypto interface {
	Encrypt(data []byte, key []byte) ([]byte, error)           // 加密
	Decrypt(data []byte, key []byte) ([]byte, error)           // 解密
	EncryptWithBase64(data []byte, key []byte) (string, error) // 加密，使用base64
	DecryptWithBase64(data string, key []byte) ([]byte, error) // 解密，使用base64
}

// AsymmetricCrypto 非对称加密接口
type AsymmetricCrypto interface {
	GenerateKey(pub string, priv string) error                           // 生成密钥对，输入公钥、私钥文件路径
	GeneratePublicKey(priv []byte) (string, error)                       // 生成公钥
	Encrypt(data []byte, key []byte) ([]byte, error)                     // 加密(公钥)
	Decrypt(data []byte, key []byte) ([]byte, error)                     // 解密(私钥)
	EncryptWithBase64(data []byte, key []byte) (string, error)           // 加密(公钥)，使用base64
	DecryptWithBase64(data string, key []byte) ([]byte, error)           // 解密(私钥)，使用base64
	Sign(data []byte, key []byte) ([]byte, error)                        // 签名（私钥）
	Verify(data []byte, key []byte, sign []byte) (bool, error)           // 校验（公钥）
	SignWithBase64(data []byte, key []byte) (string, error)              // 签名（私钥），使用base64
	VerifyWithBase64(data []byte, key []byte, sign string) (bool, error) // 校验（公钥），使用base64
}

// MessageDigest 信息摘要
type MessageDigest interface {
	Sum(data []byte) (string, error) // 摘要
}
