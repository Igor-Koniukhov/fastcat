package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/helpers"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"

	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

var (
	app      config.AppConfig
	infoLog  *log.Logger
	errorLog *log.Logger
)

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	DSN := os.Getenv("DSN")
	db := driver.ConnectMySQLDB(DSN)
	defer db.Close()

	app.DB = db
	app.Session="This is session"


	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	handlers.UserHandlers(&app)
	handlers.OrderHandlers(&app)
	handlers.SupplierHandlers(&app)
	handlers.ProductHandlers(&app)
	handlers.CartHandlers(&app)

	helpers.NewHelpers(&app)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Fatal(http.ListenAndServe(port, nil))
}
