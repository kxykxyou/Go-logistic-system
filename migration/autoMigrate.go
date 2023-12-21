// cannot work due to go version, the least requirement >= 1.20
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"logistic/internal/config"
	"logistic/internal/model"
)

func main() {
	ctx := config.GetConfigCtx()
	fmt.Printf("%s", ctx.DB.PORT)

	autoMigrate(ctx.DB.HOST, ctx.DB.PORT, ctx.DB.USER, ctx.DB.PASSWORD, ctx.DB.DBNAME)

}

func autoMigrate(host string, port string, user string, password string, name string) {

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
	fmt.Printf("%s\n", dbSource)
	db, err := gorm.Open(mysql.Open(dbSource), &gorm.Config{})
	if err != nil {
		panic("Can not connect to db source")
	}

	if err := db.AutoMigrate(&model.Location{}, &model.Product{}, &model.Recipient{}, &model.Order{}, &model.LogisticDetail{}); err != nil {
		fmt.Println(err)
	}
}
