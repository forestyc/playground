package test

import (
	"github.com/forestyc/playground/go-zero/epidemic/pkg/crypto"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var (
	private = `89369080dfc1617733a569a6edf3ea6cf547c0a62d9d5a82bc7d2c4a5a17f4c6`
	public  = `040ff1101ae15e5b7e5b863a2785e0be00a677dbf08d28978c3c8656c40fc3f62c78e0f420f3f9eb3d5e5e9a1095de39ad6480910e7573c7fc3cf80ed33eb70790`
)

// [正常系]AES测试
func TestAesNormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := "abcd哈哈a123"
	key := "310c4dbc712ce3bc"
	// 加密
	crypto, want := adapter.Symmetric["aes"].Encrypt([]byte(data), []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(crypto) > 0, true)
	// 解密
	raw, want := adapter.Symmetric["aes"].Decrypt(crypto, []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]AES测试
func TestAesAbnormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := ""
	key := ""
	// 加密
	crypto, want := adapter.Symmetric["aes"].Encrypt([]byte(data), []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(crypto) == 0, true)
	// 解密
	raw, want := adapter.Symmetric["aes"].Decrypt(crypto, []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]AES+base64测试
func TestAesWithBase64Normal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := "abcd哈哈a123"
	key := "310c4dbc712ce3bc"
	// 加密
	base64Code, want := adapter.Symmetric["aes"].EncryptWithBase64([]byte(data), []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(base64Code) > 0, true)
	// 解密
	raw, want := adapter.Symmetric["aes"].DecryptWithBase64(base64Code, []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]AES+base64测试
func TestAesWithBase64Abnormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := ""
	key := ""
	// 加密
	base64Code, want := adapter.Symmetric["aes"].EncryptWithBase64([]byte(data), []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(base64Code) == 0, true)
	// 解密
	raw, want := adapter.Symmetric["aes"].DecryptWithBase64(base64Code, []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]SM4测试
func TestSM4Normal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := "abcd哈哈a123"
	key := "310c4dbc712ce3bc"
	crypto, want := adapter.Symmetric["sm4"].Encrypt([]byte(data), []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(crypto) > 0, true)
	raw, want := adapter.Symmetric["sm4"].Decrypt(crypto, []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]SM4测试
func TestSM4Abnormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := ""
	key := ""
	crypto, want := adapter.Symmetric["sm4"].Encrypt([]byte(data), []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(crypto) == 0, true)
	raw, want := adapter.Symmetric["sm4"].Decrypt(crypto, []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]SM4+base64测试
func TestSM4WithBase64Normal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := "abcd哈哈a123"
	key := "310c4dbc712ce3bc"
	base64Code, want := adapter.Symmetric["sm4"].EncryptWithBase64([]byte(data), []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(base64Code) > 0, true)
	raw, want := adapter.Symmetric["sm4"].DecryptWithBase64(base64Code, []byte(key))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]SM4+base64测试
func TestSM4WithBase64Abnormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := ""
	key := ""
	base64Code, want := adapter.Symmetric["sm4"].EncryptWithBase64([]byte(data), []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(base64Code) == 0, true)
	raw, want := adapter.Symmetric["sm4"].DecryptWithBase64(base64Code, []byte(key))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]SM3测试
func TestSM3(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := "abcd哈哈a123"
	crypto, want := adapter.Digest["sm3"].Sum([]byte(data))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(crypto) > 0, true)
}

// [正常系]SM2加密测试
func TestSM2CryptoNormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := "abcd哈哈a123"
	crypto, want := adapter.Asymmetric["sm2"].Encrypt([]byte(data), []byte(public))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(crypto) > 0, true)
	raw, want := adapter.Asymmetric["sm2"].Decrypt(crypto, []byte(private))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]SM2加密测试
func TestSM2CryptoAbnormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := ""
	private := ``
	public := ``
	crypto, want := adapter.Asymmetric["sm2"].Encrypt([]byte(data), []byte(public))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(crypto) == 0, true)
	raw, want := adapter.Asymmetric["sm2"].Decrypt(crypto, []byte(private))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]SM2+base64加密测试
func TestSM2WithBase64CryptoNormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := "abcd哈哈a123"
	base64Code, want := adapter.Asymmetric["sm2"].EncryptWithBase64([]byte(data), []byte(public))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(base64Code) > 0, true)
	raw, want := adapter.Asymmetric["sm2"].DecryptWithBase64(base64Code, []byte(private))
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]SM2+base64加密测试
func TestSM2WithBase64CryptoAbnormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := ""
	private := ``
	public := ``
	base64Code, want := adapter.Asymmetric["sm2"].EncryptWithBase64([]byte(data), []byte(public))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(base64Code) == 0, true)
	raw, want := adapter.Asymmetric["sm2"].DecryptWithBase64(base64Code, []byte(private))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(raw) == 0, true)
}

// [正常系]SM2签名测试
func TestSM2SignNormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := "abcd哈哈a123"
	sign, want := adapter.Asymmetric["sm2"].Sign([]byte(data), []byte(private))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(sign) > 0, true)
	result, want := adapter.Asymmetric["sm2"].Verify([]byte(data), []byte(public), sign)
	assert.Equal(t, want, nil)
	assert.Equal(t, result, true)
}

// [异常系]SM2签名测试
func TestSM2SignAbnormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := ""
	private := ``
	public := ``
	sign, want := adapter.Asymmetric["sm2"].Sign([]byte(data), []byte(private))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(sign) == 0, true)
	result, want := adapter.Asymmetric["sm2"].Verify([]byte(data), []byte(public), sign)
	assert.Equal(t, want != nil, true)
	assert.Equal(t, result, false)
}

// [正常系]SM2+base64签名测试
func TestSM2SignWithBase64Normal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := "abcd哈哈a123"
	signBase64, want := adapter.Asymmetric["sm2"].Sign([]byte(data), []byte(private))
	assert.Equal(t, want, nil)
	assert.Equal(t, len(signBase64) > 0, true)
	result, want := adapter.Asymmetric["sm2"].Verify([]byte(data), []byte(public), signBase64)
	assert.Equal(t, want, nil)
	assert.Equal(t, result, true)
}

// [异常系]SM2+base64签名测试
func TestSM2SignWithBase64Abnormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	data := ""
	private := ``
	public := ``
	signBase64, want := adapter.Asymmetric["sm2"].Sign([]byte(data), []byte(private))
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(signBase64) == 0, true)
	result, want := adapter.Asymmetric["sm2"].Verify([]byte(data), []byte(public), signBase64)
	assert.Equal(t, want != nil, true)
	assert.Equal(t, result, false)
}

// [正常系]生成sm2秘钥
func TestSM2GenerateKeyNormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	publicPath := "./pub.pem"
	privatePath := "./priv.pem"
	want := adapter.Asymmetric["sm2"].GenerateKey(publicPath, privatePath)
	assert.Equal(t, want, nil)
	_, err := os.Stat(publicPath)
	assert.Equal(t, err == nil || os.IsExist(err), true)
	_, err = os.Stat(privatePath)
	assert.Equal(t, err == nil || os.IsExist(err), true)
	pubkeyPem, err := ioutil.ReadFile(publicPath)
	assert.Equal(t, err, nil)
	privkeyPem, err := ioutil.ReadFile(privatePath)
	assert.Equal(t, err, nil)
	data := "abcd哈哈哈1234"
	crypto, want := adapter.Asymmetric["sm2"].Encrypt([]byte(data), pubkeyPem)
	assert.Equal(t, want, nil)
	assert.Equal(t, len(crypto) > 0, true)
	raw, want := adapter.Asymmetric["sm2"].Decrypt(crypto, privkeyPem)
	assert.Equal(t, want, nil)
	assert.Equal(t, string(raw) == data, true)
}

// [异常系]生成sm2秘钥
func TestSM2GenerateKeyAbnormal(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	publicPath := ""
	privatePath := ""
	want := adapter.Asymmetric["sm2"].GenerateKey(publicPath, privatePath)
	assert.Equal(t, want != nil, true)
}

// [正常系]生成sm2秘钥
func TestGeneratePublicKey(t *testing.T) {
	adapter := crypto.InitCryptoAdapter()
	publicKey, err := adapter.Asymmetric[crypto.CryptoSM2].GeneratePublicKey([]byte("89369080dfc1617733a569a6edf3ea6cf547c0a62d9d5a82bc7d2c4a5a17f4c6"))
	assert.Equal(t, err, nil)
	assert.Equal(t, publicKey != "", true)
}
