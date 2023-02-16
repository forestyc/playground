package test

import (
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/sms"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 配置文件读取测试
func TestHuaweiCloudSendMsg(t *testing.T) {
	// 初始化加解密
	huawei := sms.NewHuaweiCloud(sms.Params{
		AppSecret:  "6nu843Bwc1tlGQraX04W25h6i5aW",
		AppKey:     "Zh4AZfNW85wmY75Estb5Km8HGcnn",
		Host:       "https://api.rtc.huaweicloud.com:10443/sms/batchSendSms/v1",
		TemplateId: "55ab415b225b40fd87362611d86b1acb",
		Sign:       "csms18083014",
	})
	want := huawei.SendMsg([]string{"13520548443"}, "123456")
	assert.Equal(t, want, nil)
}
