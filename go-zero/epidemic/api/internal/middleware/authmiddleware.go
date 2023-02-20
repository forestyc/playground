package middleware

import (
	"errors"
	"github.com/Baal19905/playground/go-zero/epidemic/api/internal/config"
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/token"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type AuthMiddleware struct {
	token.Token
}

func NewAuthMiddleware(tokenConf config.Token) *AuthMiddleware {
	auth := &AuthMiddleware{}
	auth.SignKey = tokenConf.SignKey
	return auth
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		xToken := r.Header.Get("X-TOKEN")
		_, err1 := m.ValidateAccessToken(xToken)
		_, err2 := m.ValidateRefreshToken(xToken)
		if err1 != nil && err2 != nil {
			logx.Error("invalid X-TOKEN", err1, err2)
			httpx.ErrorCtx(r.Context(), w, errors.New("invalid X-TOKEN"))
			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}
