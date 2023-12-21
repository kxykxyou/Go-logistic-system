package svc

import (
	"database/sql"
	"fmt"
	"logistic/internal/config"
)

func newDBConnection(driverName string, DBConfig config.DataBase) (*sql.DB, error) {
	dbConnection, err := sql.Open(driverName, makeDBSource(driverName, DBConfig))
	if err != nil {
		return nil, err
	}

	return dbConnection, nil
}

func makeDBSource(driverName string, DBConfig config.DataBase) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s?database=%s", driverName, DBConfig.USER, DBConfig.PASSWORD, DBConfig.HOST, DBConfig.PORT, DBConfig.DBNAME)
}
