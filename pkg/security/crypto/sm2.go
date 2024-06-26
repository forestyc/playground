package crypto

import (
	"bytes"
	"crypto/rand"
	"github.com/forestyc/playground/pkg/encoding/base64"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"math/big"
)

type SM2 struct {
	b64 base64.Base64
}

// Encrypt 加密(公钥)
func (s SM2) Encrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New(InvalidParameters)
	}
	pub, err := x509.ReadPublicKeyFromHex(string(key))
	if err != nil {
		return nil, err
	}
	crypto, err := sm2.Encrypt(pub, data, rand.Reader, sm2.C1C3C2)
	if err != nil {
		return nil, err
	}
	return crypto, nil
}

// Decrypt 解密(私钥)
func (s SM2) Decrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New(InvalidParameters)
	}
	private, err := x509.ReadPrivateKeyFromHex(string(key))
	if err != nil {
		return nil, err
	}
	return sm2.Decrypt(private, data, sm2.C1C3C2)
}

// EncryptWithBase64 加密(公钥)，使用base64
func (s SM2) EncryptWithBase64(data []byte, key []byte) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New(InvalidParameters)
	}
	crypto, err := s.Encrypt(data, key)
	if err != nil {
		return "", err
	}
	return s.b64.Encode(crypto), nil
}

// DecryptWithBase64 解密(私钥)，使用base64
func (s SM2) DecryptWithBase64(data string, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New(InvalidParameters)
	}
	raw, err := s.b64.Decode(data)
	if err != nil {
		return nil, err
	}
	return s.Decrypt(raw, key)
}

// Sign 签名（私钥）
func (s SM2) Sign(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New(InvalidParameters)
	}
	private, err := x509.ReadPrivateKeyFromHex(string(key))
	if err != nil {
		return nil, err
	}
	R, S, err := sm2.Sm2Sign(private, data, nil, rand.Reader)
	if err != nil {
		return nil, err
	}
	return bytes.Join([][]byte{R.Bytes(), S.Bytes()}, nil), nil
}

// Verify 校验（公钥）
func (s SM2) Verify(data []byte, key []byte, sign []byte) (bool, error) {
	if len(data) == 0 || len(key) == 0 {
		return false, errors.New(InvalidParameters)
	}
	pub, err := x509.ReadPublicKeyFromHex(string(key))
	if err != nil {
		return false, err
	}
	var R, S big.Int
	pos := len(sign) / 2
	R.SetBytes(sign[:pos])
	S.SetBytes(sign[pos:])
	return sm2.Sm2Verify(pub, data, nil, &R, &S), nil
}

// SignWithBase64 签名（私钥），使用base64
func (s SM2) SignWithBase64(data []byte, key []byte) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New(InvalidParameters)
	}
	sign, err := s.Sign(data, key)
	if err != nil {
		return "", err
	}
	return s.b64.Encode(sign), nil
}

// VerifyWithBase64 校验（公钥），使用base64
func (s SM2) VerifyWithBase64(data []byte, key []byte, sign string) (bool, error) {
	if len(data) == 0 || len(key) == 0 {
		return false, errors.New(InvalidParameters)
	}
	signRaw, err := s.b64.Decode(sign)
	if err != nil {
		return false, err
	}
	return s.Verify(data, key, signRaw)
}

// GenerateKey 生成密钥对，返回私钥、公钥、错误
func (s SM2) GenerateKey() ([]byte, []byte, error) {
	priv, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
	if err != nil {
		return nil, nil, err
	}
	return []byte(x509.WritePrivateKeyToHex(priv)),
		[]byte(x509.WritePublicKeyToHex(priv.Public().(*sm2.PublicKey))),
		nil
}
