package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm/internal/logic"
	"gorm/internal/svc"
	"gorm/internal/types"
)

func CreateGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddGoodsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateGoodsLogic(r.Context(), svcCtx)
		resp, err := l.CreateGoods(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
