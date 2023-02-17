package middleware

import (
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/token"
	"net/http"
)

type AuthMiddleware struct {
	token.Token
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//xToken := r.Header.Get("X-TOKEN")
		//m.ValidateAccessToken()

		// Passthrough to next handler if need
		next(w, r)
	}
}
