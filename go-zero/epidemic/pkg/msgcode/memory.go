package msgcode

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 内存版本

type CodeByMobile map[string]Code     // 验证码map
type FreezeByMobile map[string]Freeze // 冻结信息map

type Memory struct {
	codeByMobile   CodeByMobile
	freezeByMobile FreezeByMobile
	frozenExpireAt time.Time
	lock           sync.RWMutex
	params         Params
}

func NewMemoryMsgCode(params Params) *Memory {
	m := Memory{
		params:         params,
		codeByMobile:   make(CodeByMobile),
		freezeByMobile: make(FreezeByMobile),
	}
	go m.purge()
	return &m
}

// Gen 生成验证码
func (mcg *Memory) Gen(account string) (string, error) {
	mcg.lock.Lock()
	defer mcg.lock.Unlock()
	// 校验是否冻结
	if mcg.frozen(account) {
		return "", errors.New("验证码发送过快，请稍后重试")
	}
	// 生成验证码
	code, exists := mcg.codeByMobile[account]
	now := time.Now()
	if !exists || code.ExpireAt.Before(now) { // 不存在或已过期
		code = initCode(mcg.params.ValidTime)
		mcg.codeByMobile[account] = code
		return code.Code, nil
	} else {
		timeLeft := code.ExpireAt.Sub(time.Now()).Seconds()
		left, _ := strconv.Atoi(fmt.Sprintf("%1.0f", timeLeft))
		return "", errors.New("验证码发送过快，请于" + strconv.Itoa(left) + "秒后重试")
	}
}

// Check 校验验证码
func (mcg *Memory) Check(account, code string) error {
	mcg.lock.RLock()
	defer mcg.lock.RUnlock()
	if account == "" || code == "" {
		return errors.New("验证码错误")
	}
	codeCached, exists := mcg.codeByMobile[account]
	if !exists {
		return errors.New("验证码失效")
	} else {
		if codeCached.Wrong >= mcg.params.WrongTime {
			// 超出最大错误数，立刻清掉验证码
			delete(mcg.codeByMobile, account)
			return errors.New("验证码失效")
		}
		now := time.Now()
		if codeCached.ExpireAt.Before(now) {
			return errors.New("验证码失效")
		}
		if codeCached.Code != code {
			codeCached.Wrong += 1
			mcg.codeByMobile[account] = codeCached
			return errors.New("验证码错误")
		}
		// 验证通过，删除验证码
		delete(mcg.codeByMobile, account)
		return nil
	}
}

// 清理过期信息
// 每30分钟请一次
func (mcg *Memory) purge() {
	time.Sleep(time.Minute * 30)
	mcg.lock.Lock()
	defer mcg.lock.Unlock()
	now := time.Now()
	for k, v := range mcg.codeByMobile {
		if v.ExpireAt.Before(now) {
			delete(mcg.codeByMobile, k)
		}
	}
	for k, v := range mcg.freezeByMobile {
		if v.ExpireAt.Equal(time.Time{}) {
			if v.Window.Before(now) {
				delete(mcg.freezeByMobile, k)
			}
		} else if v.ExpireAt.Before(now) {
			delete(mcg.freezeByMobile, k)
		}
	}
}

// 冻结
func (mcg *Memory) frozen(account string) bool {
	if freeze, exists := mcg.freezeByMobile[account]; exists {
		now := time.Now()
		if freeze.ExpireAt.Equal(time.Time{}) { // 未设置冻结
			if freeze.Count >= mcg.params.DurationCount { // 窗口内请求超限
				if freeze.Window.Before(time.Now()) { // 窗口已过期，不冻结
					mcg.freezeByMobile[account] = initFreeze(mcg.params.DurationWindow)
					return false
				} else { // 未过期，需要冻结
					freeze.ExpireAt = time.Now().Add(time.Second * time.Duration(mcg.params.FreezeTime))
					mcg.freezeByMobile[account] = freeze
					return true
				}
			} else { // 窗口内请求未超限
				freeze.Count += 1
				mcg.freezeByMobile[account] = freeze
			}
		} else { // 已设置冻结
			if freeze.ExpireAt.Before(now) { // 冻结已解除
				mcg.freezeByMobile[account] = initFreeze(mcg.params.DurationWindow)
				return false
			} else { // 冻结未接触，仍请求，重新计算冻结时间
				freeze.ExpireAt = time.Now().Add(time.Second * time.Duration(mcg.params.FreezeTime))
				return true
			}
		}
	} else {
		mcg.freezeByMobile[account] = initFreeze(mcg.params.DurationWindow)
	}
	return false
}
