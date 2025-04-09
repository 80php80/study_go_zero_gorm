package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm/internal/logic"
	"gorm/internal/svc"
	"gorm/internal/types"
)

func GormHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsBase
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGormLogic(r.Context(), svcCtx)
		resp, err := l.Gorm(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
