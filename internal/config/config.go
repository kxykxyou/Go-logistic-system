package config

import (
	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/rest"
	"os"
	"strconv"
)

type Config struct {
	rest.RestConf
	Redis Redis
	DB    DataBase
}

type Redis struct {
	ADDRESS           string
	PASSWORD          string
	DB                int `json:"db"`
	IdleTimeoutSecond int
}

type DataBase struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DBNAME   string
}

func GetConfigCtx() Config {
	err := godotenv.Load()
	if err != nil {
		panic("No env file found")
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPass := os.Getenv("REDIS_REQUIREPASS")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUM"))
	redisTimeoutSecond, _ := strconv.Atoi(os.Getenv("REDIS_TIMEOUT_SECONDS"))

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if len(dbHost) == 0 || len(dbPort) == 0 || len(dbUser) == 0 || len(dbPassword) == 0 || len(dbName) == 0 {
		panic("DB ctx is not set")
	}
	ctx := Config{
		Redis: Redis{
			ADDRESS:           redisHost,
			PASSWORD:          redisPass,
			DB:                redisDB,
			IdleTimeoutSecond: redisTimeoutSecond,
		},
		DB: DataBase{
			HOST:     dbHost,
			PORT:     dbPort,
			USER:     dbUser,
			PASSWORD: dbPassword,
			DBNAME:   dbName,
		},
	}

	return ctx
}
