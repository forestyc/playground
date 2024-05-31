package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"github.com/forestyc/playground/pkg/encoding/base64"
	"github.com/pkg/errors"
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

//16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法

type AES struct {
	b64 base64.Base64
}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// Encrypt 加密
func (a AES) Encrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New(InvalidParameters)
	}
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypto := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypto, encryptBytes)
	return crypto, nil
}

// Decrypt 解密
func (a AES) Decrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New(InvalidParameters)
	}
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypto := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypto, data)
	//去除填充
	crypto, err = pkcs7UnPadding(crypto)
	if err != nil {
		return nil, err
	}
	return crypto, nil
}

// EncryptWithBase64 Aes加密 后 base64
func (a AES) EncryptWithBase64(data []byte, key []byte) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New(InvalidParameters)
	}
	res, err := a.Encrypt(data, key)
	if err != nil {
		return "", err
	}
	return a.b64.Encode(res), nil
}

// DecryptWithBase64 base64解码后 Aes 解密
func (a AES) DecryptWithBase64(data string, key []byte) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New(InvalidParameters)
	}
	dataByte, err := a.b64.Decode(data)
	if err != nil {
		return nil, err
	}
	return a.Decrypt(dataByte, key)
}
