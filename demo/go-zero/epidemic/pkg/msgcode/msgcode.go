package msgcode

// Params 验证码参数
type Params struct {
	FreezeTime     int // 验证码冻结时间（秒）
	DurationCount  int // 验证码单位时间窗口内最大发送次数（次）
	DurationWindow int // 验证码单位时间窗口（秒）
	ValidTime      int // 验证码失效时间（秒）
	WrongTime      int // 验证码容错次数（次）
}

// Generator 验证码生成器
type Generator interface {
	Gen(account string) (string, error) // 生成验证码
	Check(account, code string) error   // 校验验证码
}
