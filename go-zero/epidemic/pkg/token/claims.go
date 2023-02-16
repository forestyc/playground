package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

const (
	accessTag  = "access"
	refreshTag = "refresh"
)

type Claims struct {
	jwt.RegisteredClaims
	Tag    string `json:"tab"`
	Random string `json:"rdm"`
}

type AccessClaims struct {
	Claims
}

type RefreshClaims struct {
	Claims
}

func newWithAccessClaims(userId string, exp time.Duration) *AccessClaims {
	claim := AccessClaims{}
	claim.ID = userId
	expiresAt := time.Now().Add(exp)
	claim.ExpiresAt = jwt.NewNumericDate(expiresAt)
	claim.Tag = accessTag
	claim.Random = uuid.NewV4().String()
	return &claim
}

func newWithRefreshClaims(userId string, exp time.Duration) *RefreshClaims {
	claim := RefreshClaims{}
	claim.ID = userId
	expiresAt := time.Now().Add(exp)
	claim.ExpiresAt = jwt.NewNumericDate(expiresAt)
	claim.Tag = refreshTag
	claim.Random = uuid.NewV4().String()
	return &claim
}

func (c AccessClaims) Valid() error {
	if c.Tag != accessTag {
		return errors.New("invalid tag")
	}
	if time.Now().After(c.ExpiresAt.Time) {
		vErr := new(jwt.ValidationError)
		vErr.Inner = errors.New("token is expired")
		vErr.Errors |= jwt.ValidationErrorExpired
		return vErr
	}
	return nil
}

func (c RefreshClaims) Valid() error {
	if c.Tag != refreshTag {
		return errors.New("invalid tag")
	}
	if time.Now().After(c.ExpiresAt.Time) {
		vErr := new(jwt.ValidationError)
		vErr.Inner = errors.New("token is expired")
		vErr.Errors |= jwt.ValidationErrorExpired
		return vErr
	}
	return nil
}
