package test

import (
	"github.com/forestyc/playground/pkg/crypto"
	"github.com/go-playground/assert/v2"
	"os"
	"testing"
)

var (
	private = `89369080dfc1617733a569a6edf3ea6cf547c0a62d9d5a82bc7d2c4a5a17f4c6`
	public  = `040ff1101ae15e5b7e5b863a2785e0be00a677dbf08d28978c3c8656c40fc3f62c78e0f420f3f9eb3d5e5e9a1095de39ad6480910e7573c7fc3cf80ed33eb70790`
)

// [正常系]AES测试
func TestAesNormal(t *testing.T) {
	c := crypto.Crypto{Symmetric: crypto.AES{}}
	data := "abcd哈哈a123"
	key := "310c4dbc712ce3bc"
	// 加密
	crypto, want := c.Symmetric.Encrypt([]byte(data), []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(crypto) > 0, true)
	// 解密
	raw, want := c.Symmetric.Decrypt(crypto, []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]AES测试
func TestAesAbnormal(t *testing.T) {
	c := crypto.Crypto{Symmetric: crypto.AES{}}
	data := ""
	key := ""
	// 加密
	crypto, want := c.Symmetric.Encrypt([]byte(data), []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(crypto) == 0, true)
	// 解密
	raw, want := c.Symmetric.Decrypt(crypto, []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]AES+base64测试
func TestAesWithBase64Normal(t *testing.T) {
	c := crypto.Crypto{Symmetric: crypto.AES{}}
	data := "abcd哈哈a123"
	key := "310c4dbc712ce3bc"
	// 加密
	base64Code, want := c.Symmetric.EncryptWithBase64([]byte(data), []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(base64Code) > 0, true)
	// 解密
	raw, want := c.Symmetric.DecryptWithBase64(base64Code, []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]AES+base64测试
func TestAesWithBase64Abnormal(t *testing.T) {
	c := crypto.Crypto{Symmetric: crypto.AES{}}
	data := ""
	key := ""
	// 加密
	base64Code, want := c.Symmetric.EncryptWithBase64([]byte(data), []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(base64Code) == 0, true)
	// 解密
	raw, want := c.Symmetric.DecryptWithBase64(base64Code, []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]SM4测试
func TestSM4Normal(t *testing.T) {
	c := crypto.Crypto{Symmetric: crypto.SM4{}}
	data := "abcd哈哈a123"
	key := "310c4dbc712ce3bc"
	crypto, want := c.Symmetric.Encrypt([]byte(data), []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(crypto) > 0, true)
	raw, want := c.Symmetric.Decrypt(crypto, []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]SM4测试
func TestSM4Abnormal(t *testing.T) {
	c := crypto.Crypto{Symmetric: crypto.SM4{}}
	data := ""
	key := ""
	crypto, want := c.Symmetric.Encrypt([]byte(data), []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(crypto) == 0, true)
	raw, want := c.Symmetric.Decrypt(crypto, []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]SM4+base64测试
func TestSM4WithBase64Normal(t *testing.T) {
	c := crypto.Crypto{Symmetric: crypto.SM4{}}
	data := "abcd哈哈a123"
	key := "310c4dbc712ce3bc"
	base64Code, want := c.Symmetric.EncryptWithBase64([]byte(data), []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(base64Code) > 0, true)
	raw, want := c.Symmetric.DecryptWithBase64(base64Code, []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]SM4+base64测试
func TestSM4WithBase64Abnormal(t *testing.T) {
	c := crypto.Crypto{Symmetric: crypto.SM4{}}
	data := ""
	key := ""
	base64Code, want := c.Symmetric.EncryptWithBase64([]byte(data), []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(base64Code) == 0, true)
	raw, want := c.Symmetric.DecryptWithBase64(base64Code, []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]SM3测试
func TestSM3(t *testing.T) {
	c := crypto.Crypto{Digest: crypto.SM3{}}
	data := "abcd哈哈a123"
	crypto, want := c.Digest.Sum([]byte(data))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(crypto) > 0, true)
}

// [正常系]SM2加密测试
func TestSM2CryptoNormal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	data := "abcd哈哈a123"
	crypto, want := c.Asymmetric.Encrypt([]byte(data), []byte(public))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(crypto) > 0, true)
	raw, want := c.Asymmetric.Decrypt(crypto, []byte(private))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]SM2加密测试
func TestSM2CryptoAbnormal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	data := ""
	private := ``
	public := ``
	crypto, want := c.Asymmetric.Encrypt([]byte(data), []byte(public))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(crypto) == 0, true)
	raw, want := c.Asymmetric.Decrypt(crypto, []byte(private))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]SM2+base64加密测试
func TestSM2WithBase64CryptoNormal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	data := "abcd哈哈a123"
	base64Code, want := c.Asymmetric.EncryptWithBase64([]byte(data), []byte(public))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(base64Code) > 0, true)
	raw, want := c.Asymmetric.DecryptWithBase64(base64Code, []byte(private))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]SM2+base64加密测试
func TestSM2WithBase64CryptoAbnormal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	data := ""
	private := ``
	public := ``
	base64Code, want := c.Asymmetric.EncryptWithBase64([]byte(data), []byte(public))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(base64Code) == 0, true)
	raw, want := c.Asymmetric.DecryptWithBase64(base64Code, []byte(private))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]SM2签名测试
func TestSM2SignNormal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	data := "abcd哈哈a123"
	sign, want := c.Asymmetric.Sign([]byte(data), []byte(private))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(sign) > 0, true)
	result, want := c.Asymmetric.Verify([]byte(data), []byte(public), sign)
	assert.Equal(t, want, nil)
	assert.Equal(t, result, true)
}

// [异常系]SM2签名测试
func TestSM2SignAbnormal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	data := ""
	private := ``
	public := ``
	sign, want := c.Asymmetric.Sign([]byte(data), []byte(private))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(sign) == 0, true)
	result, want := c.Asymmetric.Verify([]byte(data), []byte(public), sign)
	assert.Equal(t, want != nil, true)
	assert.Equal(t, result, false)
}

// [正常系]SM2+base64签名测试
func TestSM2SignWithBase64Normal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	data := "abcd哈哈a123"
	signBase64, want := c.Asymmetric.SignWithBase64([]byte(data), []byte(private))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(signBase64) > 0, true)
	result, want := c.Asymmetric.VerifyWithBase64([]byte(data), []byte(public), signBase64)
	assert.Equal(t, want, nil)
	assert.Equal(t, result, true)
}

// [异常系]SM2+base64签名测试
func TestSM2SignWithBase64Abnormal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	data := ""
	private := ``
	public := ``
	signBase64, want := c.Asymmetric.SignWithBase64([]byte(data), []byte(private))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(signBase64) == 0, true)
	result, want := c.Asymmetric.VerifyWithBase64([]byte(data), []byte(public), signBase64)
	assert.Equal(t, want != nil, true)
	assert.Equal(t, result, false)
}

// [正常系]生成sm2秘钥
func TestSM2GenerateKeyNormal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	publicPath := "./pub.pem"
	privatePath := "./priv.pem"
	want := c.Asymmetric.GenerateKey(publicPath, privatePath)
	assert.Equal(t, want, nil)
	_, err := os.Stat(publicPath)
	assert.Equal(t, err == nil || os.IsExist(err), true)
	_, err = os.Stat(privatePath)
	assert.Equal(t, err == nil || os.IsExist(err), true)
	pubkeyPem, err := os.ReadFile(publicPath)
	assert.Equal(t, err, nil)
	privkeyPem, err := os.ReadFile(privatePath)
	assert.Equal(t, err, nil)
	data := "abcd哈哈哈1234"
	crypto, want := c.Asymmetric.Encrypt([]byte(data), pubkeyPem)
	assert.Equal(t, want, nil)
	assert.Equal(t, len(crypto) > 0, true)
	raw, want := c.Asymmetric.Decrypt(crypto, privkeyPem)
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]生成sm2秘钥
func TestSM2GenerateKeyAbnormal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	publicPath := ""
	privatePath := ""
	want := c.Asymmetric.GenerateKey(publicPath, privatePath)
	assert.Equal(t, want != nil, true)
}

// [正常系]生成sm2秘钥
func TestGeneratePublicKey(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.SM2{}}
	publicKey, err := c.Asymmetric.GeneratePublicKey([]byte("89369080dfc1617733a569a6edf3ea6cf547c0a62d9d5a82bc7d2c4a5a17f4c6"))
	assert.Equal(t, err, nil)
	assert.Equal(t, publicKey != "", true)
}
