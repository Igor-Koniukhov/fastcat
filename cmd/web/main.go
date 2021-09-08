package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/server"
	web "github.com/igor-koniukhov/webLogger/v3"
	"github.com/subosito/gotenv"
	"os"
	"os/signal"
	"syscall"
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
	go RunUpToDateSuppliersInfo(600)

	srv := new(server.Server)
	go func() {
		err := srv.Run(port, routes(&app))
		web.Log.Fatal(err, err, " got an error while running http server")
	}()

	web.Log.Info("FastCat application Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	web.Log.Info("FastCat application Shutting Down")

	err := srv.Shutdown(context.Background())
	web.Log.Error(err, err, "got an error on DB connection close")
}

