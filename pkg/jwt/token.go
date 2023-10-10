package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/forestyc/playground/pkg/util"
	"github.com/golang-jwt/jwt/v4"
)

type Option func(*Claims)

type Claims struct {
	jwt.RegisteredClaims
}

// NewToken generate tokenString with id and secret
func NewToken(id string, secret []byte, option ...Option) (string, error) {
	var token string
	var err error
	c := &Claims{
		jwt.RegisteredClaims{
			Issuer:    "forestyc",
			Subject:   util.GenUuid(),
			NotBefore: &jwt.NumericDate{Time: time.Now()},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ID:        id,
		},
	}
	for _, o := range option {
		o(c)
	}
	if token, err = genTokenString(c, secret); err != nil {
		return "", err
	}
	return token, nil
}

// ValidateToken return claims if tokenString is validate.
func ValidateToken(tokenString string, secret []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims.(*Claims), nil
}

// WithIssuer adds issuer. By default forestyc.
func WithIssuer(issuer string) Option {
	return func(c *Claims) {
		c.Issuer = issuer
	}
}

// WithSubject adds subject. By default uuid.
func WithSubject(subject string) Option {
	return func(c *Claims) {
		c.Subject = subject
	}
}

// WithAudience adds audience.
func WithAudience(audience string) Option {
	return func(c *Claims) {
		c.Audience = append(c.Audience, audience)
	}
}

// WithExpireAt adds expireAt.
func WithExpireAt(expireAt time.Time) Option {
	return func(c *Claims) {
		c.ExpiresAt = &jwt.NumericDate{Time: expireAt}
	}
}

// WithNotBefore adds notBefore. Default time.Now().
func WithNotBefore(notBefore time.Time) Option {
	return func(c *Claims) {
		c.NotBefore = &jwt.NumericDate{Time: notBefore}
	}
}

// WithIssuedAt adds issuedAt. By default time.Now().
func WithIssuedAt(issuedAt time.Time) Option {
	return func(c *Claims) {
		c.IssuedAt = &jwt.NumericDate{Time: issuedAt}
	}
}

// genTokenString generate tokenString with claim and secret
func genTokenString(claims jwt.Claims, secret []byte) (string, error) {
	if secret == nil {
		return "", errors.New("invalid conf")
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tk.SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, err
}
