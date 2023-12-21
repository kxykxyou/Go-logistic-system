package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"logistic/internal/config"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cfg := config.GetConfigCtx()

	err := insertSeedData(cfg.DB.HOST, cfg.DB.PORT, cfg.DB.USER, cfg.DB.PASSWORD, cfg.DB.DBNAME)
	if err != nil {
		panic(fmt.Sprintf("insert seed data failed: %s", err))
	}
}

func insertSeedData(host string, port string, user string, password string, name string) error {
	if len(host) == 0 || len(port) == 0 || len(user) == 0 || len(password) == 0 || len(name) == 0 {
		return errors.New("DB config is not set")
	}

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)

	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	pwd, _ := os.Getwd()

	fileNames, _ := filepath.Glob(filepath.Join(pwd, "migration", "seeders", "*.sql"))

	for _, fileName := range fileNames {
		sqlFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			return errors.New(fmt.Sprintf("file reading failed: %s", string(sqlFile)))
		}

		sqlStatements := strings.Split(string(sqlFile), ";")
		if err := executeSql(sqlStatements, db); err != nil {
			return err
		}

	}

	return nil

}

func executeSql(sqlStatements []string, db *sql.DB) error {

	if sqlStatements[len(sqlStatements)-1] == "" {
		sqlStatements = sqlStatements[:len(sqlStatements)-1]
	}

	for _, statement := range sqlStatements {
		if _, err := db.Exec(
			statement,
		); err != nil {
			if _, e := db.Exec("ROLLBACK"); e != nil {
				return err
			}
			return err
		}
	}

	return nil
}
