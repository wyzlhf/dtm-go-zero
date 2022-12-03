// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"dtm-go-zero/order/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/order/quickCreate",
				Handler: createHandler(serverCtx),
			},
		},
	)
}