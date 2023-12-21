package logic

import (
	"logistic/internal/model"
	"logistic/internal/svc"
	"net/http"
)

type QueryLogisticDetailLogic struct {
	HttpRequestCtx *http.Request
	SvcCtx         *svc.Context
}

type QueryLogisticDetailResult struct {
	Sno            uint                    `json:"sno"`
	TrackingStatus string                  `json:"tracking_status"`
	Details        *[]model.LogisticDetail `json:"details"`
}

func NewQueryLogisticDetailLogic(r *http.Request, svcCtx *svc.Context) *QueryLogisticDetailLogic {
	return &QueryLogisticDetailLogic{
		HttpRequestCtx: r,
		SvcCtx:         svcCtx,
	}
}

func (logic QueryLogisticDetailLogic) QueryLogisticDetailsByOrderId(orderId uint) QueryLogisticDetailResult {
	var queryResult []model.LogisticDetail
	logic.SvcCtx.DB.Find(&queryResult, model.LogisticDetail{OrderId: orderId})

	result := QueryLogisticDetailResult{
		Sno:            orderId,
		TrackingStatus: queryResult[len(queryResult)-1].Status,
		Details:        &queryResult,
	}

	return result
}
