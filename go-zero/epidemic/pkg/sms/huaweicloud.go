package sms

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HuaweiCloud struct {
	params Params
}

// Result 华为云应答
type Result struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Result      []SmsID `json:"result"`
}

// SmsID 华为云短信ID列表
type SmsID struct {
	SmsMsgId   string `json:"smsMsgId"`
	From       string `json:"from"`
	OriginTo   string `json:"originTo"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}

func NewHuaweiCloud(params Params) *HuaweiCloud {
	return &HuaweiCloud{
		params: params,
	}
}

// SendMsg 发送短信
func (h *HuaweiCloud) SendMsg(to []string, templateParas string) error {
	values := url.Values{}
	values.Set("from", h.params.Sign)
	values.Set("to", h.formatTo(to))
	values.Set("templateId", h.params.TemplateId)
	if templateParas != "" {
		values.Set("templateParas", fmt.Sprintf(`["%s"]`, templateParas))
	}
	request, err := http.NewRequest("POST", h.params.Host, ioutil.NopCloser(strings.NewReader(values.Encode())))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", `WSSE realm="SDP",profile="UsernameToken",type="Appkey"`)
	request.Header.Set("X-WSSE", h.genXWSSE())

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	var result Result
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.Code == "000000" {
		return nil
	} else {
		return errors.New(string(body))
	}
}

// genXWSSE 生成xwsse
func (h *HuaweiCloud) genXWSSE() string {
	xwsse := `UsernameToken Username="` + h.params.AppKey + `", PasswordDigest="%s", Nonce="%s", Created="%s"`
	nonce := h.genNonce()
	created := time.Now().Format("2006-01-02T15:04:05Z")
	password := h.params.AppSecret
	tmp := sha256.Sum256([]byte(nonce + created + password))
	passwordDigest := tmp[:]
	return fmt.Sprintf(xwsse, base64.StdEncoding.EncodeToString(passwordDigest), nonce, created)
}

// genNonce 生成nonce
func (h *HuaweiCloud) genNonce() string {
	return uuid.NewString()
}

// formatTo 格式化多个手机号
func (h *HuaweiCloud) formatTo(to []string) string {
	return strings.Join(to, ",")
}
