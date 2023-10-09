package sms

// Params 初始化参数
type Params struct {
	AppKey     string
	AppSecret  string
	Host       string // 短信访问域名
	TemplateId string
	Sign       string
}

// Sms 云短信发送接口
type Sms interface {
	SendMsg(to []string, Msg string) error // 发送短信
}
