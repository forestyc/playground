package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"github.com/forestyc/playground/pkg/encoding"
	"github.com/pkg/errors"
)

type RSA struct {
}

// GenerateKey 生成密钥对，返回私钥、公钥、错误
func (r RSA) GenerateKey() ([]byte, []byte, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}
	pubStream, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	// private
	privPem, _ := encoding.PemEncode("RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(privKey))
	// public
	pubPem, _ := encoding.PemEncode("PUBLIC KEY", pubStream)
	return privPem, pubPem, nil
}

// Encrypt 加密(公钥)
func (r RSA) Encrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	pub, err := r.parsePublic(key)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, data, nil)
}

// Decrypt 解密(私钥)
func (r RSA) Decrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	priv, err := r.parsePrivate(key)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, data, nil)
}

// EncryptWithBase64 加密(公钥)，使用base64
func (r RSA) EncryptWithBase64(data []byte, key []byte) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New("invalid params")
	}
	ciphertext, err := r.Encrypt(data, key)
	if err != nil {
		return "", err
	}
	return encoding.Base64Encode(ciphertext), nil
}

// DecryptWithBase64 解密(私钥)，使用base64
func (r RSA) DecryptWithBase64(data string, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	plaintext, err := encoding.Base64Decode(data)
	if err != nil {
		return nil, err
	}
	return r.Decrypt(plaintext, key)
}

// Sign 签名（私钥）
func (r RSA) Sign(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	priv, err := r.parsePrivate(key)
	if err != nil {
		return nil, err
	}
	sh256 := sha256.New()
	sh256.Write(data)
	return rsa.SignPSS(rand.Reader, priv, crypto.SHA256, sh256.Sum(nil), nil)
}

// Verify 校验（公钥）
func (r RSA) Verify(data []byte, key []byte, sign []byte) (bool, error) {
	if len(data) == 0 || len(key) == 0 || len(sign) == 0 {
		return false, errors.New("invalid params")
	}
	pub, err := r.parsePublic(key)
	if err != nil {
		return false, err
	}
	sh256 := sha256.New()
	sh256.Write(data)
	if err = rsa.VerifyPSS(pub, crypto.SHA256, sh256.Sum(nil), sign, nil); err != nil {
		return false, err
	}
	return true, nil
}

// SignWithBase64 签名（私钥），使用base64
func (r RSA) SignWithBase64(data []byte, key []byte) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New("invalid params")
	}
	sign, err := r.Sign(data, key)
	if err != nil {
		return "", err
	}
	return encoding.Base64Encode(sign), nil
}

// VerifyWithBase64 校验（公钥），使用base64
func (r RSA) VerifyWithBase64(data []byte, key []byte, sign string) (bool, error) {
	if len(data) == 0 || len(key) == 0 {
		return false, errors.New("invalid params")
	}
	signRaw, err := encoding.Base64Decode(sign)
	if err != nil {
		return false, err
	}
	return r.Verify(data, key, signRaw)
}

// 解析私钥
func (r RSA) parsePrivate(key []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(encoding.PemDecode(key))
}

// 解析公钥
func (r RSA) parsePublic(key []byte) (*rsa.PublicKey, error) {
	pub, err := x509.ParsePKIXPublicKey(encoding.PemDecode(key))
	return pub.(*rsa.PublicKey), err
}
