package test

import (
	"fmt"
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

// [正常系]生成rsa秘钥
func TestRSAGenerateKeyNormal(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.RSA{}}
	publicPath := "./pub.pem"
	privatePath := "./priv.pem"
	want := c.Asymmetric.GenerateKey(publicPath, privatePath)
	assert.Equal(t, want, nil)
	_, err := os.Stat(publicPath)
	assert.Equal(t, err == nil || os.IsExist(err), true)
	_, err = os.Stat(privatePath)
	assert.Equal(t, err == nil || os.IsExist(err), true)
}

// [正常系]生成rsa秘钥
func TestRSAGeneratePublicKey(t *testing.T) {
	c := crypto.Crypto{Asymmetric: crypto.RSA{}}
	priv := `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`
	publicKey, err := c.Asymmetric.GeneratePublicKey([]byte(priv))
	fmt.Println(publicKey)
	assert.Equal(t, err, nil)
	assert.Equal(t, publicKey != "", true)
}
