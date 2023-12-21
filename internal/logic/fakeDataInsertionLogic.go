package logic

import (
	"fmt"
	"gorm.io/gorm"
	"logistic/internal/model"
	"logistic/internal/svc"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var allStatus = []string{"Created", "Package Received", "In Transit", "Out of Delivery", "Delivered", "Returned to Sender", "Exception"}

type FakeDataInsertionLogic struct {
	HttpRequestCtx *http.Request
	SvcCtx         *svc.Context
}

type FakeData struct {
	FakeOrders          []*model.Order
	FakeLogisticDetails []*model.LogisticDetail
}

func NewFakeDataInsertionLogic(r *http.Request, svcCtx *svc.Context) *FakeDataInsertionLogic {
	return &FakeDataInsertionLogic{
		HttpRequestCtx: r,
		SvcCtx:         svcCtx,
	}
}

func (logic FakeDataInsertionLogic) InsertFakeData() (result interface{}, e error) {
	var num int
	var err error
	num, err = strconv.Atoi(logic.HttpRequestCtx.URL.Query().Get("num"))

	if err != nil {
		return nil, err
	}

	fakeData := generateRandomData(num)

	err = bulkInsertFakeData(logic.SvcCtx.DB, &fakeData)
	if err != nil {
		return nil, err
	} else {
		return fmt.Sprintf("%d rows inserted", num), nil
	}

}

func generateRandomData(num int) FakeData {
	var fakeOrders []*model.Order
	var fakeLogisticDetails []*model.LogisticDetail

	for i := 0; i < num-1; i++ {
		newOrder := model.Order{
			RecipientId: uint(rand.Intn(4) + 1),
			ProductId:   uint(rand.Intn(4) + 1),
		}
		fakeOrders = append(fakeOrders, &newOrder)
	}

	dates := randDatesForAnOrder()
	for i := 0; i < len(allStatus); i++ {
		newDetail := model.LogisticDetail{
			OrderId:    uint(rand.Intn(4) + 1),
			LocationId: uint(rand.Intn(4) + 1),
			Date:       dates[i],
			Status:     allStatus[i],
		}
		fakeLogisticDetails = append(fakeLogisticDetails, &newDetail)
	}

	return FakeData{
		FakeOrders:          fakeOrders,
		FakeLogisticDetails: fakeLogisticDetails,
	}
}

func bulkInsertFakeData(db *gorm.DB, fakeData *FakeData) error {
	tx := db.Begin()

	if result := tx.Create(&fakeData.FakeOrders); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result := tx.Create(&fakeData.FakeLogisticDetails); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return tx.Commit().Error
}

func randDatesForAnOrder() []time.Time {
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	duration := endDate.Sub(startDate)
	timeDiff := time.Duration(rand.Int63n(int64(duration)))

	firstDate := startDate.Add(timeDiff)
	var dates []time.Time

	dates = append(dates, firstDate)
	var dateDiffs []int
	for i := 0; i < 6; i++ {
		dateDiffs = append(dateDiffs, rand.Intn(2))
	}

	tempDate := firstDate.AddDate(0, 0, 0)
	for _, dateDiff := range dateDiffs {
		tempDate = tempDate.AddDate(0, 0, dateDiff)
		dates = append(dates, tempDate)
	}
	return dates
}
