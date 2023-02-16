package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewToken(t *testing.T) {
	token := Token{
		SignKey:       "adsfadsfadsf",
		AccessExpires: 100,
		RefreshExpire: 100,
	}
	access, refresh, err := token.GenToken("abcd")
	assert.Equal(t, err, nil)
	_, err = token.ValidateAccessToken(access)
	assert.Equal(t, err, nil)
	_, err = token.ValidateRefreshToken(refresh)
	assert.Equal(t, err, nil)
}

func TestValidateAccessToken(t *testing.T) {
	token := Token{
		SignKey:       "adsfadsfadsf",
		AccessExpires: 5,
		RefreshExpire: 100,
	}
	access, _, err := token.GenToken("abcd")
	assert.Equal(t, err, nil)
	access2, _, _ := token.GenToken("ffff")
	// 超时测试
	time.Sleep(time.Second * 6)
	_, err = token.ValidateAccessToken(access)
	assert.Equal(t, err != nil, true)

	// 篡改测试
	tmp := []byte(access)
	tmp[0] += 1
	access = string(tmp)
	_, err = token.ValidateAccessToken(access)
	assert.Equal(t, err != nil, true)
	// 非本用户token测试
	_, err = token.ValidateAccessToken(access2)
	assert.Equal(t, err != nil, true)
}

func TestValidateRefreshToken(t *testing.T) {
	token := Token{
		SignKey:       "adsfadsfadsf",
		AccessExpires: 5,
		RefreshExpire: 5,
	}
	_, refresh, err := token.GenToken("abcd")
	assert.Equal(t, err, nil)
	_, refresh2, _ := token.GenToken("ffff")
	// 超时测试
	time.Sleep(time.Second * 6)
	_, err = token.ValidateRefreshToken(refresh)
	assert.Equal(t, err != nil, true)

	// 篡改测试
	tmp := []byte(refresh)
	tmp[0] += 1
	refresh = string(tmp)
	_, err = token.ValidateRefreshToken(refresh)
	assert.Equal(t, err != nil, true)
	// 非本用户token测试
	_, err = token.ValidateRefreshToken(refresh2)
	assert.Equal(t, err != nil, true)
}

func TestRefresh(t *testing.T) {
	token := Token{
		SignKey:       "adsfadsfadsf",
		AccessExpires: 5,
		RefreshExpire: 5,
	}
	access, refresh, err := token.GenToken("abcd")
	assert.Equal(t, err, nil)
	access2, refresh2, err := token.Refresh(refresh)
	_, err = token.ValidateAccessToken(access2)
	assert.Equal(t, err, nil)
	_, err = token.ValidateRefreshToken(refresh2)
	assert.Equal(t, err, nil)
	assert.Equal(t, access == access2, false)
	assert.Equal(t, refresh == refresh2, false)
}
