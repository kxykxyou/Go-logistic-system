package http

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func Success(data interface{}) SuccessCtx {
	return SuccessCtx{
		Success: true,
		Data:    data,
	}
}

func SuccessResponse(ctx context.Context, w http.ResponseWriter, data interface{}) {
	var result interface{}
	if data == nil {
		result = struct{}{}
	} else {
		result = data
	}

	httpx.OkJsonCtx(ctx, w, Success(result))
}

func ErrorResponse(ctx context.Context, w http.ResponseWriter, err error) {
	httpx.ErrorCtx(ctx, w, err)
}
