package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/handlers"
	"github.com/igor-koniukhov/fastcat/services"
	"github.com/igor-koniukhov/webLogger/v2"

	"time"

	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

var (
	app config.AppConfig
)

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	db := driver.ConnectMySQLDB()
	defer db.Close()
	app.DB = db
	app.Session = "This is session"
	app.TimeFormat = time.Now().UTC().Format("15:04:05||02.01.2006")
	fmt.Println(app.TimeFormat)

	setWebLoggerParameters()

	go runUpToDateDB(3600)

	handlers.UserHandlers(&app)
	handlers.OrderHandlers(&app)
	handlers.SupplierHandlers(&app)
	handlers.ProductHandlers(&app)
	handlers.CartHandlers(&app)

	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Fatal(http.ListenAndServe(port, nil))
}


func runUpToDateDB(t time.Duration) {
	rest := services.NewRestMenuRepository(&app)
	services.NewRestMenu(rest)
	suppL := services.ParserRestMenu.GetListSuppliers()
	app.ChanelSupplierId = make(chan int, len(suppL.Restaurants))
	app.ChanelLockUnlock = make(chan int, 1)
	for {
		services.ParserRestMenu.ParsedDataWriter()
		fmt.Println("Menu is up-to-date ")
		time.Sleep(time.Second * t)
	}
}

func setWebLoggerParameters() {
	logSet := webLogger.NewLogStruct(&webLogger.LogParameters{
		OutWriter:  webLogger.ConsoleAndFile,
		FilePath:   "./logs",
		LogFile:    "/logger.log",
		TimeFormat: "[15:04:05||2006.01.02]",
	})
	webLogger.NewLog(logSet)
}
