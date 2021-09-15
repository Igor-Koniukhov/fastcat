package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/parser"
	"github.com/igor-koniukhov/fastcat/internal/server"
	"github.com/igor-koniukhov/fastcat/services"
	web "github.com/igor-koniukhov/webLogger/v3"
	"github.com/subosito/gotenv"
	"log"
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
	s := services.NewSetter(&app)
	driver, err := s.SetAndRun()
	if err != nil {
		log.Println(err)
	}
	defer driver.SQL.Close()
	go RunUpToDateSuppliersInfo(600)
	go func() {
		err := srv.Serve(
			os.Getenv("PORT"),
			routes(&app, driver.SQL),
		)
		if err != nil {
			log.Fatal(err)
		}
	}()
	web.Log.Info("FastCat application Started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	web.Log.Info("FastCat application Shutting Down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func RunUpToDateSuppliersInfo(t time.Duration) {

	app.ChanIdSupplier = make(chan int,4 )
	defer close(app.ChanIdSupplier)
	for {
		parser.ParseRestMenu.ParsedDataWriter()
		fmt.Println("Menu is up-to-date ")
		time.Sleep(time.Second * t)
	}
}
