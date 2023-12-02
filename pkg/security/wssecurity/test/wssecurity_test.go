package test

import (
	"testing"
	"time"

	"github.com/forestyc/playground/pkg/crypto"
	"github.com/forestyc/playground/pkg/security/wssecurity"
	"github.com/forestyc/playground/pkg/utils"
	"github.com/go-playground/assert/v2"
)

// [正常系]生成appKey
func TestGenAppKey(t *testing.T) {
	appKey := utils.GenAppKey()
	assert.Equal(t, len(appKey) > 0, true)
	t.Logf("appKey=[%s]\n", appKey)
}

// [正常系]生成appSecret
func TestGenAppSecret(t *testing.T) {
	appSecret, err := utils.GenAppSecret("28599f31d80a4cfab02dd6c33214f028", "test") //967acbb541daf9ce8319d3c23406e4ec7383c9da70398b6bcefe4bbb80405509
	assert.Equal(t, err, nil)
	assert.Equal(t, len(appSecret) > 0, true)
	t.Logf("appSecret=[%s]\n", appSecret)
}

// [正常系]生成appSecret
func TestCheckNormal(t *testing.T) {
	var err error
	var xwsseStr, digest, appSecret string
	xwsseResult := &wssecurity.XWSSE{}
	var sm3 crypto.SM3
	appSecret = "967acbb541daf9ce8319d3c23406e4ec7383c9da70398b6bcefe4bbb80405509"
	xwsse := wssecurity.XWSSE{
		AppKey:  "28599f31d80a4cfab02dd6c33214f028",
		Nonce:   utils.GenUuid(),
		Created: time.Now().Format("2006-01-02T15:04:05Z"),
	}
	digest, err = sm3.Sum([]byte(xwsse.Nonce + xwsse.Created + appSecret))
	assert.Equal(t, err, nil)
	xwsseStr, err = xwsse.Marshal(appSecret)
	assert.Equal(t, err, nil)
	assert.Equal(t, xwsseStr == "", false)
	t.Logf("X-WSSE=[%s]\n", xwsseStr)
	err = xwsseResult.Check(`WSSE profile="UsernameToken"`, xwsseStr, getAppSecret)
	assert.Equal(t, err, nil)
	assert.Equal(t, xwsseResult.AppKey == xwsse.AppKey, true)
	assert.Equal(t, xwsseResult.Nonce == xwsse.Nonce, true)
	assert.Equal(t, xwsseResult.Created == xwsse.Created, true)
	assert.Equal(t, xwsseResult.PasswordDigest == digest, true)
}

// [异常系]生成appSecret
func TestCheckAbnormal(t *testing.T) {
	xwsse := "asdfasdfasdf"
	xwsseResult := &wssecurity.XWSSE{}
	err := xwsseResult.Check(`WSSE profile="UsernameToken"`, xwsse, getAppSecret)
	assert.Equal(t, err != nil, true)
}

// 获取appSecret
func getAppSecret(appKey string) (appSecret string, err error) {
	return "967acbb541daf9ce8319d3c23406e4ec7383c9da70398b6bcefe4bbb80405509", nil
}
