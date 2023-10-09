package msgcode

import (
	"github.com/mailru/easyjson"
	"math/rand"
	"time"
)

// Code 验证码
type Code struct {
	Code     string    `json:"code"`      // 验证码
	Wrong    int       `json:"wrong"`     // 错误次数
	ExpireAt time.Time `json:"expire_at"` // 过期时间
}

// 初始化code
func initCode(valid int) Code {
	return Code{
		Code:     genCode(),
		ExpireAt: time.Now().Add(time.Second * time.Duration(valid)),
		Wrong:    0,
	}
}

// Freeze 冻结结构
type Freeze struct {
	ExpireAt time.Time `json:"expire_at"` // 冻结过期时间
	Window   time.Time `json:"window"`    // 冻结窗口结束时间
	Count    int       `json:"count"`     // 请求次数
}

// 初始化Freeze
func initFreeze(window int) Freeze {
	return Freeze{
		Count:  1,
		Window: time.Now().Add(time.Second * time.Duration(window)),
	}
}

// 随机码
var randomCodes = [...]byte{
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
}

// 生成验证码
func genCode() string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var code = make([]byte, 6)
	for j := 0; j < 6; j++ {
		index := r.Int() % len(randomCodes)
		code[j] = randomCodes[index]
	}
	return string(code)
}

// code转换json
func code2Json(code Code) (string, error) {
	jsonByte, err := easyjson.Marshal(code)
	if err != nil {
		return "", err
	}
	return string(jsonByte), nil
}

// json转换code
func json2Code(json []byte) (Code, error) {
	var code Code
	err := easyjson.Unmarshal(json, &code)
	return code, err
}

// freeze转换json
func freeze2Json(freeze Freeze) (string, error) {
	jsonByte, err := easyjson.Marshal(freeze)
	if err != nil {
		return "", err
	}
	return string(jsonByte), nil
}

// json转换freeze
func json2Freeze(json []byte) (Freeze, error) {
	var freeze Freeze
	err := easyjson.Unmarshal(json, &freeze)
	return freeze, err
}
