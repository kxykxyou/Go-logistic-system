package logic

import (
	"gorm.io/gorm"
	"logistic/internal/model"
	"logistic/internal/svc"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var allStatus = []string{"Created", "Package Received", "In Transit", "Out of Delivery", "Delivered", "Returned to Sender", "Exception"}

type FakeDataInsertionLogic struct {
	HttpRequestCtx *http.Request
	SvcCtx         *svc.Context
}

type FirstClassFakeData struct {
	FakeLocations  []*model.Location
	FakeRecipients []*model.Recipient
	FakeProducts   []*model.Product
}

type SecondClassFakeData struct {
	FakeOrders []*model.Order
}

type ThirdClassFakeData struct {
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

	tx := logic.SvcCtx.DB.Begin()
	firstClassFakeData := generateFirstClassRandomData()
	if err := insertFirstClassFakeData(tx, &firstClassFakeData); err != nil {
		tx.Rollback()
		return nil, err
	}

	secondClassFakeData := generateSecondClassRandomData(tx, num)
	if err := insertSecondClassFakeData(tx, &secondClassFakeData); err != nil {
		tx.Rollback()
		return nil, err
	}

	thirdClassFakeData := generateThirdClassRandomData(tx)
	if err := insertThirdClassFakeData(tx, &thirdClassFakeData); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return "ok", nil
}

func generateFirstClassRandomData() FirstClassFakeData {

	var fakeLocations []*model.Location
	var fakeRecipients []*model.Recipient
	var fakeProducts []*model.Product

	for i := 1; i < 10; i++ {
		s := strconv.FormatInt(int64(i), 10)
		newLocation := model.Location{
			Title:   "縣市" + s,
			City:    "城市" + s,
			Address: "地址" + s,
		}
		fakeLocations = append(fakeLocations, &newLocation)

		newRecipient := model.Recipient{
			Name:    "王" + s,
			Address: strings.Repeat(s, 20),
			Phone:   "09" + strings.Repeat(s, 8),
		}
		fakeRecipients = append(fakeRecipients, &newRecipient)

		newProduct := model.Product{
			Name: "商品" + s,
		}
		fakeProducts = append(fakeProducts, &newProduct)

	}

	return FirstClassFakeData{
		FakeLocations:  fakeLocations,
		FakeRecipients: fakeRecipients,
		FakeProducts:   fakeProducts,
	}
}

func generateSecondClassRandomData(tx *gorm.DB, num int) SecondClassFakeData {

	var recipientIds []uint
	tx.Model(&model.Recipient{}).Pluck("id", &recipientIds)

	var productIds []uint
	tx.Model(&model.Product{}).Pluck("id", &productIds)

	var fakeOrders []*model.Order
	for i := 1; i < num-1; i++ {
		newOrder := model.Order{
			RecipientId: recipientIds[rand.Intn(len(recipientIds))],
			ProductId:   productIds[rand.Intn(len(productIds))],
		}
		fakeOrders = append(fakeOrders, &newOrder)
	}

	return SecondClassFakeData{
		FakeOrders: fakeOrders,
	}
}

func generateThirdClassRandomData(tx *gorm.DB) ThirdClassFakeData {

	var locationIds []uint
	tx.Model(&model.Location{}).Pluck("id", &locationIds)

	var orderIds []uint
	tx.Model(&model.Order{}).Pluck("id", &orderIds)

	var fakeLogisticDetails []*model.LogisticDetail
	dates := randDatesForAnOrder()
	for i := 0; i < len(allStatus); i++ {
		orderId := orderIds[rand.Intn(len(orderIds))]
		locationId := locationIds[rand.Intn(len(locationIds))]
		newDetail := model.LogisticDetail{
			OrderId:    orderId,
			LocationId: locationId,
			Date:       dates[i],
			Status:     allStatus[i],
		}
		fakeLogisticDetails = append(fakeLogisticDetails, &newDetail)

	}

	return ThirdClassFakeData{
		FakeLogisticDetails: fakeLogisticDetails,
	}
}

func insertFirstClassFakeData(tx *gorm.DB, firstClassFakeData *FirstClassFakeData) error {
	//if tx.Find(&model.Recipient{}).RowsAffected != 0 ||
	//	tx.Find(&model.Product{}).RowsAffected != 0 ||
	//	tx.Find(&model.Location{}).RowsAffected != 0 {
	//	return nil
	//}

	if result := tx.Create(&firstClassFakeData.FakeLocations); result.Error != nil {
		return result.Error
	}
	if result := tx.Create(&firstClassFakeData.FakeRecipients); result.Error != nil {
		return result.Error
	}
	if result := tx.Create(&firstClassFakeData.FakeProducts); result.Error != nil {
		return result.Error
	}

	return nil
}

func insertSecondClassFakeData(tx *gorm.DB, secondClassFakeData *SecondClassFakeData) error {
	//if tx.Find(&model.Recipient{}).RowsAffected == 0 || tx.Find(&model.Product{}).RowsAffected == 0 {
	//	return errors.New("not enough number of recipient or product in database")
	//}

	if result := tx.Create(&secondClassFakeData.FakeOrders); result.Error != nil {
		return result.Error
	}

	return nil
}

func insertThirdClassFakeData(tx *gorm.DB, thirdClassFakeData *ThirdClassFakeData) error {
	//if tx.Find(&model.Order{}).RowsAffected == 0 || tx.Find(&model.Location{}).RowsAffected == 0 {
	//	return errors.New("not enough number of order or location in database")
	//}

	if result := tx.Create(&thirdClassFakeData.FakeLogisticDetails); result.Error != nil {
		return result.Error
	}

	return nil
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
