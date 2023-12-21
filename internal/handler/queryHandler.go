package handler

import (
	internalHttp "logistic/internal/http"
	"logistic/internal/logic"
	"logistic/internal/svc"
	"net/http"
	"strconv"
)

func QueryHandler(svcCtx *svc.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewQueryLogisticDetailLogic(r, svcCtx)

		orderId, err := strconv.Atoi(r.URL.Query().Get("sno"))
		if err != nil {
			internalHttp.ErrorResponse(r.Context(), w, err)
			return
		}

		data := l.QueryLogisticDetailsByOrderId(uint(orderId))
		internalHttp.SuccessResponse(r.Context(), w, data)
	}
}
