package crypto

import (
	"github.com/forestyc/playground/pkg/encoding/base64"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm4"
)

type SM4 struct {
	b64 base64.Base64
}

// Encrypt 加密
func (s SM4) Encrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New(InvalidParameters)
	}
	return sm4.Sm4Ecb(key, data, true)
}

// Decrypt 解密
func (s SM4) Decrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New(InvalidParameters)
	}
	return sm4.Sm4Ecb(key, data, false)
}

// EncryptWithBase64 加密，使用base64
func (s SM4) EncryptWithBase64(data []byte, key []byte) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New(InvalidParameters)
	}
	crypto, err := s.Encrypt(data, key)
	if err != nil {
		return "", err
	}
	return s.b64.Encode(crypto), nil
}

// DecryptWithBase64 解密，使用base64
func (s SM4) DecryptWithBase64(data string, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New(InvalidParameters)
	}
	crypto, err := s.b64.Decode(data)
	if err != nil {
		return nil, err
	}
	return s.Decrypt(crypto, key)
}
