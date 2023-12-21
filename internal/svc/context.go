package svc

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"logistic/internal/config"
)

type Context struct {
	// TODO
	DB    *gorm.DB
	Cache *redis.Client
}

func GetInitSvcContext() Context {

	cfg := config.GetConfigCtx()

	DBDsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.USER,
		cfg.DB.PASSWORD,
		cfg.DB.HOST,
		cfg.DB.PORT,
		cfg.DB.DBNAME,
	)

	db, err := gorm.Open(mysql.Open(DBDsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database.")
	}

	rds := GetRedisClient(cfg)

	return Context{
		DB:    db,
		Cache: rds,
	}
}
