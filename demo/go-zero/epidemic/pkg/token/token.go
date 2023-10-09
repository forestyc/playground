package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Token 双Token
type Token struct {
	SignKey       string
	AccessExpires int
	RefreshExpire int
}

// GenToken 生成token
func (t Token) GenToken(id string) (string, string, error) {
	var err error
	var access, refresh string
	accessClaims := newWithAccessClaims(id, time.Duration(t.AccessExpires)*time.Second)
	if access, err = t.accessToken(accessClaims); err != nil {
		return access, refresh, err
	}
	refreshClaims := newWithRefreshClaims(id, time.Duration(t.RefreshExpire)*time.Second)
	if refresh, err = t.refreshToken(refreshClaims); err != nil {
		return access, refresh, err
	}
	return access, refresh, err
}

// ValidateAccessToken 校验AccessToken
func (t Token) ValidateAccessToken(token string) (*AccessClaims, error) {
	tk, err := jwt.ParseWithClaims(token, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.SignKey), nil
	})

	if err != nil {
		return nil, err
	}
	if !tk.Valid {
		return nil, errors.New("invalid token")
	}
	return tk.Claims.(*AccessClaims), nil
}

// ValidateRefreshToken 校验RefreshToken
func (t Token) ValidateRefreshToken(token string) (*RefreshClaims, error) {
	tk, err := jwt.ParseWithClaims(token, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !tk.Valid {
		return nil, errors.New("invalid token")
	}
	return tk.Claims.(*RefreshClaims), nil
}

// Refresh 刷新双token
func (t Token) Refresh(token string) (string, string, error) {
	if refreshClaims, err := t.ValidateRefreshToken(token); err != nil {
		return "", "", err
	} else {
		return t.GenToken(refreshClaims.ID)
	}
}

// 生成AccessToken
func (t Token) accessToken(claims jwt.Claims) (string, error) {
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tk.SignedString([]byte(t.SignKey))
	if err != nil {
		return "", err
	}
	return token, err
}

// 生成RefreshToken
func (t Token) refreshToken(claims jwt.Claims) (string, error) {
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tk.SignedString([]byte(t.SignKey))
	if err != nil {
		return token, err
	}
	return token, nil
}
