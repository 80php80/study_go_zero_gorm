package handler

import (
	"fmt"
	"gorm/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm/internal/logic"
	"gorm/internal/svc"
)

func GetGoodsByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetGoodsByIdLogic(r.Context(), svcCtx)
		var req types.GetUserByIdReq

		if err := httpx.ParsePath(r, &req); err != nil {
			fmt.Printf("Error parsing path: %v\n", err) // 打印错误日志
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Printf("Parsed request: %+v\n", req) // 打印解析结果
		resp, err := l.GetGoodsById(req.ID)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
