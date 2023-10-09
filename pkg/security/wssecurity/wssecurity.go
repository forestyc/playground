package wssecurity

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/forestyc/playground/pkg/crypto"
	uuid "github.com/satori/go.uuid"
)

type XWSSE struct {
	PasswordDigest string
	AppKey         string
	Nonce          string
	Created        string
}

// Check x-wsse 校验
// Authorization=`WSSE profile="UsernameToken"`
// X-WSSE: UsernameToken Username="bob", PasswordDigest="quR/EWLAV4xLf9Zqyw4pDmfV9OY=", Nonce="d36e316282959a9ed4c89851497a717f", Created="2006-01-02T15:04:05Z"
// PasswordDigest=base64(sm3(Nonce + Created + AppSecret))
// Nonce 随机数
// Created 时间， 2006-01-02T15:04:05Z
func (x *XWSSE) Check(authorization, xwsse string, getAppSecret ApiCheck) (err error) {
	// 校验Authorization，固定值
	if authorization != `WSSE profile="UsernameToken"` {
		return errors.New("invalid Authorization")
	}
	// 校验X-WSSE
	if err = x.UnMarshal(xwsse); err != nil {
		return
	}
	// 获取app-secret
	var appSecret string
	if appSecret, err = getAppSecret(x.AppKey); err != nil {
		return
	}
	// 校验摘要
	if err = x.CheckSum(appSecret); err != nil {
		return
	}
	return
}

// Marshal 序列化
func (x *XWSSE) Marshal(appSecret string) (xwsse string, err error) {
	format := `UsernameToken Username="` + x.AppKey + `", PasswordDigest="%s", Nonce="%s", Created="%s"`
	if len(x.Nonce) == 0 {
		x.Nonce = uuid.NewV4().String()
	}
	if len(x.Created) == 0 {
		x.Created = time.Now().Format("2006-01-02T15:04:05Z")
	}
	var sm3 crypto.SM3
	var passwordDigest string
	if passwordDigest, err = sm3.Sum([]byte(x.Nonce + x.Created + appSecret)); err != nil {
		return
	}
	xwsse = fmt.Sprintf(format, passwordDigest, x.Nonce, x.Created)
	return
}

// UnMarshal 反序列化
func (x *XWSSE) UnMarshal(xwsse string) (err error) {
	resultMap := make(map[string]string)
	xwsse = strings.Replace(xwsse, " ", "", -1)
	xwsse = strings.Replace(xwsse, "UsernameToken", "", -1)
	xwsse = strings.Replace(xwsse, `"`, "", -1)
	paras := strings.Split(xwsse, ",")
	for i := 0; i < len(paras); i++ {
		keyValue := strings.Split(paras[i], "=")
		if len(keyValue) == 2 {
			resultMap[keyValue[0]] = keyValue[1]
		} else if len(keyValue) > 2 {
			// base64后面有=的情况
			resultMap[keyValue[0]] = keyValue[1]
			for i := 0; i < len(keyValue)-2; i++ {
				resultMap[keyValue[0]] += "="
			}
		}
	}
	x.Nonce = resultMap["Nonce"]
	x.Created = resultMap["Created"]
	x.AppKey = resultMap["Username"]
	x.PasswordDigest = resultMap["PasswordDigest"]
	return
}

// CheckSum 校验PasswordDigest
func (x *XWSSE) CheckSum(appSecret string) (err error) {
	var createdTime time.Time
	createdTime, err = time.ParseInLocation("2006-01-02T15:04:05Z", x.Created, time.Local)
	if err != nil {
		return
	}
	now := time.Now()
	// created比当前时间靠后或者是30秒之前的请求，直接拒绝
	if createdTime.After(now) || now.Sub(createdTime) > time.Second*RequestIn {
		return errors.New("invalid request")
	}
	var want string
	want, err = x.Sum(appSecret)
	if err != nil {
		return
	}
	if want != x.PasswordDigest {
		return errors.New("invalid passwordDigest")
	}
	return err
}

// Sum 生成PasswordDigest
// PasswordDigest=sm3(Nonce + Created + AppSecret)
func (x *XWSSE) Sum(appSecret string) (string, error) {
	sm3 := crypto.SM3{}
	return sm3.Sum([]byte(x.Nonce + x.Created + appSecret))
}
