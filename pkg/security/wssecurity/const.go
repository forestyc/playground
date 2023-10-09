package wssecurity

type ApiCheck func(appKey string) (appSecret string, err error)

const (
	RequestIn = 30 // RequestIn秒内的请求
)
