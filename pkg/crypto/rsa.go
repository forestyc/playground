package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

type RSA struct {
}

// GenerateKey 生成密钥对，输入公钥、私钥文件路径
func (r RSA) GenerateKey(pub, priv string) error {
	if len(pub) == 0 || len(priv) == 0 {
		return errors.New("invalid params")
	}
	privKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}
	privStream := x509.MarshalPKCS1PrivateKey(privKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privStream,
	}
	file, err := os.Create(priv)
	if err != nil {
		return err
	}
	defer file.Close()
	pem.Encode(file, block)
	pubStream, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubStream,
	}
	file, err = os.Create(pub)
	if err != nil {
		return err
	}
	defer file.Close()
	pem.Encode(file, block)
	return nil
}

// GeneratePublicKey 生成公钥
func (r RSA) GeneratePublicKey(priv []byte) (string, error) {
	if len(priv) == 0 {
		return "", errors.New("invalid params")
	}
	privKey, err := r.parsePrivate(priv)
	if err != nil {
		return "", err
	}
	pubStream, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		return "", err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubStream,
	}
	pub := pem.EncodeToMemory(block)
	return string(pub), err
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
	return Base64Encode(ciphertext), nil
}

// DecryptWithBase64 解密(私钥)，使用base64
func (r RSA) DecryptWithBase64(data string, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	plaintext, err := Base64Decode(data)
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
	return Base64Encode(sign), nil
}

// VerifyWithBase64 校验（公钥），使用base64
func (r RSA) VerifyWithBase64(data []byte, key []byte, sign string) (bool, error) {
	if len(data) == 0 || len(key) == 0 {
		return false, errors.New("invalid params")
	}
	signRaw, err := Base64Decode(sign)
	if err != nil {
		return false, err
	}
	return r.Verify(data, key, signRaw)
}

// 解析私钥
func (r RSA) parsePrivate(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// 解析公钥
func (r RSA) parsePublic(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	return pub.(*rsa.PublicKey), err
}
