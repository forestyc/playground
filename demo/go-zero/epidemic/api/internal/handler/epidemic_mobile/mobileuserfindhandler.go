package epidemic_mobile

import (
	"net/http"

	"github.com/forestyc/playground/go-zero/epidemic/api/internal/logic/epidemic_mobile"
	"github.com/forestyc/playground/go-zero/epidemic/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MobileUserFindHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := epidemic_mobile.NewMobileUserFindLogic(r.Context(), svcCtx)
		token := r.Header.Get("X-TOKEN")
		resp, err := l.MobileUserFind(token)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
