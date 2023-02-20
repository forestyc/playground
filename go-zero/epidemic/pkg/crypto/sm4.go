package crypto

import (
	"errors"
	"github.com/tjfoc/gmsm/sm4"
)

type SM4 struct {
}

// Encrypt 加密
func (s SM4) Encrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	return sm4.Sm4Ecb(key, data, true)
}

// Decrypt 解密
func (s SM4) Decrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	return sm4.Sm4Ecb(key, data, false)
}

// EncryptWithBase64 加密，使用base64
func (s SM4) EncryptWithBase64(data []byte, key []byte) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New("invalid params")
	}
	crypto, err := s.Encrypt(data, key)
	if err != nil {
		return "", err
	}
	return Base64Encode(crypto), nil
}

// DecryptWithBase64 解密，使用base64
func (s SM4) DecryptWithBase64(data string, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("invalid params")
	}
	crypto, err := Base64Decode(data)
	if err != nil {
		return nil, err
	}
	return s.Decrypt(crypto, key)
}
