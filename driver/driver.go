package driver

import (
	"database/sql"
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"log"
)

var err error

const DataSName = "fastcat:fastCat@tcp(127.0.0.1:3306)/fastcat_db"

func ConnectMySQLDB(app *config.AppConfig) *sql.DB {
	app.DB, err = sql.Open("mysql", DataSName)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("you connected to DB")
	}

	err = app.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return app.DB
}
