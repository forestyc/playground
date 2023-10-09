package msgcode

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"
)

const (
	CodeHashKey   = "hash_msgcode_code"   // field:account, value:{"code":"123457","wrong":0,"expire_at":"time,RFC3339"}
	FreezeHashKey = "hash_msgcode_freeze" // field:account, value:{"count":2,"window":"time,RFC3339","expire_at":"time,RFC3339"}
)

type Redis struct {
	params Params
	redis  *redis.Redis
}

func NewRedisMsgCode(params Params, redis *redis.Redis) *Redis {
	return &Redis{
		params: params,
		redis:  redis,
	}
}

// Gen 生成验证码
func (r *Redis) Gen(account string) (string, error) {
	// 校验是否冻结
	frozen, err := r.frozen(account)
	if err != nil {
		return "", err
	}
	if frozen {
		return "", errors.New("验证码发送过快，请稍后重试")
	}
	now := time.Now()
	// 生成验证码
	var code Code
	ok, err := r.hgetCode(account, &code)
	if err != nil {
		return "", err
	}
	if !ok || code.ExpireAt.Before(now) { // 不存在或已过期
		code = initCode(r.params.ValidTime)
		err = r.hsetCode(code, account)
		if err != nil {
			return "", err
		}
		return code.Code, nil
	} else {
		timeLeft := code.ExpireAt.Sub(time.Now()).Seconds()
		left, _ := strconv.Atoi(fmt.Sprintf("%1.0f", timeLeft))
		return "", errors.New("验证码发送过快，请于" + strconv.Itoa(left) + "秒后重试")
	}
}

// Check 校验验证码
func (r *Redis) Check(account, code string) error {
	if account == "" || code == "" {
		return errors.New("验证码错误")
	}
	var codeCached Code
	ok, err := r.hgetCode(account, &codeCached)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("验证码失效")
	} else {
		if codeCached.Wrong >= r.params.WrongTime {
			// 超出最大错误数，立刻清掉验证码
			if _, err = r.redis.Hdel(CodeHashKey, account); err != nil {
				return err
			}
			return errors.New("验证码失效")
		}
		now := time.Now()
		if codeCached.ExpireAt.Before(now) {
			return errors.New("验证码失效")
		}
		if codeCached.Code != code {
			codeCached.Wrong += 1
			if err := r.hsetCode(codeCached, account); err != nil {
				return err
			}
			return errors.New("验证码错误")
		}
		// 验证通过，删除验证码

		if _, err = r.redis.Hdel(CodeHashKey, account); err != nil {
			return err
		}
		return nil
	}
}

// 冻结
func (r *Redis) frozen(account string) (bool, error) {
	var freeze Freeze
	ok, err := r.hgetFreeze(account, &freeze)
	if err != nil {
		return false, err
	}
	if ok {
		now := time.Now()
		if freeze.ExpireAt.Equal(time.Time{}) { // 未设置冻结
			if freeze.Count >= r.params.DurationCount { // 窗口内请求超限
				if freeze.Window.Before(time.Now()) { // 窗口已过期，不冻结
					freeze = initFreeze(r.params.DurationWindow)
					if err = r.hsetFreeze(freeze, account); err != nil {
						return false, err
					}
					return false, nil
				} else { // 未过期，需要冻结
					freeze.ExpireAt = time.Now().Add(time.Second * time.Duration(r.params.FreezeTime))
					if err = r.hsetFreeze(freeze, account); err != nil {
						return false, err
					}
					return true, nil
				}
			} else { // 窗口内请求未超限
				freeze.Count += 1
				if err = r.hsetFreeze(freeze, account); err != nil {
					return false, err
				}
			}
		} else { // 已设置冻结
			if freeze.ExpireAt.Before(now) { // 冻结已解除
				freeze = initFreeze(r.params.DurationWindow)
				if err = r.hsetFreeze(freeze, account); err != nil {
					return false, err
				}
				return false, nil
			} else { // 冻结未接触，仍请求，重新计算冻结时间
				freeze.ExpireAt = time.Now().Add(time.Second * time.Duration(r.params.FreezeTime))
				if err = r.hsetFreeze(freeze, account); err != nil {
					return false, err
				}
				return true, nil
			}
		}
	} else {
		freeze = initFreeze(r.params.DurationWindow)
		if err = r.hsetFreeze(freeze, account); err != nil {
			return false, err
		}
	}
	return false, err
}

// redis操作hset Code
func (r *Redis) hsetCode(code Code, account string) error {
	json, err := code2Json(code)
	if err != nil {
		return err
	}
	if err = r.redis.Hset(CodeHashKey, account, json); err != nil {
		return err
	}
	return nil
}

// redis操作hget Code
func (r *Redis) hgetCode(account string, code *Code) (bool, error) {
	codeStr, err := r.redis.Hget(CodeHashKey, account)
	if err != nil && err != redis.Nil {
		return false, err
	} else if err == redis.Nil {
		return false, nil
	}
	if len(codeStr) != 0 {
		*code, err = json2Code([]byte(codeStr))
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// redis操作hset Freeze
func (r *Redis) hsetFreeze(freeze Freeze, account string) error {
	json, err := freeze2Json(freeze)
	if err != nil {
		return err
	}
	if err = r.redis.Hset(FreezeHashKey, account, json); err != nil {
		return err
	}
	return nil
}

// redis操作hget Freeze
func (r *Redis) hgetFreeze(account string, freeze *Freeze) (bool, error) {
	freezeStr, err := r.redis.Hget(FreezeHashKey, account)
	if err != nil && err != redis.Nil {
		return false, err
	} else if err == redis.Nil {
		return false, nil
	}
	if len(freezeStr) != 0 {
		*freeze, err = json2Freeze([]byte(freezeStr))
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
