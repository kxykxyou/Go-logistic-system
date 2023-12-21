package main

import (
	"database/sql"
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
	ctx := config.GetConfigCtx()

	createDataBase(ctx.DB.HOST, ctx.DB.PORT, ctx.DB.USER, ctx.DB.PASSWORD, ctx.DB.DBNAME)
}

func createDataBase(host string, port string, user string, password string, name string) {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)

	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	pwd, err := os.Getwd()

	sqlFile, err := ioutil.ReadFile(filepath.Join(pwd, "migration", "sql", "logistic_initialize_schema.sql"))

	sqlStatements := strings.Split(string(sqlFile), ";")

	if sqlStatements[len(sqlStatements)-1] == "" {
		sqlStatements = sqlStatements[:len(sqlStatements)-1]
	}

	for _, statement := range sqlStatements {
		if _, err = db.Exec(
			statement,
		); err != nil {
			if _, e := db.Exec("ROLLBACK"); e != nil {
				panic(e)
			}
			panic(err)
		}
	}
}
