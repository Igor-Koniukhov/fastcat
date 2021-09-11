package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/server"
	web "github.com/igor-koniukhov/webLogger/v3"
	"github.com/subosito/gotenv"
	"os"
	"os/signal"
	"time"
)

var app config.AppConfig

func init() {
	gotenv.Load()
}

func main() {
	srv := new(server.Server)
	driver, err := SetAndRun()
	defer driver.MySQL.Close()

	go RunUpToDateSuppliersInfo(600)
	go func() {
		err := srv.Serve(
			os.Getenv("PORT"),
			routes(&app, driver.MySQL),
			)
		web.Log.Fatal(err, err, " Got an error while running http server")
	}()
	web.Log.Info("FastCat application Started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	web.Log.Info("FastCat application Shutting Down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = srv.Shutdown(ctx)
	web.Log.Error(err, "Error on DB connection close: ", err)
}

