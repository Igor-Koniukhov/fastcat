package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var err error


const (
	TableUser = "users"
	TableOrders = "orders"
	TableAddress = "address"
	TableCartProduct = "cart_products"
	TableProducts = "products"
)

func ConnectMySQLDB() *sql.DB {

	DSN := os.Getenv("DSN")
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
