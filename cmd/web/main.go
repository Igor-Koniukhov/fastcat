package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

var app config.AppConfig

func init() {
	gotenv.Load()
}

func main() {
	db := driver.ConnectMySQLDB()
	defer db.Close()
	port := os.Getenv("PORT")

	SetAppConfigParameters(db)
	SetWebLoggerParameters()
	go RunUpToDateSuppliersInfo(1)

	handlers.UserHandlers(&app)
	handlers.OrderHandlers(&app)
	handlers.SupplierHandlers(&app)
	handlers.ProductHandlers(&app)
	handlers.CartHandlers(&app)

	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Fatal(http.ListenAndServe(port, nil))
}





