package crypto

import (
	"crypto/rand"
	"errors"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"io/ioutil"
)

type SM2 struct {
}

// Encrypt 加密(公钥)
func (s SM2) Encrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	pub, err := x509.ReadPublicKeyFromHex(string(key))
	if err != nil {
		return nil, err
	}
	crypto, err := pub.EncryptAsn1(data, rand.Reader)
	if err != nil {
		return nil, err
	}
	return crypto, nil
}

// Decrypt 解密(私钥)
func (s SM2) Decrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	private, err := x509.ReadPrivateKeyFromHex(string(key))
	if err != nil {
		return nil, err
	}
	return private.DecryptAsn1(data)
}

// EncryptWithBase64 加密(公钥)，使用base64
func (s SM2) EncryptWithBase64(data []byte, key []byte) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New("invalid params")
	}
	crypto, err := s.Encrypt(data, key)
	if err != nil {
		return "", err
	}
	return Base64Encode(crypto), nil
}

// DecryptWithBase64 解密(私钥)，使用base64
func (s SM2) DecryptWithBase64(data string, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	raw, err := Base64Decode(data)
	if err != nil {
		return nil, err
	}
	return s.Decrypt(raw, key)
}

// Sign 签名（私钥）
func (s SM2) Sign(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	private, err := x509.ReadPrivateKeyFromHex(string(key))
	if err != nil {
		return nil, err
	}
	return private.Sign(rand.Reader, data, nil)
}

// Verify 校验（公钥）
func (s SM2) Verify(data []byte, key []byte, sign []byte) (bool, error) {
	if len(data) == 0 || len(key) == 0 {
		return false, errors.New("invalid params")
	}
	pub, err := x509.ReadPublicKeyFromHex(string(key))
	if err != nil {
		return false, err
	}
	return pub.Verify(data, sign), nil
}

// SignWithBase64 签名（私钥），使用base64
func (s SM2) SignWithBase64(data []byte, key []byte) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New("invalid params")
	}
	sign, err := s.Sign(data, key)
	if err != nil {
		return "", err
	}
	return Base64Encode(sign), nil
}

// VerifyWithBase64 校验（公钥），使用base64
func (s SM2) VerifyWithBase64(data []byte, key []byte, sign string) (bool, error) {
	if len(data) == 0 || len(key) == 0 {
		return false, errors.New("invalid params")
	}
	signRaw, err := Base64Decode(sign)
	if err != nil {
		return false, err
	}
	return s.Verify(data, key, signRaw)
}

// GenerateKey 生成密钥对，返回公钥、私钥文件路径
func (s SM2) GenerateKey(pubKeyPath, privKeyPath string) error {
	if pubKeyPath == "" || privKeyPath == "" {
		return errors.New("invalid params")
	}
	priv, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
	if err != nil {
		return err
	}
	public := priv.Public().(*sm2.PublicKey)
	pubkeyHex := x509.WritePublicKeyToHex(public)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(pubKeyPath, []byte(pubkeyHex), 0644)
	if err != nil {
		return err
	}
	privkeyHex := x509.WritePrivateKeyToHex(priv)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(privKeyPath, []byte(privkeyHex), 0644)
	if err != nil {
		return err
	}
	return nil
}

// GeneratePublicKey 生成公钥(hex)
func (s SM2) GeneratePublicKey(priv []byte) (string, error) {
	private, err := x509.ReadPrivateKeyFromHex(string(priv))
	if err != nil {
		return "", err
	}
	public := private.Public().(*sm2.PublicKey)
	hex := x509.WritePublicKeyToHex(public)
	if err != nil {
		return "", err
	}
	return hex, nil
}
