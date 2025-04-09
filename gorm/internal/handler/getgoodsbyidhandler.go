package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm/internal/logic"
	"gorm/internal/svc"
)

func GetGoodsByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetGoodsByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetGoodsById()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
