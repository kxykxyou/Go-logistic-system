package handler

import (
	internalHttp "logistic/internal/http"
	"logistic/internal/logic"
	"logistic/internal/svc"
	"net/http"
)

func FakeDataHandler(svcCtx *svc.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFakeDataInsertionLogic(r, svcCtx)
		result, err := l.InsertFakeData()

		if err != nil {
			internalHttp.ErrorResponse(r.Context(), w, err)
		} else {
			internalHttp.SuccessResponse(r.Context(), w, result)
		}
	}
}
