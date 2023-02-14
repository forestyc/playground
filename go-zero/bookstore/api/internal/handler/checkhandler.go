package handler

import (
	"net/http"

	"github.com/Baal19905/playground/go-zero/bookstore/api/internal/logic"
	"github.com/Baal19905/playground/go-zero/bookstore/api/internal/svc"
	"github.com/Baal19905/playground/go-zero/bookstore/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCheckLogic(r.Context(), svcCtx)
		resp, err := l.Check(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
