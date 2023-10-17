package jwt

import (
	"testing"
	"time"

	"github.com/forestyc/playground/pkg/utils"
	"github.com/go-playground/assert/v2"
)

var secret string = utils.GenUuid()

func TestTokenWithIssuer(t *testing.T) {
	id := "abcd"
	issuer := "hahaha"
	token, err := NewToken(id, []byte(secret), WithIssuer(issuer))
	assert.Equal(t, err, nil)
	assert.NotEqual(t, token, "")
	claims, err := ValidateToken(token, []byte(secret))
	assert.Equal(t, err, nil)
	assert.NotEqual(t, claims, nil)
	assert.Equal(t, claims.ID, id)
	assert.Equal(t, claims.Issuer, issuer)
}

func TestTokenWithSubject(t *testing.T) {
	id := "abcd"
	subject := "hahaha"
	token, err := NewToken(id, []byte(secret), WithSubject(subject))
	assert.Equal(t, err, nil)
	assert.NotEqual(t, token, "")
	claims, err := ValidateToken(token, []byte(secret))
	assert.Equal(t, err, nil)
	assert.NotEqual(t, claims, nil)
	assert.Equal(t, claims.ID, id)
	assert.Equal(t, claims.Subject, subject)
}

func TestTokenWithAudience(t *testing.T) {
	id := "abcd"
	audience := "read"
	token, err := NewToken(id, []byte(secret), WithAudience(audience))
	assert.Equal(t, err, nil)
	assert.NotEqual(t, token, "")
	claims, err := ValidateToken(token, []byte(secret))
	assert.Equal(t, err, nil)
	assert.NotEqual(t, claims, nil)
	assert.Equal(t, claims.ID, id)
	assert.Equal(t, claims.Audience[0], audience)
}

func TestTokenWithExpireAt(t *testing.T) {
	id := "abcd"
	token, err := NewToken(id, []byte(secret), WithExpireAt(time.Now().Add(time.Second)))
	assert.Equal(t, err, nil)
	assert.NotEqual(t, token, "")
	claims, err := ValidateToken(token, []byte(secret))
	assert.Equal(t, err, nil)
	assert.NotEqual(t, claims, nil)
	assert.Equal(t, claims.ID, id)
	// wait 1s
	time.Sleep(time.Second)
	_, err = ValidateToken(token, []byte(secret))
	assert.NotEqual(t, err, nil)
}
