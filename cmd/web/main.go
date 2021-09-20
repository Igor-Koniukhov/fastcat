package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-koniukhov/fastcat/driver"
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/parser"
	"github.com/igor-koniukhov/fastcat/internal/render"
	"github.com/igor-koniukhov/fastcat/internal/server"
	web "github.com/igor-koniukhov/webLogger/v2"
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
	dr, err := driver.ConnectDB("DSN")
	if err != nil {
		log.Fatal(err)
	}
	defer dr.SQL.Close()
	err = SetAndRun()
	if err != nil {
		log.Fatal(err)
	}

	go parser.RunUpToDateSuppliersInfo(600)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err, err)

	}
	go func() {
		err := srv.Serve(
			os.Getenv("PORT"),
			routes(&app, dr.SQL),
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

func SetAndRun() error {
	rest := parser.NewRestMenuParser(&app)
	parser.NewRestMenu(rest)
	render.NewTemplates(&app)
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not parse template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseTemplateCache = false
	app.TimeFormat = time.Now().UTC().Format("2006-01-02 15:04:05.999999")
	app.BearerString = os.Getenv("BearerString")
	logSet := web.NewLogStruct(&web.LogParameters{
		OutWriter:  web.ConsoleAndFile,
		FilePath:   "./logs",
		LogFile:    "/logger.log",
		TimeFormat: "[15:04:05||2006.01.02]",
	})
	web.NewLog(logSet)
	return nil
}

