package test

import (
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/msgcode"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strings"
	"testing"
	"time"
)

var (
	r *redis.Redis
)

func initTest(t *testing.T) {
	r = redis.New("82.156.37.89:1443", func(r *redis.Redis) {
		r.Pass = "Keydb@1234567"
	})
	r.Del("hash_msgcode_code", "hash_msgcode_freeze")
}

// [正常系]内存版获取验证码并验证
func TestRedisSendAndCheckNormal(t *testing.T) {
	initTest(t)
	p := msgcode.Params{
		FreezeTime:     10,
		DurationCount:  2,
		DurationWindow: 5,
		ValidTime:      2,
		WrongTime:      2,
	}
	mc := msgcode.NewRedisMsgCode(p, r)
	code, want := mc.Gen("13520548443")
	assert.Equal(t, want, nil)
	assert.Equal(t, len(code) == 6, true)
	assert.Equal(t, isNumber(code), true)
	want = mc.Check("13520548443", code)
	assert.Equal(t, want, nil)
}

// [异常系]内存版获取验证码并验证失效验证码
func TestRedisSendAndCheckAbnormal(t *testing.T) {
	initTest(t)
	p := msgcode.Params{
		FreezeTime:     10,
		DurationCount:  2,
		DurationWindow: 5,
		ValidTime:      2,
		WrongTime:      2,
	}
	mc := msgcode.NewRedisMsgCode(p, r)
	code, want := mc.Gen("13520548443")
	assert.Equal(t, want, nil)
	assert.Equal(t, len(code) == 6, true)
	assert.Equal(t, isNumber(code), true)
	time.Sleep(2 * time.Second)
	want = mc.Check("13520548443", code)
	assert.Equal(t, want != nil, true)
	assert.Equal(t, want.Error() == "验证码失效", true)
}

// [正常系]内存版获取验证码，两次
func TestRedisDoubleTimeSendAndCheckNormal(t *testing.T) {
	initTest(t)
	p := msgcode.Params{
		FreezeTime:     10,
		DurationCount:  2,
		DurationWindow: 5,
		ValidTime:      2,
		WrongTime:      2,
	}
	mc := msgcode.NewRedisMsgCode(p, r)
	for i := 0; i < 2; i++ {
		code, want := mc.Gen("13520548443")
		assert.Equal(t, want, nil)
		assert.Equal(t, len(code) == 6, true)
		assert.Equal(t, isNumber(code), true)
		want = mc.Check("13520548443", code)
		assert.Equal(t, want, nil)
	}
}

// [异常系]内存版获取验证码，两次
func TestRedisDoubleTimeSendNormal(t *testing.T) {
	initTest(t)
	p := msgcode.Params{
		FreezeTime:     10,
		DurationCount:  2,
		DurationWindow: 2,
		ValidTime:      2,
		WrongTime:      2,
	}
	mc := msgcode.NewRedisMsgCode(p, r)
	code, want := mc.Gen("13520548443")
	assert.Equal(t, want, nil)
	assert.Equal(t, len(code) == 6, true)
	assert.Equal(t, isNumber(code), true)

	code, want = mc.Gen("13520548443")
	assert.Equal(t, want != nil, true)
	assert.Equal(t, strings.HasPrefix(want.Error(), "短信验证码发送过快，请于"), true)

	time.Sleep(time.Second * time.Duration(p.ValidTime))
	code, want = mc.Gen("13520548443")
	assert.Equal(t, want, nil)
	assert.Equal(t, len(code) == 6, true)
	assert.Equal(t, isNumber(code), true)
}

// [异常系]内存版冻结
func TestRedisFrozenNormal(t *testing.T) {
	initTest(t)
	p := msgcode.Params{
		FreezeTime:     10,
		DurationCount:  2,
		DurationWindow: 5,
		ValidTime:      1,
		WrongTime:      2,
	}
	mc := msgcode.NewRedisMsgCode(p, r)
	for i := 0; i < 2; i++ {
		code, want := mc.Gen("13520548443")
		assert.Equal(t, want, nil)
		assert.Equal(t, len(code) == 6, true)
		assert.Equal(t, isNumber(code), true)
		want = mc.Check("13520548443", code)
		assert.Equal(t, want, nil)
	}
	code, want := mc.Gen("13520548443")
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(code) == 0, true)
	assert.Equal(t, want.Error() == "短信验证码发送过快，请稍后重试", true)
}

// [正常系]内存版冻结解除
func TestRedisFrozenExpiredNormal(t *testing.T) {
	initTest(t)
	p := msgcode.Params{
		FreezeTime:     2,
		DurationCount:  2,
		DurationWindow: 5,
		ValidTime:      1,
		WrongTime:      2,
	}
	mc := msgcode.NewRedisMsgCode(p, r)
	for i := 0; i < 2; i++ {
		code, want := mc.Gen("13520548443")
		assert.Equal(t, want, nil)
		assert.Equal(t, len(code) == 6, true)
		assert.Equal(t, isNumber(code), true)
		want = mc.Check("13520548443", code)
		assert.Equal(t, want, nil)
	}
	code, want := mc.Gen("13520548443")
	assert.Equal(t, want != nil, true)
	assert.Equal(t, len(code) == 0, true)
	assert.Equal(t, want.Error() == "短信验证码发送过快，请稍后重试", true)
	time.Sleep(time.Second * time.Duration(p.FreezeTime+1))
	code, want = mc.Gen("13520548443")
	assert.Equal(t, want, nil)
	assert.Equal(t, len(code) == 6, true)
	assert.Equal(t, isNumber(code), true)
	want = mc.Check("13520548443", code)
	assert.Equal(t, want, nil)
}

// [异常系]内存版验证码验证错误
func TestRedisWrongCodeNormal(t *testing.T) {
	initTest(t)
	p := msgcode.Params{
		FreezeTime:     2,
		DurationCount:  2,
		DurationWindow: 5,
		ValidTime:      60,
		WrongTime:      2,
	}
	mc := msgcode.NewRedisMsgCode(p, r)
	code, want := mc.Gen("13520548443")
	assert.Equal(t, want, nil)
	assert.Equal(t, len(code) == 6, true)
	assert.Equal(t, isNumber(code), true)
	for i := 0; i < p.WrongTime; i++ {
		want = mc.Check("13520548443", "1")
		assert.Equal(t, want != nil, true)
		assert.Equal(t, want.Error() == "验证码错误", true)
	}
	want = mc.Check("13520548443", "1")
	assert.Equal(t, want != nil, true)
	assert.Equal(t, want.Error() == "验证码失效", true)
}
