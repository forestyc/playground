package handler

import (
	"net/http"

	"github.com/Baal19905/playground/go-zero/epidemic/api/internal/logic"
	"github.com/Baal19905/playground/go-zero/epidemic/api/internal/svc"
	"github.com/Baal19905/playground/go-zero/epidemic/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MsgCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommonMsgCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMsgCodeLogic(r.Context(), svcCtx)
		resp, err := l.MsgCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}