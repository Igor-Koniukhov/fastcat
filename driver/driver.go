package driver

import (
	"database/sql"
	"fmt"
	"log"
)

var err error


const (
	TableUser = "users"
	TableOrders = "orders"
	TableAddress = "address"
	TableCartProduct = "cart_products"
	TableProducts = "products"
)

func ConnectMySQLDB(DSN string) *sql.DB {
	db, err := sql.Open("mysql", DSN)
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
