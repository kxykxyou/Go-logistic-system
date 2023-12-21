package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"logistic/internal/svc"
	"net/http"
)

func RegisterHandler(server *rest.Server, svcCtx *svc.Context) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/fake",
				Handler: FakeDataHandler(svcCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/query",
				Handler: QueryHandler(svcCtx),
			},
		})
}
