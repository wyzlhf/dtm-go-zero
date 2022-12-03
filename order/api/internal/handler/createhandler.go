package handler

import (
	"net/http"

	"dtm-go-zero/order/api/internal/logic"
	"dtm-go-zero/order/api/internal/svc"
	"dtm-go-zero/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuickCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
