package driver

import (
	"database/sql"
	"fmt"
	"log"
)

var err error

const DataSName = "fastcat:fastCat@tcp(127.0.0.1:3306)/fastcat_db"
const TableUser = "user"

func ConnectMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", DataSName)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("you connected to DB")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
